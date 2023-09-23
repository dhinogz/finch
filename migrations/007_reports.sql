-- +goose Up
-- +goose StatementBegin
CREATE TABLE reports (
    id SERIAL PRIMARY KEY,
    user_id SERIAL UNIQUE REFERENCES users (id) ON DELETE CASCADE,
    report_type TEXT,
    location_id SERIAL UNIQUE REFERENCES locations (id) ON DELETE CASCADE,
    report_description TEXT,
    location_point SERIAL UNIQUE REFERENCES location_points (id) ON DELETE CASCADE
    

)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reports
-- +goose StatementEnd
