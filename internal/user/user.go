package user

type User struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Service struct {
	store store
}

func (s *Service) Read(id string) (*User, error) {
	user := &User{}
	err := s.store.Read(id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) ReadByUsername(username string) (*User, error) {
	user := &User{}
	err := s.store.ReadByUsername(username, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewService() (*Service, error) {
	nstore, err := newStore()
	if err != nil {
		return nil, err
	}

	return &Service{
		store: nstore,
	}, nil
}
