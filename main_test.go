package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestRootHandler tests the root endpoint
func TestRootHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	// Act
	rootHandler(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: expected %v, received %v",
			http.StatusOK, status)
	}
	expected := "The server is running"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned a wrong body: expected %v, received %v",
			expected, rr.Body.String())
	}
}

// TestReportHandler tests the /report
func TestReportHandler(t *testing.T) {
	// Arrange
	req := httptest.NewRequest("GET", "/report", nil)
	rr := httptest.NewRecorder()

	// Act
	reportHandler(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: expected %v, received %v",
			http.StatusOK, status)
	}
	expected := "Report generated successfully!"
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned a wrong body: expected %v, received %v",
			expected, rr.Body.String())
	}
}
