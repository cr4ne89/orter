
-- +migrate Up
CREATE TABLE daily_log_items (
    daily_log_id UUID REFERENCES daily_logs(day_id),
    item_id UUID REFERENCES items(item_id),
    contents JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (daily_log_id, item_id)
);

-- +migrate Down
DROP TABLE daily_log_items;
