package accounts

import (
	"net/http"

	"github.com/alebas1/ca-particuliers/internal/application/dtos"
	"github.com/alebas1/ca-particuliers/internal/application/services"
	"github.com/alebas1/ca-particuliers/internal/infrastructure/cav1"
	"github.com/samber/mo"
)

func ListAccounts(username string, password []string, regionCode string) ([]dtos.AccountResult, error) {
	service := services.NewAccountService(cav1.NewCAV1Repository(mo.None[*http.Client]()))
	list, err := service.ListAllAccounts(dtos.ListAccountsCommand{
		Username:   username,
		Passcode:   password,
		RegionCode: regionCode,
	})
	if err != nil {
		return nil, err
	}

	return list, nil
}
