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

func TestCreateMessageSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"DeviceId":"xxxxx",
			"Message":"Hello world"
		}`))
	if err != nil {
		t.Fatal(err)
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(`INSERT INTO messages\(id, device_id, message, msg_time\) VALUES\(gen_random_uuid\(\), \$1, \$2, \$3\)`).WithArgs("xxxxx", "Hello world", sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateMessageHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusCreated, rr.Code)
}

func TestCreateMessageFailsForEmptyBody(t *testing.T) {
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
	handler := http.HandlerFunc(h.CreateMessageHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateMessageFailsForMalformedJSON(t *testing.T) {
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
	handler := http.HandlerFunc(h.CreateMessageHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateMessageFailsIfDeviceIdIsEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"Message":"Hello"
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
	handler := http.HandlerFunc(h.CreateMessageHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateMessageFailsIfMessageIsEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"DeviceId":"xxxx"
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
	handler := http.HandlerFunc(h.CreateMessageHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestCreateMessageFailsIfDbFails(t *testing.T) {
	req, err := http.NewRequest("POST", "/", strings.NewReader(`{
			"DeviceId":"xxxx",
			"Message":"Hello world"
		}`))
	if err != nil {
		t.Fatal(err)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mock.ExpectExec(`INSERT INTO messages\(id, device_id, message, msg_time\) VALUES\(gen_random_uuid\(\), \$1, \$2, \$3\)`).WithArgs("xxxxx", "Hello world", sqlmock.AnyArg()).WillReturnError(fmt.Errorf("some error"))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.CreateMessageHandler(db))
	handler.ServeHTTP(rr, req)

	assert := assert.New(t)
	assert.Equal(http.StatusInternalServerError, rr.Code)
	err = mock.ExpectationsWereMet()
}
