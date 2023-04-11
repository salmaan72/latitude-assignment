package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
	"github.com/salmaan72/latitude-assignment/internal/user/ledger"
	"github.com/salmaan72/latitude-assignment/internal/verifier"
)

type Service struct {
	store           store
	ledgerService   *ledger.Service
	verifierService *verifier.Service
}

type status string

var (
	StatusPending       = status("pending")
	StatusApproved      = status("approved")
	StatusEmailVerified = status("email_verified")
	StatusPhoneVerified = status("phone_verified")
)

type User struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Email     string     `json:"email,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Address   *Address   `json:"address,omitempty"`
	Status    status     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`

	Password string `json:"password,omitempty"`
}

type Address struct {
	Line1   string `json:"line1,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
	Pincode string `json:"pincode,omitempty"`
}

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (u *User) fetchFromModelsBasic(um *UserModel, am *AddressModel) {
	u.ID = um.ID
	u.CreatedAt = um.CreatedAt
	u.UpdatedAt = um.UpdatedAt
	u.Status = um.Status
}

func (s *Service) CreateUser(ctx context.Context, u *User) (*User, error) {
	err := s.store.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	_, err = s.ledgerService.Create(ctx, &ledger.Ledger{
		UserID: u.ID,
	})
	if err != nil {
		return nil, err
	}

	_, err = s.verifierService.CreateDefaults(ctx, &verifier.Verifier{
		UserID:   u.ID,
		Password: s.verifierService.PassHash([]byte(u.Password)),
	})
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) ReadByEmail(ctx context.Context, email string) (*User, error) {
	u := &User{}
	err := s.store.ReadByEmail(ctx, email, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) VerifyUser(ctx context.Context, userID uuid.UUID, commType, otp string) (*User, error) {
	verifier, err := s.verifierService.Verify(ctx, userID, commType, otp)
	if err != nil {
		return nil, err
	}

	if verifier.Email.IsVerified && verifier.Phone.IsVerified {
		// approved status
		err = s.store.Update(ctx, userID, StatusApproved)
	} else if verifier.Email.IsVerified && commType == "email" {
		// email_verifier status
		err = s.store.Update(ctx, userID, StatusEmailVerified)
	} else if verifier.Phone.IsVerified && commType == "phone" {
		// phone_verified status
		err = s.store.Update(ctx, userID, StatusPhoneVerified)
	}
	if err != nil {
		return nil, err
	}

	u := &User{}
	err = s.store.Read(ctx, userID, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func NewService(db *datastore.Client, ledgerService *ledger.Service, verifierService *verifier.Service) (*Service, error) {
	nstore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Service{
		store:           nstore,
		ledgerService:   ledgerService,
		verifierService: verifierService,
	}, nil
}
