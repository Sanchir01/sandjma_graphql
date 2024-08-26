-- +goose Up
-- +goose StatementBegin
ALTER TABLE colors ADD COLUMN  css_variables VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
