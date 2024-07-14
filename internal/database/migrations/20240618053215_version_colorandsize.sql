-- +goose Up
-- +goose StatementBegin
ALTER TABLE size ADD COLUMN version INT DEFAULT 1;
ALTER TABLE categories ADD COLUMN version INT DEFAULT 1;
ALTER TABLE colors ADD COLUMN version INT DEFAULT 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
