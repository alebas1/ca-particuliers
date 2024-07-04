package entities

import "fmt"

type Account struct {
	Number                string
	Name                  string
	Index                 int
	FamilyCode            int
	Balance               float64
	LoanAmount            float64
	LoanAmountOutstanding float64
	LoanInstalment        float64
	LoanPeriodicity       string
}

func (a *Account) Validate() error {
	if a.Number == "" {
		return fmt.Errorf("account number is required")
	}
	if a.Name == "" {
		return fmt.Errorf("account name is required")
	}
	if a.Index < 0 {
		return fmt.Errorf("account index is invalid")
	}
	return nil
}
