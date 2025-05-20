package storage

import (
	"context"
	"testing"

	"github.com/roku-zeros/mortage-calc/services/calc/internal/models"
)

func TestMortagesStorage(t *testing.T) {
	storage := NewStorage(context.Background())

	calculation1 := models.Calculation{
		Params: models.Params{
			ObjectCost:     100000,
			InitialPayment: 20000,
			Months:         240,
		},
	}
	calculation2 := models.Calculation{
		Params: models.Params{
			ObjectCost:     150000,
			InitialPayment: 30000,
			Months:         180,
		},
	}

	storage.CreateMortage(context.Background(), calculation1)
	storage.CreateMortage(context.Background(), calculation2)

	calcs := storage.GetAllMortages(context.Background())

	if len(calcs) != 2 {
		t.Errorf("expected 2 mortgages, got %d", len(calcs))
	}

	if calcs[0].Params.ObjectCost != calculation1.Params.ObjectCost {
		t.Errorf("first calculation object cost mismatch: expected %d, got %d", calculation1.Params.ObjectCost, calcs[0].Params.ObjectCost)
	}

	if calcs[1].Params.ObjectCost != calculation2.Params.ObjectCost {
		t.Errorf("second calculation object cost mismatch: expected %d, got %d", calculation2.Params.ObjectCost, calcs[1].Params.ObjectCost)
	}
}
