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

func (rm *ReportModel) Insert(location string, user_id int) error {
	locstmt := `
		INSERT INTO location_points (lat, lng)
		VALUES (52.31, 53.32);
	`
	// parts := strings.Split(location, ",")
	// latitude, _ := strconv.ParseFloat(parts[0], 64)
	// longitude, _ := strconv.ParseFloat(parts[1], 64)
	_, err := rm.DB.Exec(locstmt)
	if err != nil {
		return err
	}

	stmt := `
		INSERT INTO reports (user_id, report_type, report_description)
		VALUES ($1, $2, $3) RETURNING id
	`
	_, err = rm.DB.Exec(stmt, user_id, "bang", "bangarang")
	if err != nil {
		return err
	}
	return nil
}
