package server_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/roku-zeros/mortage-calc/services/calc/internal/models"
	"github.com/roku-zeros/mortage-calc/services/calc/internal/providers"
	storage "github.com/roku-zeros/mortage-calc/services/calc/internal/repository/database"
	"github.com/roku-zeros/mortage-calc/services/calc/internal/server"
)

func TestPing(t *testing.T) {
	storage := storage.NewStorage(context.Background())
	provider := providers.NewMortageProvider(storage)
	s := server.New(provider)

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.Ping)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "Calc service OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestExecute(t *testing.T) {
	storage := storage.NewStorage(context.Background())
	provider := providers.NewMortageProvider(storage)

	s := server.New(provider)

	truePtr := new(bool)
	*truePtr = true
	reqBody := models.Params{
		ObjectCost:     3_000_000,
		InitialPayment: 1_500_000,
		Months:         120,
		Program: &models.Program{
			Military: truePtr,
		},
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/execute", bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.Execute)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response models.Calculation
	json.NewDecoder(rr.Body).Decode(&response) //nolint

}

func TestCache(t *testing.T) {
	storage := storage.NewStorage(context.Background())
	provider := providers.NewMortageProvider(storage)
	s := server.New(provider)

	req, err := http.NewRequest("GET", "/cache", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.Cache)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var calculations []models.Calculation
	json.NewDecoder(rr.Body).Decode(&calculations) //nolint
}
