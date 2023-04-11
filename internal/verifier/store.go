package verifier

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type store interface {
	Create(ctx context.Context, v *Verifier) error
	ReadByUserID(ctx context.Context, userID uuid.UUID, v *Verifier) error
	Update(ctx context.Context, userID uuid.UUID, commType, otp string) error
}

type verifierStore struct {
	db *datastore.Client
}

type VerifierModel struct {
	datastore.Model
	UserID          uuid.UUID
	Password        string
	EmailOTP        string
	IsEmailVerified bool
	PhoneOTP        string
	IsPhoneVerified bool
}

func (vm *VerifierModel) prepare(verifier *Verifier) {
	vm.UserID = verifier.UserID
	vm.EmailOTP = verifier.Email.OTP
	vm.IsEmailVerified = verifier.Email.IsVerified
	vm.PhoneOTP = verifier.Phone.OTP
	vm.IsPhoneVerified = verifier.Phone.IsVerified
	vm.Password = verifier.Password

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

func (vs *verifierStore) ReadByUserID(ctx context.Context, userID uuid.UUID, v *Verifier) error {
	verModel := &VerifierModel{}
	err := vs.db.Where("user_id=?", userID).Find(verModel).Error
	if err != nil {
		return err
	}

	structToModel(v, verModel)
	return nil
}

func (vs *verifierStore) Update(ctx context.Context, userID uuid.UUID, commType, otp string) error {
	var err error
	if commType == "email" {
		err = vs.db.Model(&VerifierModel{}).Where(&VerifierModel{
			UserID:   userID,
			EmailOTP: otp,
		}).Update(
			"is_email_verified", true,
		).Error
	} else if commType == "phone" {
		err = vs.db.Model(&VerifierModel{}).Where(&VerifierModel{
			UserID:   userID,
			PhoneOTP: otp,
		}).Update(
			"is_phone_verified", true,
		).Error
	} else {
		return errors.New("invalid type provided")
	}
	if err != nil {
		return err
	}

	return nil
}

func structToModel(v *Verifier, model *VerifierModel) {
	v.ID = model.ID
	v.UserID = model.UserID
	v.Password = model.Password
	v.Email = &VerificationDetails{
		OTP:        model.EmailOTP,
		IsVerified: model.IsEmailVerified,
	}
	v.Phone = &VerificationDetails{
		OTP:        model.PhoneOTP,
		IsVerified: model.IsPhoneVerified,
	}
	v.CreatedAt = model.CreatedAt
	v.UpdatedAt = model.UpdatedAt
}

func newStore(db *datastore.Client) (*verifierStore, error) {
	return &verifierStore{
		db: db,
	}, nil
}
