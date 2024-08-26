-- +goose Up
-- +goose StatementBegin
UPDATE products SET category_id = (SELECT id FROM categories LIMIT 1);
ALTER TABLE products ADD CONSTRAINT products_category_id_fkey FOREIGN KEY (category_id) REFERENCES categories(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
