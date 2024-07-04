package interfaces

import "github.com/alebas1/ca-particuliers/internal/domain/entities"

type BankRepository interface {
	GetAllAccounts(user entities.User) ([]entities.Account, error)
	ListOperations(account entities.Account) ([]entities.Operation, error)
}
