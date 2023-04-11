CREATE TABLE IF NOT EXISTS card_models(
    id TEXT PRIMARY KEY,
    ledger_id TEXT NOT NULL,
    type TEXT NOT NULL,
    number TEXT NOT NULL,
    cvv TEXT NOT NULL,
    expiry timestamptz NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);
