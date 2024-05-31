-- +goose Up
-- +goose StatementBegin
-- Удаление столбца category_id из таблицы products
ALTER TABLE products DROP COLUMN IF EXISTS category_id;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
