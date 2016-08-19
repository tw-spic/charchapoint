package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	h "github.com/tw-spic/charchapoint/handlers"

	"github.com/stretchr/testify/assert"
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

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler())
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusCreated, rr.Code)
}

func TestCreateZoneFailsForEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler())
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateZoneFailsForMalformedJSON(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader("}{"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler())
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

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler())
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

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler())
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

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler())
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

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateZoneHandler())
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}
