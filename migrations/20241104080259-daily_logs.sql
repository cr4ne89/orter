
-- +migrate Up
CREATE TABLE daily_logs (
    day_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(user_id),
    date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +migrate Down
DROP TABLE daily_logs;
