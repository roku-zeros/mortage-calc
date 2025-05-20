package providers

import (
	"context"

	"github.com/roku-zeros/mortage-calc/services/calc/internal/models"

	mortagecalc "github.com/roku-zeros/mortage-calc/lib/mortgagecalc"
	pkgerrors "github.com/roku-zeros/mortage-calc/services/calc/internal/errors"
)

func (p *MortageProvider) CreateMortage(ctx context.Context, params models.Params) (models.Calculation, error) {
	if params.Program == nil {
		return models.Calculation{}, pkgerrors.ErrNoProgram
	}
	if float32(params.InitialPayment) < float32(params.ObjectCost)*0.2 {
		return models.Calculation{}, pkgerrors.ErrBadInitialPayment
	}

	var rate int
	count := 0
	if params.Program.Salary != nil && *params.Program.Salary {
		count++
		rate = 8
	}
	if params.Program.Military != nil && *params.Program.Military {
		count++
		rate = 9
	}
	if params.Program.Base != nil && *params.Program.Base {
		count++
		rate = 10
	}
	if count == 0 {
		return models.Calculation{}, pkgerrors.ErrNoProgram
	} else if count > 1 {
		return models.Calculation{}, pkgerrors.ErrMoreThanOneProgram
	}

	monthlyPayment, overpayment, loanSum, lastPaymentDate := mortagecalc.CalculateMortgage(params.ObjectCost, params.InitialPayment, params.Months, rate)
	aggregate := models.Aggregate{
		Rate:            rate,
		LoanSum:         loanSum,
		MonthlyPayment:  monthlyPayment,
		Overpayment:     overpayment,
		LastPaymentDate: lastPaymentDate.Format("2006-01-02"),
	}

	calculation := models.Calculation{
		Params:     params,
		Program:    *params.Program,
		Aggregates: aggregate,
	}

	id := p.storage.CreateMortage(ctx, calculation)
	calculation.ID = id

	return calculation, nil
}

func (p *MortageProvider) GetAllMortages(ctx context.Context) ([]models.Calculation, error) {
	calculations := p.storage.GetAllMortages(ctx)
	if len(calculations) == 0 {
		return nil, pkgerrors.ErrEmptyCache
	}
	return calculations, nil
}
