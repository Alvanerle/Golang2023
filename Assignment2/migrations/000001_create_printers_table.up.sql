CREATE TABLE IF NOT EXISTS printers (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255),
    is_color BOOLEAN,
    ip_address VARCHAR(255),
    status VARCHAR(255) NOT NULL,
    supported_paper_sizes TEXT[] NOT NULL,
    description TEXT NOT NULL,
    battery_left INTEGER NOT NULL DEFAULT -1
);