-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       name VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       phone VARCHAR(15) NOT NULL UNIQUE NOT NULL,
                       email TEXT NOT NULL UNIQUE NOT NULL,
                       avatar_path TEXT DEFAULT ('https://www.google.com/url?sa=i&url=https%3A%2F%2Fru.dreamstime.com%2F%25D0%25B0%25D0%25BD%25D0%25BE%25D0%25BD%25D0%25B8%25D0%25BC%25D0%25BD%25D1%258B%25D0%25B9-%25D0%25B3%25D0%25B5%25D0%25BD%25D0%25B4%25D0%25B5%25D1%2580%25D0%25BD%25D0%25BE-%25D0%25BD%25D0%25B5%25D0%25B9%25D1%2582%25D1%2580%25D0%25B0%25D0%25BB%25D1%258C%25D0%25BD%25D1%258B%25D0%25B9-%25D0%25B0%25D0%25B2%25D0%25B0%25D1%2582%25D0%25B0%25D1%2580-%25D1%2581%25D0%25B8%25D0%25BB%25D1%2583%25D1%258D%25D1%2582-%25D0%25B3%25D0%25BE%25D0%25BB%25D0%25BE%25D0%25B2%25D1%258B-%25D0%25B8%25D0%25BD%25D0%25BA%25D0%25BE%25D0%25B3%25D0%25BD%25D0%25B8%25D1%2582%25D0%25BE-image227531366&psig=AOvVaw3ne_h-rK4XxvxDPC-lzVwc&ust=1717639583815000&source=images&cd=vfe&opi=89978449&ved=0CBIQjRxqFwoTCLi2oZyww4YDFQAAAAAdAAAAABAE'),
                       role VARCHAR(50) NOT NULL
);
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
