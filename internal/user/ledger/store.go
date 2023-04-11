package ledger

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/salmaan72/latitude-assignment/internal/clients/datastore"
)

type store interface {
	Create(ctx context.Context, ledger *Ledger) error
	ReadByUserID(ctx context.Context, userID uuid.UUID, ledger *Ledger) error
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
	now := time.Now()
	lm.CreatedAt = &now
	lm.UpdatedAt = &now

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

func (cm *CardModel) prepare(lm *LedgerModel, card *Card) {
	cm.ID = uuid.New()
	now := time.Now()
	cm.CreatedAt = &now
	cm.UpdatedAt = &now

	cm.LedgerID = lm.ID
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
		cm.prepare(lm, &copy)

		cardModelList = append(cardModelList, *cm)
	}

	err = nls.db.Create(cardModelList).Error
	if err != nil {
		return err
	}

	ledger.fetchFromModelsBasic(lm)

	return nil
}

func (nls *newLedgerStore) ReadByUserID(ctx context.Context, userID uuid.UUID, ledger *Ledger) error {
	ledgerModel := &LedgerModel{}
	err := nls.db.Where("user_id=?", userID.String()).Find(ledgerModel).Error
	if err != nil {
		return err
	}

	cardModel := &CardModel{}
	err = nls.db.Where("ledger_id=?", ledgerModel.ID.String()).Find(cardModel).Error
	if err != nil {
		return err
	}

	modelToStruct(ledger, ledgerModel, cardModel)
	return nil
}

func modelToStruct(l *Ledger, ledgerModel *LedgerModel, cardModel *CardModel) {
	l.ID = ledgerModel.ID
	l.UserID = ledgerModel.UserID
	l.AccountNumber = ledgerModel.AccountNumber
	l.CurrentBalance = ledgerModel.CurrentBalance
	l.Cards = []Card{
		{
			Type:   cardType(cardModel.Type),
			Number: cardModel.Number,
			CVV:    cardModel.CVV,
			Expiry: cardModel.Expiry,
		},
	}
}

func newStore(db *datastore.Client) (*newLedgerStore, error) {
	return &newLedgerStore{
		db: db,
	}, nil
}
