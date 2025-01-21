-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    mom_tests (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        company_id INTEGER NOT NULL,
        hypothesis TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE mom_tests;

-- +goose StatementEnd