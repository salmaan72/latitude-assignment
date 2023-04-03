package ledger

type Service struct {
	store store
}

type Ledger struct {
	ID             string `json:"id,omitempty"`
	UserID         string `json:"userID,omitempty"`
	CurrentBalance int64  `json:"currentBalance,omitempty"`
}

func (s *Service) Read(userID string) (*Ledger, error) {
	l := &Ledger{}
	err := s.store.Read(userID, l)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func NewService() *Service {
	ns := newStore()
	return &Service{
		store: ns,
	}
}
