package services

import (
	"fmt"

	"github.com/alebas1/ca-particuliers/internal/application/dtos"
	"github.com/alebas1/ca-particuliers/internal/domain/entities"
	"github.com/alebas1/ca-particuliers/internal/infrastructure/interfaces"
)

type AccountService struct {
	bankRepository interfaces.BankRepository
}

func NewAccountService(bankRepository interfaces.BankRepository) *AccountService {
	return &AccountService{
		bankRepository: bankRepository,
	}
}

func (s *AccountService) ListAllAccounts(listAccountsCmd dtos.ListAccountsCommand) ([]dtos.AccountResult, error) {
	user := entities.User{
		Username:   listAccountsCmd.Username,
		Password:   listAccountsCmd.Passcode,
		RegionCode: listAccountsCmd.RegionCode,
	}
	err := user.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid user input: %v", err)
	}

	accounts, err := s.bankRepository.GetAllAccounts(user)
	if err != nil {
		return nil, err
	}

	response := []dtos.AccountResult{}
	for _, account := range accounts {
		err := account.Validate()
		if err != nil {
			return nil, fmt.Errorf("invalid account: %v", err)
		}
		response = append(response, dtos.AccountResult{
			Number:                account.Number,
			Name:                  account.Name,
			Balance:               account.Balance,
			LoanAmount:            account.LoanAmount,
			LoanAmountOutstanding: account.LoanAmountOutstanding,
			LoanInstalment:        account.LoanInstalment,
			LoanPeriodicity:       account.LoanPeriodicity,
		})
	}

	return response, nil
}
