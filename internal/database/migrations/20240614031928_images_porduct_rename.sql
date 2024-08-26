-- +goose Up
-- +goose StatementBegin
ALTER TABLE products RENAME COLUMN password TO images;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
