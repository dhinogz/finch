package data

import (
	"database/sql"
	"errors"

	"googlemaps.github.io/maps"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	User   UserModel
	Report ReportModel
	Map    MapModel
}

func NewModel(db *sql.DB, gmaps *maps.Client) Models {
	return Models{
		User:   UserModel{DB: db},
		Report: ReportModel{DB: db},
		Map: MapModel{
			DB:    db,
			GMaps: gmaps,
		},
	}
}
