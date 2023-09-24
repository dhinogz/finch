package data

import (
	"database/sql"
)

type Report struct {
	ID          int
	Type        string
	Description string
}

type ReportModel struct {
	DB *sql.DB
}

func (rm *ReportModel) Insert(user_id int, reportType string, description string) error {
	locstmt := `
		INSERT INTO location_points (lat, lng)
		VALUES (37.782, -122.445);
	`

	_, err := rm.DB.Exec(locstmt)
	if err != nil {
		return err
	}

	stmt := `
		INSERT INTO reports (user_id, report_type, report_description)
		VALUES ($1, $2, $3) RETURNING id
	`
	_, err = rm.DB.Exec(stmt, user_id, reportType, description)
	if err != nil {
		return err
	}
	return nil
}
