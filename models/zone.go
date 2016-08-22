package models

import (
	"database/sql"
)

type Zone struct {
	Id          *int64
	Name        *string
	Description *string
	Lat         *float64
	Long        *float64
	Radius      *float64
}

const (
	SaveZoneQuery = `INSERT INTO zones (name,description,lat,long,radius) VALUES ($1,$2,$3,$4,$5)`
)

func (z *Zone) SaveToDb(db *sql.DB) error {
	_, err := db.Exec(SaveZoneQuery, z.Name, z.Description, z.Lat, z.Long, z.Radius)
	return err
}
