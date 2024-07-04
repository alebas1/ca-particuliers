package cav1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/alebas1/ca-particuliers/internal/domain/entities"
	"github.com/samber/mo"
)

func (a *CAV1Repository) createSession(user entities.User) error {
	session := NewSession()
	regionalBankAlias, err := getRegionalBankAlias(user.RegionCode)
	if err != nil {
		return fmt.Errorf("failed to get regional bank URL fragment: %v", err)
	}
	session.RegionalBankAlias = regionalBankAlias
	session.Referer = fmt.Sprintf("https://www.credit-agricole.fr/%s/particulier", session.RegionalBankAlias)

	a.session = mo.Some(session)
	err = a.getKeypad()
	if err != nil {
		return fmt.Errorf("failed to get keypad: %v", err)
	}

	err = a.passSecurityCheck(user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("failed to pass security check: %v", err)
	}
	a.session.MustGet().SetAuthenticated()
	return nil
}

type KeypadResponse struct {
	Layout []string `json:"keyLayout"`
	Id     string   `json:"keypadId"`
}

func (a *CAV1Repository) getKeypad() error {
	getKeypadUrl := fmt.Sprintf("%s/acceder-a-mes-comptes.authenticationKeypad.json", a.session.MustGet().Referer)

	req, err := http.NewRequest(http.MethodPost, getKeypadUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Referer", fmt.Sprintf("%s/acceder-a-mes-comptes.html", a.session.MustGet().Referer))
	req.Header.Set("User-Agent", "CA Particuliers/1.0.0")

	res, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fetch keypad from api: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch key pad from api: Error code=%v", res.StatusCode)
	}

	var keypadResponse KeypadResponse
	err = json.NewDecoder(res.Body).Decode(&keypadResponse)
	if err != nil {
		return fmt.Errorf("failed to decode keypad response: %v", err)
	}

	a.session.MustGet().AppendCookies(res.Cookies())
	a.session.MustGet().Keypad = Keypad{
		Id:     keypadResponse.Id,
		Layout: keypadResponse.Layout,
	}
	return nil
}

func (a *CAV1Repository) passSecurityCheck(username string, passcode []string) error {
	computedPasscode, err := computePasscodeCombination(passcode, a.session.MustGet().Keypad.Layout)
	if err != nil {
		return fmt.Errorf("failed to compute passcode combination: %v", err)
	}

	var param = url.Values{}
	param.Set("j_password", strings.Join(computedPasscode, ","))
	param.Set("path", "/content/npc/start")
	param.Set("j_path_ressource", fmt.Sprintf("%%2F%s%%2Fparticulier%%2Foperations%%2Fsynthese.html", a.session.MustGet().RegionalBankAlias))
	param.Set("j_username", username)
	param.Set("keypadId", a.session.MustGet().Keypad.Id)
	param.Set("j_validate", "true")
	payload := strings.NewReader(param.Encode())

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/acceder-a-mes-comptes.html/j_security_check", a.session.MustGet().Referer), payload)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Referer", fmt.Sprintf("%s/acceder-a-mes-comptes.html", a.session.MustGet().Referer))
	req.Header.Set("User-Agent", "CA Particuliers/1.0.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	jar, err := cookiejar.New(nil)
	if err != nil {
		return fmt.Errorf("failed to create cookie jar: %v", err)
	}
	jar.SetCookies(req.URL, a.session.MustGet().Cookies)
	a.httpClient.Jar = jar

	res, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to authenticate user with api: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to authenticate user with api: Error code=%v", res)
	}
	a.session.MustGet().AppendCookies(res.Cookies())

	return nil
}
