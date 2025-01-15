-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    leads (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        linked_in_url TEXT NOT NULL,
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE leads;

-- +goose StatementEnd