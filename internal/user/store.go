package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type store interface {
	Create(ctx context.Context, u *User) error
}

type newUserStore struct {
	db *datastore.Client
}

type UserModel struct {
	datastore.Model
	Username string
	Email    string
	Phone    string
	Status   status
}

func (um *UserModel) prepare(u *User) {
	um.Username = u.Username
	um.Email = u.Email
	um.Phone = u.Phone
	um.Status = StatusPending

	um.ID = uuid.New()
	now := time.Now()
	um.CreatedAt = &now
	um.UpdatedAt = &now
}

type AddressModel struct {
	datastore.Model
	Line1   string
	City    string
	State   string
	Country string
	Pincode string
	UserID  uuid.UUID
}

func (am *AddressModel) prepare(um *UserModel, u *User) {
	am.Line1 = u.Address.Line1
	am.City = u.Address.City
	am.State = u.Address.State
	am.Country = u.Address.Country
	am.Pincode = u.Address.Pincode
	am.UserID = um.ID

	am.ID = uuid.New()
	now := time.Now()
	am.CreatedAt = &now
	am.UpdatedAt = &now
}

func (nus *newUserStore) Create(ctx context.Context, user *User) error {
	um := &UserModel{}
	um.prepare(user)

	am := &AddressModel{}
	am.prepare(um, user)

	err := nus.db.Create(um).Error
	if err != nil {
		return err
	}

	err = nus.db.Create(am).Error
	if err != nil {
		return err
	}

	user.fetchFromModelsBasic(um, am)

	return nil
}

func newStore(db *datastore.Client) (*newUserStore, error) {
	return &newUserStore{
		db: db,
	}, nil
}
