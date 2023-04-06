package user

// import (
// 	"errors"
// )

// type mockStore struct{}

// var users = []User{
// 	{
// 		ID:       "1",
// 		Username: "salmaan",
// 		Password: "pass1",
// 	},
// 	{
// 		ID:       "2",
// 		Username: "ayshu",
// 		Password: "pass2",
// 	},
// 	{
// 		ID:       "9",
// 		Username: "admin",
// 		Password: "adminpass",
// 	},
// }

// func (ms *mockStore) Read(id string, u *User) error {
// 	user, err := findByID(id)
// 	if err != nil {
// 		return err
// 	}

// 	*u = *user

// 	return nil
// }

// func (ms *mockStore) ReadByUsername(username string, u *User) error {
// 	user, err := findByUsername(username)
// 	if err != nil {
// 		return err
// 	}

// 	*u = *user

// 	return nil
// }

// func findByID(id string) (*User, error) {
// 	user := new(User)
// 	for _, u := range users {
// 		if u.ID == id {
// 			*user = u
// 			break
// 		}
// 	}

// 	if user.ID == "" {
// 		return nil, errors.New("no user found")
// 	}

// 	return user, nil
// }

// func findByUsername(username string) (*User, error) {
// 	user := new(User)
// 	for _, u := range users {
// 		if u.Username == username {
// 			*user = u
// 			break
// 		}
// 	}

// 	if user.Username == "" {
// 		return nil, errors.New("no user found")
// 	}

// 	return user, nil
// }

// func newMockStore() (*mockStore, error) {
// 	return &mockStore{}, nil
// }
