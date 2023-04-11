package verifier

import (
	"context"
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
	ID        uuid.UUID
	UserID    uuid.UUID
	Email     *VerificationDetails
	Phone     *VerificationDetails
	CreatedAt *time.Time
	UpdatedAt *time.Time
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

	v.Email.OTP = fmt.Sprintf("%04d", rand.Int63n(1e4))
	v.Email.IsVerified = false
	v.Phone.OTP = fmt.Sprintf("%04d", rand.Int63n(1e4))
	v.Phone.IsVerified = false

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

func NewService(db *datastore.Client) (*Service, error) {
	newStore, err := newStore(db)
	if err != nil {
		return nil, err
	}

	return &Service{
		store: newStore,
	}, nil
}
