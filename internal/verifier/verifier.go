package verifier

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type Service struct {
	store store
}

type Verifier struct {
	ID        uuid.UUID            `json:"id,omitempty"`
	UserID    uuid.UUID            `json:"userID,omitempty"`
	Password  string               `json:"password,omitempty"`
	Email     *VerificationDetails `json:"email,omitempty"`
	Phone     *VerificationDetails `json:"phone,omitempty"`
	CreatedAt *time.Time           `json:"createdAt,omitempty"`
	UpdatedAt *time.Time           `json:"updatedAt,omitempty"`
}

func (v *Verifier) checkInit() *Verifier {
	if v.Email == nil {
		v.Email = &VerificationDetails{}
	}
	if v.Phone == nil {
		v.Phone = &VerificationDetails{}
	}

	return v
}

func (ver *Verifier) validate() error {
	if ver.UserID.String() == "" {
		return errors.New("user id not provided")
	}

	return nil
}

type VerificationDetails struct {
	OTP        string
	IsVerified bool
}

func (s *Service) CreateDefaults(ctx context.Context, v *Verifier) (*Verifier, error) {
	v.ID = uuid.New()
	now := time.Now()
	v.CreatedAt = &now
	v.UpdatedAt = &now

	rand.Seed(time.Now().Unix())
	v.checkInit().Email.OTP = fmt.Sprintf("%04d", rand.Int63n(1e4))
	v.checkInit().Email.IsVerified = false
	rand.Seed(time.Now().Unix())
	v.checkInit().Phone.OTP = fmt.Sprintf("%04d", rand.Int63n(1e4))
	v.checkInit().Phone.IsVerified = false

	err := v.validate()
	if err != nil {
		return nil, err
	}

	err = s.store.Create(ctx, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Service) PassHash(pass []byte) string {
	hashed := sha256.Sum256(pass)

	return fmt.Sprintf("%x", hashed)
}

func (s *Service) Verify(ctx context.Context, userID uuid.UUID, comType string, otp string) (*Verifier, error) {
	v := &Verifier{}
	err := s.store.ReadByUserID(ctx, userID, v)
	if err != nil {
		return nil, err
	}

	if comType == "email" {
		if v.Email.OTP != otp {
			return nil, errors.New("invalid otp")
		}
		v.Email.IsVerified = true
	} else if comType == "phone" {
		if v.Phone.OTP != otp {
			return nil, errors.New("invalid otp")
		}
		v.Phone.IsVerified = true
	}
	err = s.store.Update(ctx, userID, comType, otp)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Service) ReadByUserID(ctx context.Context, userID uuid.UUID) (*Verifier, error) {
	v := &Verifier{}
	err := s.store.ReadByUserID(ctx, userID, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func NewService(db *datastore.Client) (*Service, error) {
	newStore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Service{
		store: newStore,
	}, nil
}
