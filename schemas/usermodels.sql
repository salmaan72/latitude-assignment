CREATE TYPE status AS ENUM ('approved', 'pending');
CREATE TABLE IF NOT EXISTS User_models(
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    phone TEXT NOT NULL,
    status status DEFAULT 'pending',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);
