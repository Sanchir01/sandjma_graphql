-- +goose Up
-- +goose StatementBegin
CREATE TABLE colors (
                        id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                        name VARCHAR(255) NOT NULL,
                        slug VARCHAR(255) NOT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        UNIQUE (slug)
);
CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON colors
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
CREATE TABLE product_colors (
                                product_id UUID NOT NULL,
                                color_id UUID NOT NULL,
                                FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
                                FOREIGN KEY (color_id) REFERENCES colors(id) ON DELETE CASCADE,
                                PRIMARY KEY (product_id, color_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
