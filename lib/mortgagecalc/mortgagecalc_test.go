package mortagecalc

import (
	"testing"
	"time"
)

func TestCalculateMortgage(t *testing.T) {
	tests := []struct {
		objectCost             int
		initialPayment         int
		months                 int
		rate                   int
		expectedMonthlyPayment int
		expectedOverpayment    int
		expectedLoanSum        int
	}{
		{5_000_000, 1_000_000, 240, 8, 33_458, 4_029_920, 4_000_000},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			monthlyPayment, overpayment, loanSum, lastPaymentDate := CalculateMortgage(tt.objectCost, tt.initialPayment, tt.months, tt.rate)

			if monthlyPayment != tt.expectedMonthlyPayment {
				t.Errorf("expected monthly payment %d, got %d", tt.expectedMonthlyPayment, monthlyPayment)
			}
			if overpayment != tt.expectedOverpayment {
				t.Errorf("expected overpayment %d, got %d", tt.expectedOverpayment, overpayment)
			}
			if loanSum != tt.expectedLoanSum {
				t.Errorf("expected loan sum %d, got %d", tt.expectedLoanSum, loanSum)
			}

			// Проверяем дату последнего платежа
			startDate := time.Now()
			expectedLastPaymentDate := startDate.AddDate(0, tt.months, 0)
			if lastPaymentDate.Year() != expectedLastPaymentDate.Year() || lastPaymentDate.Month() != expectedLastPaymentDate.Month() || lastPaymentDate.Day() != expectedLastPaymentDate.Day() {
				t.Errorf("expected last payment date %v, got %v", expectedLastPaymentDate.Format("2006-01-02"), lastPaymentDate.Format("2006-01-02"))
			}
		})
	}
}
