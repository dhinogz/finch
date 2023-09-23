-- +goose Up
-- +goose StatementBegin
CREATE TABLE routes(
    id SERIAL PRIMARY KEY,
    user_id SERIAL UNIQUE REFERENCES users (id) ON DELETE CASCADE,
    start_id SERIAL UNIQUE REFERENCES location_points (id) ON DELETE CASCADE,
    destination_id SERIAL UNIQUE REFERENCES location_points (id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE routes
-- +goose StatementEnd
