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
	SaveZoneQuery             = `INSERT INTO zones (name,description,lat,long,radius) VALUES ($1,$2,$3,$4,$5)`
	GetZonesWithinRadiusQuery = `SELECT id,name,description,lat,long,radius FROM zones WHERE ACOS( SIN( RADIANS( lat ) ) * SIN( RADIANS( $1 ) ) + COS( RADIANS( lat ) ) * COS( RADIANS( $1 )) * COS( RADIANS( long ) - RADIANS( $2 )) ) * 6380 < $3;` //6380 is approx radius of earth in km
)

func (z *Zone) SaveToDb(db *sql.DB) error {
	_, err := db.Exec(SaveZoneQuery, z.Name, z.Description, z.Lat, z.Long, z.Radius)
	return err
}

// GetZonesWithinRadiusFrom returns all the zones within the given radius from given point.
// Input is latitude and longitude of the point in radian and radius in km.
func GetZonesWithinRadiusFrom(lat, long, radius float64, db *sql.DB) ([]Zone, error) {
	zones := make([]Zone, 0)
	rows, err := db.Query(GetZonesWithinRadiusQuery, lat, long, radius)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return zones, err
	}
	for rows.Next() {
		z := Zone{}
		err := rows.Scan(&z.Id, &z.Name, &z.Description, &z.Lat, &z.Long, &z.Radius)
		if err != nil {
			return zones, err
		}
		zones = append(zones, z)
	}
	return zones, nil
}
