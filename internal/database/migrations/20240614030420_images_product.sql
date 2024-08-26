-- +goose Up
-- +goose StatementBegin
ALTER TABLE products ADD COLUMN password TEXT[] DEFAULT ARRAY[]::TEXT[];
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
