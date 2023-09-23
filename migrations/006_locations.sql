-- +goose Up
-- +goose StatementBegin
CREATE TABLE locations(
    id SERIAL PRIMARY KEY,
    name TEXT
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE locations
-- +goose StatementEnd
