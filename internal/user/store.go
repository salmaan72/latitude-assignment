package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type store interface {
	Create(ctx context.Context, u *User) error
	ReadByEmail(ctx context.Context, email string, u *User) error
	Read(ctx context.Context, id uuid.UUID, u *User) error
	Update(ctx context.Context, id uuid.UUID, status status) error
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

func (nus *newUserStore) ReadByEmail(ctx context.Context, email string, u *User) error {
	userModel := &UserModel{}
	err := nus.db.Where("email=?", email).Find(userModel).Error
	if err != nil {
		return err
	}

	structToModel(u, userModel)
	return nil
}

func (nus *newUserStore) Read(ctx context.Context, id uuid.UUID, u *User) error {
	userModel := &UserModel{}
	err := nus.db.Where("id=?", id.String()).Find(userModel).Error
	if err != nil {
		return err
	}

	structToModel(u, userModel)
	return nil
}

func (nus *newUserStore) Update(ctx context.Context, id uuid.UUID, status status) error {
	err := nus.db.Model(&UserModel{}).Where(
		"id=?", id.String(),
	).Update(
		"status", status,
	).Error
	if err != nil {
		return err
	}

	return nil
}

func structToModel(u *User, model *UserModel) {
	u.ID = model.ID
	u.Email = model.Email
	u.Username = model.Username
	u.Phone = model.Phone
	u.Status = model.Status
}

func newStore(db *datastore.Client) (*newUserStore, error) {
	return &newUserStore{
		db: db,
	}, nil
}
