-- +goose Up
-- +goose StatementBegin
-- Шаг 1: Добавление новой колонки category_id типа UUID
ALTER TABLE products ADD COLUMN category_id UUID;

-- Шаг 2: Установка значения category_id для всех существующих записей



-- Шаг 2: Установка значения по умолчанию для всех существующих записей

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
