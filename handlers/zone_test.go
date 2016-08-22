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

////////////////////////////////// Create Zone /////////////////////////////////////////////////
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

///////////////////////////////////////////////////////////////////////////////////////////////

///////////////////////////// Get zone ////////////////////////////////////////////////////////
func TestGetZoneSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/zone?lat=10.123&long=11.234&radius=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "lat", "long", "radius"}).
		AddRow(1, "one", "Oneee", 10.111, 11.222, 10).
		AddRow(2, "two", "Twooo", 10.222, 11.333, 10)

	mock.ExpectQuery(`SELECT id,name,description,lat,long,radius FROM zones WHERE ACOS\( SIN\( RADIANS\( lat \) \) \* SIN\( RADIANS\( \$1 \) \) \+ COS\( RADIANS\( lat \) \) \* COS\( RADIANS\( \$1 \)\) \* COS\( RADIANS\( long \) - RADIANS\( \$2 \)\) \) \* 6380 < \$3;`).WithArgs(10.123, 11.234, 1.0).WillReturnRows(rows)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, rr.Code)
	assert.Equal(`[{"Id":1,"Name":"one","Description":"Oneee","Lat":10.111,"Long":11.222,"Radius":10},{"Id":2,"Name":"two","Description":"Twooo","Lat":10.222,"Long":11.333,"Radius":10}]`, rr.Body.String())
}

func TestGetZoneFailsIfLatIsEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/zone?long=10.123&radius=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestGetZoneFailsIfLongIsEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/zone?lat=10.123&radius=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestGetZoneFailsIfRadiusIsEmpty(t *testing.T) {
	req, err := http.NewRequest("GET", "/zone?lat=11.234&long=10.123", nil)
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestGetZoneFailsIfLatIsInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/zone?lat=foo&long=10.123&radius=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestGetZoneFailsIfLongIsInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/zone?lat=11.234&long=foo&radius=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestGetZoneFailsIfRadiusIsInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/zone?lat=11.234&long=10.123&radius=foo", nil)
	if err != nil {
		t.Fatal(err)
	}
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestGetZoneFailsIfDbFails(t *testing.T) {
	req, err := http.NewRequest("GET", "/zone?lat=11.234&long=10.123&radius=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectQuery(`SELECT id,name,description,lat,long,radius FROM zones WHERE ACOS\( SIN\( RADIANS\( lat \) \) \* SIN\( RADIANS\( \$1 \) \) \+ COS\( RADIANS\( lat \) \) \* COS\( RADIANS\( \$1 \)\) \* COS\( RADIANS\( long \) - RADIANS\( \$2 \)\) \) \* 6380 < \$3;`).WithArgs(10.123, 11.234, 1.0).WillReturnError(fmt.Errorf("some error"))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetZoneHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusInternalServerError, rr.Code)
	err = mock.ExpectationsWereMet()
}
