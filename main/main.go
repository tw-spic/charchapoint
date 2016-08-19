package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	c "github.com/tw-spic/charchapoint/config"
	h "github.com/tw-spic/charchapoint/handlers"

	gh "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	f, err := os.OpenFile("CharchaPoint.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	conf, err := c.ReadFromFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	connString := fmt.Sprintf("user=%s password=%s dbname=smart_scale sslmode=disable",
		conf.DBUsername, conf.DBPassword)
	db, err := sql.Open("postgres", connString)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", h.ServeIndexPage())
	r.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), gh.LoggingHandler(f, r)))
}
