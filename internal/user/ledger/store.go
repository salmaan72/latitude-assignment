package ledger

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type store interface {
	Create(ctx context.Context, ledger *Ledger) error
}

type newLedgerStore struct {
	db *datastore.Client
}

type LedgerModel struct {
	datastore.Model
	UserID         uuid.UUID
	AccountNumber  string
	CurrentBalance int64
}

func (lm *LedgerModel) prepare(ledger *Ledger) {
	lm.ID = uuid.New()
	lm.CreatedAt = time.Now()
	lm.UpdatedAt = time.Now()

	lm.UserID = ledger.UserID
	lm.AccountNumber = ledger.AccountNumber
	lm.CurrentBalance = ledger.CurrentBalance
}

type CardModel struct {
	datastore.Model
	LedgerID uuid.UUID
	Type     string
	Number   string
	CVV      string
	Expiry   *time.Time
}

func (cm *CardModel) prepare(ledgerID uuid.UUID, card *Card) {
	cm.ID = uuid.New()
	cm.CreatedAt = time.Now()
	cm.UpdatedAt = time.Now()

	cm.LedgerID = ledgerID
	cm.Type = string(card.Type)
	cm.Number = card.Number
	cm.CVV = card.CVV
	cm.Expiry = card.Expiry
}

func (nls *newLedgerStore) Create(ctx context.Context, ledger *Ledger) error {
	lm := &LedgerModel{}
	lm.prepare(ledger)
	err := nls.db.Create(lm).Error
	if err != nil {
		return err
	}

	cardModelList := make([]CardModel, 0, len(ledger.Cards))
	for _, c := range ledger.Cards {
		cm := &CardModel{}
		copy := c
		cm.prepare(ledger.ID, &copy)

		cardModelList = append(cardModelList, *cm)
	}

	err = nls.db.Create(cardModelList).Error
	if err != nil {
		return err
	}

	ledger.fetchFromModelsBasic(lm)

	return nil
}

func newStore(db *datastore.Client) (*newLedgerStore, error) {
	return &newLedgerStore{
		db: db,
	}, nil
}
