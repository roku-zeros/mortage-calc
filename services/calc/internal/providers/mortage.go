package providers

import (
	"context"
	mortagecalc "mortage-calc/lib/mortgagecalc"
	pkgerrors "mortage-calc/services/calc/internal/errors"
	"mortage-calc/services/calc/internal/models"
)

func (p *MortageProvider) CreateMortage(ctx context.Context, params models.Params) error {
	if params.Program == nil {
		return pkgerrors.ErrNoProgram
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
		return pkgerrors.ErrNoProgram
	} else if count > 1 {
		return pkgerrors.ErrMoreThanOneProgram
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

	p.storage.CreateMortage(ctx, calculation)
	return nil
}

func (p *MortageProvider) GetAllMortages(ctx context.Context) ([]models.Calculation, error) {
	calculations := p.storage.GetAllMortages(ctx)
	if len(calculations) == 0 {
		return nil, pkgerrors.ErrEmptyCache
	}
	return calculations, nil
}
