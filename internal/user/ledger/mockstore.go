package ledger

// import "errors"

// type store interface {
// 	Read(userID string, l *Ledger) error
// 	Update(userID string, inputData *Ledger) error
// }

// type mockStore struct{}

// var userledgerData = []Ledger{
// 	{
// 		ID:             "zzz",
// 		UserID:         "1",
// 		CurrentBalance: 1300,
// 	},
// 	{
// 		ID:             "kkk",
// 		UserID:         "2",
// 		CurrentBalance: 9812,
// 	},
// }

// func (ms *mockStore) Read(userID string, l *Ledger) error {
// 	foundLedger, err := findByUserID(userID)
// 	if err != nil {
// 		return err
// 	}

// 	*l = *foundLedger
// 	return nil
// }

// func (ms *mockStore) Update(userID string, inputData *Ledger) error {
// 	for idx, u := range userledgerData {
// 		if u.UserID == userID {
// 			userledgerData[idx].CurrentBalance = inputData.CurrentBalance
// 			return nil
// 		}
// 	}

// 	return errors.New("ledger associated with this user not found")
// }

// func findByUserID(userID string) (*Ledger, error) {
// 	l := &Ledger{}
// 	for _, u := range userledgerData {
// 		if u.UserID == userID {
// 			*l = u
// 			break
// 		}
// 	}

// 	if l.ID == "" {
// 		return nil, errors.New("ledger associated with this user not found")
// 	}

// 	return l, nil
// }

// func newStore() *mockStore {
// 	return &mockStore{}
// }
