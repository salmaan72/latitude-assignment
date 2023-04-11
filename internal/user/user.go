package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
	"github.com/salmaan72/latitude-assignment/internal/user/ledger"
)

type Service struct {
	store         store
	ledgerService *ledger.Service
}

type status string

var (
	StatusPending  = status("pending")
	StatusApproved = status("approved")
)

type User struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	Username  string     `json:"username,omitempty"`
	Email     string     `json:"email,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Address   *Address   `json:"address,omitempty"`
	Status    *status    `json:"status,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type Address struct {
	Line1   string `json:"line1,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
	Pincode string `json:"pincode,omitempty"`
}

func (u *User) fetchFromModelsBasic(um *UserModel, am *AddressModel) {
	u.ID = um.ID
	u.CreatedAt = um.CreatedAt
	u.UpdatedAt = um.UpdatedAt
	u.Status = &um.Status
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

	return u, nil
}

// func (s *Service) Read(id string) (*User, error) {
// 	user := &User{}
// 	err := s.store.Read(id, user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (s *Service) ReadByUsername(username string) (*User, error) {
// 	user := &User{}
// 	err := s.store.ReadByUsername(username, user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

func NewService(db *datastore.Client, ledgerService *ledger.Service) (*Service, error) {
	nstore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Service{
		store:         nstore,
		ledgerService: ledgerService,
	}, nil
}
