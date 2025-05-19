package models

type Calculation struct {
	ID         uint64    `json:"id"`
	Params     Params    `json:"params"`
	Program    Program   `json:"program"`
	Aggregates Aggregate `json:"aggregates"`
}

type Params struct {
	ObjectCost     int      `json:"object_cost"`
	InitialPayment int      `json:"initial_payment"`
	Months         int      `json:"months"`
	Program        *Program `json:"program,omitempty"`
}

type Program struct {
	Salary   *bool `json:"salary,omitempty"`
	Military *bool `json:"military,omitempty"`
	Base     *bool `json:"base,omitempty"`
}

type Aggregate struct {
	Rate            int    `json:"rate"`
	LoanSum         int    `json:"loan_sum"`
	MonthlyPayment  int    `json:"monthly_payment"`
	Overpayment     int    `json:"overpayment"`
	LastPaymentDate string `json:"last_payment_date"`
}
