package data

import "database/sql"

type Report struct {
	ID          int
	Type        string
	Description string
}

type ReportModel struct {
	DB *sql.DB
}

func (rm *ReportModel) Insert(reportType, description string) error {
	r := Report{}
	stmt := `
		INSERT INTO reports (report_type, report_description)
		VALUES ($1, $2) RETURNING id
	`
	row := rm.DB.QueryRow(stmt, reportType, description)
	err := row.Scan(&r.ID)
	if err != nil {
		return err
	}
	return nil
}
