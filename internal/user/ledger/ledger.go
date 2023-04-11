package ledger

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type Service struct {
	store store
}

type cardType string

var (
	CardTypeCredit = cardType("credit")
	CardTypeDebit  = cardType("debit")
)

type Ledger struct {
	ID             uuid.UUID  `json:"id,omitempty"`
	UserID         uuid.UUID  `json:"userID,omitempty"`
	AccountNumber  string     `json:"accountNumber,omitempty"`
	CurrentBalance int64      `json:"currentBalance,omitempty"`
	Cards          []Card     `json:"cards,omitempty"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
}

type Card struct {
	Type   cardType   `json:"type,omitempty"`
	Number string     `json:"number,omitempty"`
	CVV    string     `json:"cvv,omitempty"`
	Expiry *time.Time `json:"expiry,omitempty"`
}

func (l *Ledger) fetchFromModelsBasic(lm *LedgerModel) {
	l.ID = lm.ID
	l.CreatedAt = lm.CreatedAt
	l.UpdatedAt = lm.UpdatedAt
}

func (s *Service) Create(ctx context.Context, ledger *Ledger) (*Ledger, error) {
	now := time.Now()
	expiry := now.AddDate(1, 0, 0)

	ledger.AccountNumber = fmt.Sprintf("%016d", rand.Int63n(1e16))
	ledger.CurrentBalance = int64(0)
	ledger.Cards = []Card{
		{
			Type:   CardTypeDebit,
			Number: fmt.Sprintf("%012d", rand.Int63n(1e12)),
			CVV:    fmt.Sprintf("%03d", rand.Int63n(1e3)),
			Expiry: &expiry,
		},
	}

	err := s.store.Create(ctx, ledger)
	if err != nil {
		return nil, err
	}

	return ledger, nil
}

func NewService(db *datastore.Client) (*Service, error) {
	ns, err := newStore(db)
	if err != nil {
		return nil, err
	}
	return &Service{
		store: ns,
	}, nil
}
