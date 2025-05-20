package providers

import (
	"context"
	"testing"

	"github.com/roku-zeros/mortage-calc/services/calc/internal/models"

	pkgerrors "github.com/roku-zeros/mortage-calc/services/calc/internal/errors"
	storage "github.com/roku-zeros/mortage-calc/services/calc/internal/repository/database"
)

func TestCreateMortage(t *testing.T) {
	ctx := context.Background()
	storage := storage.NewStorage(context.Background())
	provider := &MortageProvider{storage: storage}
	truePtr := new(bool)
	*truePtr = true

	tests := []struct {
		name        string
		params      models.Params
		expectError error
	}{
		{
			name: "No program",
			params: models.Params{
				Program: nil,
			},
			expectError: pkgerrors.ErrNoProgram,
		},
		{
			name: "No program",
			params: models.Params{
				Program: &models.Program{
					Base:     new(bool),
					Salary:   new(bool),
					Military: new(bool),
				},
			},
			expectError: pkgerrors.ErrNoProgram,
		},
		{
			name: "More than one program",
			params: models.Params{
				Program: &models.Program{
					Base:     truePtr,
					Salary:   truePtr,
					Military: truePtr,
				},
			},
			expectError: pkgerrors.ErrMoreThanOneProgram,
		},
		{
			name: "The initial payment should be more",
			params: models.Params{
				ObjectCost:     5_000_000,
				InitialPayment: 900_000,
				Program: &models.Program{
					Base: truePtr,
				},
			},
			expectError: pkgerrors.ErrBadInitialPayment,
		},
		{
			name: "Valid program",
			params: models.Params{
				Program: &models.Program{
					Base: truePtr,
				},
			},
			expectError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := provider.CreateMortage(ctx, tt.params)
			if tt.expectError != nil {
				if err != tt.expectError {
					t.Errorf("expected error %v, got %v", tt.expectError, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				} else {
					calculation := models.Calculation{}
					storage.CreateMortage(context.Background(), calculation)
					if len(storage.GetAllMortages(context.Background())) == 0 {
						t.Error("expected non-zero ID after saving calculation")
					}
				}
			}
		})
	}
}

func TestGetAllMortages(t *testing.T) {
	ctx := context.Background()
	storage := storage.NewStorage(context.Background())
	provider := &MortageProvider{storage: storage}
	truePtr := new(bool)
	*truePtr = true

	tests := []struct {
		name        string
		storageData []models.Calculation
		expectError error
	}{
		{
			name:        "Empty cache",
			storageData: nil,
			expectError: pkgerrors.ErrEmptyCache,
		},
		{
			name: "Non-empty cache",
			storageData: []models.Calculation{
				{
					ID: 0,
					Params: models.Params{
						ObjectCost:     3_000_000,
						InitialPayment: 1_500_000,
						Months:         120,
					},
					Program: models.Program{
						Military: truePtr,
					},
				},
			},
			expectError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.storageData != nil {
				for _, calc := range tt.storageData {
					storage.CreateMortage(context.Background(), calc)
				}
			}

			result, err := provider.GetAllMortages(ctx)

			if tt.expectError != nil {
				if err != tt.expectError {
					t.Errorf("expected error %v, got %v", tt.expectError, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				} else if len(result) != len(tt.storageData) {
					t.Errorf("expected %d results, got %d", len(tt.storageData), len(result))
				}
				for i, calc := range result {
					if calc != tt.storageData[i] {
						t.Errorf("expected result %v, got %v", tt.storageData[i], calc)
					}
				}
			}
		})
	}
}
