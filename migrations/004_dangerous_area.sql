-- +goose Up
-- +goose StatementBegin
CREATE TABLE dangerous_area(
    id SERIAL PRIMARY KEY,
    date_created TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    point_a_id INT UNIQUE REFERENCES location_points (id) ON DELETE CASCADE,
    point_b_id INT UNIQUE REFERENCES location_points (id) ON DELETE CASCADE

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE dangerous_area
-- +goose StatementEnd
