package dtos

type ListAccountsCommand struct {
	Username   string
	Passcode   []string
	RegionCode string
}

type AccountResult struct {
	Number                string
	Name                  string
	Balance               float64
	LoanAmount            float64
	LoanAmountOutstanding float64
	LoanInstalment        float64
	LoanPeriodicity       string
}
