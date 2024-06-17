-- +goose Up
-- +goose StatementBegin
CREATE TABLE size (
                        id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                        name VARCHAR(255) NOT NULL,
                        slug VARCHAR(255) NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        UNIQUE (slug)
);

CREATE TABLE product_sizes (
                               product_id UUID REFERENCES products(id) ON DELETE CASCADE,
                               size_id UUID REFERENCES size(id) ON DELETE CASCADE,
                               PRIMARY KEY (product_id, size_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
