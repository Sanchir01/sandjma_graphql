-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN password VARCHAR(255) NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
