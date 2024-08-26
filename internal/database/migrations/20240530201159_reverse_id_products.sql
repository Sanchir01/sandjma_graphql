-- +goose Up
-- +goose StatementBegin
ALTER TABLE products DROP CONSTRAINT IF EXISTS products_pkey;

-- Установка нового столбца UUID в качестве первичного ключа
ALTER TABLE products ADD CONSTRAINT products_new_pkey PRIMARY KEY (uuid);

-- Удаление старого столбца id
ALTER TABLE products DROP COLUMN id;

-- Переименование столбца new_uuid в id
ALTER TABLE products RENAME COLUMN uuid TO id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
