package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(healthCheck)
	handler.ServeHTTP(rr, req)

	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}

	expected_body := "doodle-guessr server is running\n"
	if rr.Body.String() != expected_body {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected_body)
	}
}

func TestHealthCheckStatusMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(healthCheck)
	handler.ServeHTTP(rr, req)

	expectedStatus := http.StatusMethodNotAllowed
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}
}
