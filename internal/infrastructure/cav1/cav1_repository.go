package cav1

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"github.com/alebas1/ca-particuliers/internal/domain/entities"
	"github.com/samber/mo"
)

type AccountFamily struct {
	Code int
	Name string
}

var ACCOUNT_FAMILY = []AccountFamily{
	{Code: 1, Name: "COMPTES"},
	{Code: 3, Name: "EPARGNE_DISPONIBLE"},
	{Code: 4, Name: "CREDITS"},
	{Code: 7, Name: "EPARGNE_AUTRE"},
}

type CAV1Repository struct {
	session    mo.Option[*Session]
	httpClient *http.Client
}

func NewCAV1Repository(httpClient mo.Option[*http.Client]) *CAV1Repository {
	return &CAV1Repository{
		session: mo.None[*Session](),
		httpClient: httpClient.OrElse(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: false,
				},
			},
		}),
	}
}

type AccountResponse struct {
	Number                string  `json:"numeroCompte"`
	Name                  string  `json:"libelleProduit"`
	Index                 int     `json:"index"`
	FamilyCode            int     `json:"codeFamille"`
	Balance               float64 `json:"solde"`
	SavingsBalance        float64 `json:"montantEpargne"`
	LoanAmount            float64 `json:"montantInitial"`
	LoanAmountOutstanding float64 `json:"montantRestantDu"`
	LoanInstalment        float64 `json:"montantEcheance"`
	LoanPeriodicity       string  `json:"libellePeriodicite"`
}

func (a *CAV1Repository) GetAllAccounts(user entities.User) ([]entities.Account, error) {
	err := a.createSession(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

	accounts := []entities.Account{}
	for _, family := range ACCOUNT_FAMILY {
		getAccountsURL := fmt.Sprintf("%s/operations/synthese/jcr:content.produits-valorisation.json/%d", a.session.MustGet().Referer, family.Code)
		req, err := http.NewRequest(http.MethodGet, getAccountsURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request to get accounts: %v", err)
		}
		req.Header.Set("User-Agent", "CA Particuliers/1.0.0")
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create cookie jar: %v", err)
		}
		jar.SetCookies(req.URL, a.session.MustGet().Cookies)
		a.httpClient.Jar = jar
		res, err := a.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to get accounts: %v", err)
		}
		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to get accounts: Error code=%v", res)
		}

		var accountsResponse []AccountResponse
		err = json.NewDecoder(res.Body).Decode(&accountsResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to decode accounts response: %v", err)
		}

		for _, accountResponse := range accountsResponse {
			account := entities.Account{
				Number:                accountResponse.Number,
				Name:                  accountResponse.Name,
				Index:                 accountResponse.Index,
				FamilyCode:            accountResponse.FamilyCode,
				LoanAmount:            accountResponse.LoanAmount,
				LoanAmountOutstanding: accountResponse.LoanAmountOutstanding,
				LoanInstalment:        accountResponse.LoanInstalment,
				LoanPeriodicity:       accountResponse.LoanPeriodicity,
			}

			if accountResponse.SavingsBalance != 0 {
				account.Balance = accountResponse.SavingsBalance
			} else {
				account.Balance = accountResponse.Balance
			}

			accounts = append(accounts, account)
		}
	}
	return accounts, nil
}

func (a *CAV1Repository) ListOperations(account entities.Account) ([]entities.Operation, error) {
	return nil, nil
}
