CREATE TABLE IF NOT EXISTS Ledger_models(
    id TEXT PRIMARY KEY,
    user_id TEXT,
    account_number TEXT,
    current_balance FLOAT,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);
