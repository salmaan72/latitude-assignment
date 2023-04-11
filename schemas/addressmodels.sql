CREATE TABLE IF NOT EXISTS Address_models(
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    line1 TEXT,
    city TEXT,
    state TEXT,
    country TEXT,
    pincode TEXT,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);
