package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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
