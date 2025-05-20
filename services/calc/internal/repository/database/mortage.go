package storage

import (
	"context"

	"github.com/roku-zeros/mortage-calc/services/calc/internal/models"
)

func (s *Storage) CreateMortage(ctx context.Context, calculation models.Calculation) (id int) {
	s.db.Set(calculation)
	return s.db.GetCurrID()
}

func (s *Storage) GetAllMortages(ctx context.Context) []models.Calculation {
	id := s.db.GetCurrID()
	calculations := make([]models.Calculation, id+1)
	for i := range id + 1 {
		val, ok := s.db.Get(i)
		if !ok {
			continue
		}
		calc := val.(models.Calculation)
		calc.ID = i
		calculations[i] = calc
	}
	return calculations
}
