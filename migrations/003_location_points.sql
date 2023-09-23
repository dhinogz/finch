-- +goose Up
-- +goose StatementBegin
CREATE TABLE location_points(
    id SERIAL PRIMARY KEY,
    lat DOUBLE PRECISION NOT NULL,
    lng DOUBLE PRECISION NOT NULL,
    name TEXT
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE location_points
-- +goose StatementEnd
