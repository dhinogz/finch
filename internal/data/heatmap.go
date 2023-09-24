package data

func (mm *MapModel) GetDangerousArea() ([]*Place, error) {
	stmt := `SELECT l.lat, l.lng
		FROM location_points l
	`

	rows, err := mm.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ps := []*Place{}

	for rows.Next() {
		p := &Place{}
		err = rows.Scan(&p.Lat, &p.Lng)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}

	if rows.Err(); err != nil {
		return nil, err
	}

	return ps, nil
}
