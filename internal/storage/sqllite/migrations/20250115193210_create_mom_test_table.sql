-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    companies (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(60) NOT NULL,
        description TEXT NOT NULL,
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE companies;

-- +goose StatementEnd