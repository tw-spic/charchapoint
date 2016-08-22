package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	h "github.com/tw-spic/charchapoint/handlers"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreateZoneSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"Name":"Zone1",
			"Description": "Awesome zone",
			"Lat":10.532,
			"Long":11.324,
			"Radius":10
		}`))
	if err != nil {
		t.Fatal(err)
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(`INSERT INTO zones \(name,description,lat,long,radius\) VALUES \(\$1,\$2,\$3,\$4,\$5\)`).WithArgs("Zone1", "Awesome zone", 10.532, 11.324, 10.0).WillReturnResult(sqlmock.NewResult(1, 1))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusCreated, rr.Code)
}

func TestCreateZoneFailsForEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateZoneFailsForMalformedJSON(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader("}{"))
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateZoneFailsIfNameIsEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"Description": "Awesome zone",
			"Lat":10.532,
			"Long":11.324,
			"Radius":10
		}`))
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateZoneFailsIfLatIsEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"Name":"Zone1",
			"Description": "Awesome zone",
			"Long":11.324,
			"Radius":10
		}`))
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateZoneFailsIfLongIsEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"Name":"Zone1",
			"Description": "Awesome zone",
			"Lat":10.532,
			"Radius":10
		}`))
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateZoneFailsIfRadiusIsEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"Name":"Zone1",
			"Description": "Awesome zone",
			"Lat":10.532,
			"Long":11.324
		}`))
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateZoneFailsIfDbFails(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"Name":"Zone1",
			"Description": "Awesome zone",
			"Lat":10.532,
			"Long":11.324,
			"Radius":10
		}`))
	if err != nil {
		t.Fatal(err)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(`INSERT INTO zones \(name,description,lat,long,radius\) VALUES \(\$1,\$2,\$3,\$4,\$5\)`).WithArgs("Zone1", "Awesome zone", 10.532, 11.324, 10.0).WillReturnError(fmt.Errorf("some error"))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusInternalServerError, rr.Code)
	err = mock.ExpectationsWereMet()
}
