CREATE TABLE IF NOT EXISTS verifier_models(
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    password TEXT NOT NULL,
    email_otp TEXT,
    is_email_verified BOOLEAN,
    phone_otp TEXT,
    is_phone_verified BOOLEAN,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);
