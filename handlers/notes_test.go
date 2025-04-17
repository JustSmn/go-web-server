package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestController struct{}

func (t *TestController) HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (t *TestController) CreatePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestHomePageReturnsOK(t *testing.T) {
	controller := &TestController{}

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	controller.HomePage(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
}

func TestCreatePageReturnsOK(t *testing.T) {
	controller := &TestController{}

	req := httptest.NewRequest("GET", "/create", nil)
	rr := httptest.NewRecorder()

	controller.CreatePage(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
}
