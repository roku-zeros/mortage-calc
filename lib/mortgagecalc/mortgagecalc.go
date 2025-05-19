package mortagecalc

import (
	"math"
	"time"
)

func CalculateMortgage(objectCost, initialPayment, months int, rate int) (
    monthlyPayment, overpayment, loanSum int, lastPaymentDate time.Time) {
    loanSum = objectCost - initialPayment
    monthlyRate := float64(rate) / 100 / 12
    monthlyPayment = int(math.Ceil(float64(loanSum) * (monthlyRate * math.Pow(1+monthlyRate, float64(months)) / (math.Pow(1+monthlyRate, float64(months)) - 1))))
    overpayment = (monthlyPayment * months) - loanSum

    startDate := time.Now()
    lastPaymentDate = startDate.AddDate(0, months, 0)

    return monthlyPayment, overpayment, loanSum, lastPaymentDate
}

