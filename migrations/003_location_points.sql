-- +goose Up
-- +goose StatementBegin
CREATE TABLE location_points(
    id SERIAL PRIMARY KEY,
    lat INT NOT NULL,
    lng INT NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE location_points
-- +goose StatementEnd
