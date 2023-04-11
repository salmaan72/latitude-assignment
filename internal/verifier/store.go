package verifier

import (
	"context"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type store interface {
	Create(ctx context.Context, v *Verifier) error
}

type verifierStore struct {
	db *datastore.Client
}

type VerifierModel struct {
	*datastore.Model
	UserID          uuid.UUID
	EmailOTP        string
	IsEmailVerified bool
	PhoneOTP        string
	isPhoneVerified bool
}

func (vm *VerifierModel) prepare(verifier *Verifier) {
	vm.UserID = verifier.UserID
	vm.EmailOTP = verifier.Email.OTP
	vm.IsEmailVerified = verifier.Email.IsVerified
	vm.PhoneOTP = verifier.Phone.OTP
	vm.isPhoneVerified = verifier.Phone.IsVerified

	vm.ID = verifier.ID
	vm.CreatedAt = verifier.CreatedAt
}

func (vs *verifierStore) Create(ctx context.Context, v *Verifier) error {
	vm := &VerifierModel{}
	vm.prepare(v)

	err := vs.db.Create(vm).Error
	if err != nil {
		return err
	}

	return nil
}

func newStore(db *datastore.Client) (*verifierStore, error) {
	return &verifierStore{
		db: db,
	}, nil
}
