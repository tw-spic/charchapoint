package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	m "github.com/tw-spic/charchapoint/models"
)

func CreateMessageHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			log.Println("Create message: Empty request body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		decoder := json.NewDecoder(r.Body)
		var msg m.Message
		err := decoder.Decode(&msg)
		if err != nil {
			log.Println("Create message:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if msg.DeviceId == nil || msg.Message == nil {
			log.Println("Create message empty value in request body: ", msg)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now := time.Now()
		msg.MsgTime = &now

		err = msg.SaveToDb(db)
		if err != nil {
			log.Println("Create message save to db :", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
