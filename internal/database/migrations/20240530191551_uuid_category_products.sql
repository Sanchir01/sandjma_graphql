-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Добавление столбца uuid для категорий
ALTER TABLE categories ADD COLUMN IF NOT EXISTS uuid UUID DEFAULT uuid_generate_v4();

-- Добавление столбца uuid для продуктов
ALTER TABLE products ADD COLUMN IF NOT EXISTS uuid UUID DEFAULT uuid_generate_v4();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
