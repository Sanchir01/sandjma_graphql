-- +goose Up
-- +goose StatementBegin
-- Установка расширения для генерации UUID


-- Удаление старого первичного ключа
ALTER TABLE categories DROP CONSTRAINT IF EXISTS categories_pkey;

-- Замена старого столбца id на новый столбец UUID
ALTER TABLE categories RENAME COLUMN uuid TO id;

-- Установка нового столбца UUID в качестве первичного ключа
ALTER TABLE categories ADD CONSTRAINT categories_pkey PRIMARY KEY (id);

-- Удаление старого столбца id, если требуется
-- ALTER TABLE categories DROP COLUMN old_id;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
