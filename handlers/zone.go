package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	m "github.com/tw-spic/charchapoint/models"
)

func CreateZoneHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			log.Println("Create zone: Empty request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		decoder := json.NewDecoder(r.Body)
		var z m.Zone
		err := decoder.Decode(&z)
		if err != nil {
			log.Println("Create zone:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if z.Name == nil || z.Lat == nil || z.Long == nil || z.Radius == nil {
			log.Println("Create zone empty value in request body: ", z)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = z.SaveToDb(db)
		if err != nil {
			log.Println("Create zone save to db :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetZoneHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lat, err := strconv.ParseFloat(r.FormValue("lat"), 64)
		if err != nil {
			log.Println("Get zone invalid lat: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		long, err := strconv.ParseFloat(r.FormValue("long"), 64)
		if err != nil {
			log.Println("Get zone invalid long: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		radius, err := strconv.ParseFloat(r.FormValue("radius"), 64)
		if err != nil {
			log.Println("Get zone invalid radius: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		zones, err := m.GetZonesWithinRadiusFrom(lat, long, radius, db)
		if err != nil {
			log.Println("Get zone error fetching data from db:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(zones)
		if err != nil {
			log.Println("Get zone error converting to JSON:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}
