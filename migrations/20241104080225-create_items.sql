
-- +migrate Up
CREATE TABLE items (
    item_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(user_id),
    name TEXT NOT NULL,
    description TEXT,
    item_type TEXT CHECK (item_type IN ('text', 'checkbox', 'time')) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- +migrate Down
DROP TABLE items;
