package authenticator

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alebas1/ca-particuliers/pkg/regionalbankurlaliases"
)

type Authenticator struct {
	Username   string
	Passcode   []string
	RegionCode string
	Cookies    string
}

type KeypadResponse struct {
	KeyLayout []string `json:"keyLayout"`
	KeypadId  string   `json:"keypadId"`
}

func newAuthenticator(username string, passcode []string, regionCode string) *Authenticator {
	return &Authenticator{
		Username:   username,
		Passcode:   passcode,
		RegionCode: regionCode,
	}
}

func CreateSession(username string, passcode []string, regionCode string) (*Authenticator, error) {
	session := newAuthenticator(username, passcode, regionCode)
	err := session.authenticate()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (a *Authenticator) authenticate() error {
	regionalBankUrlFragment, err := regionalbankurlaliases.GetRegionalBankAlias(a.RegionCode)
	if err != nil {
		return fmt.Errorf("failed to get regional bank URL fragment: %v", err)
	}
	requestURL := fmt.Sprintf("https://www.credit-agricole.fr/%s/particulier", regionalBankUrlFragment)

	_, err = a.getKeypad(requestURL)
	if err != nil {
		return fmt.Errorf("failed to authenticate with keypad: %v", err)
	}

	// Load authentication keypad cookies
	return nil
}

func (a *Authenticator) getKeypad(requestURL string) (string, error) {
	// Create a new HTTP client with default settings
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false, // Enforce SSL/TLS verification
			},
		},
	}

	// Create a new POST request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/acceder-a-mes-comptes.authenticationKeypad.json", requestURL), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Referer", fmt.Sprintf("%s/acceder-a-mes-comptes.html", requestURL))
	req.Header.Set("User-Agent", "CA Particuliers/1.0.0")

	res, _ := client.Do(req)
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to authenticate with keypad: Error code=%v", res)
	}

	var keypadResponse KeypadResponse
	err = json.NewDecoder(res.Body).Decode(&keypadResponse)
	if err != nil {
		return "", err
	}

	// TODO: To put in auth payload as form data
	j_password := ComputePasswordCombination(a.Passcode, keypadResponse.KeyLayout)
	_ = j_password
	keypadId := keypadResponse.KeypadId
	_ = keypadId
	j_username := a.Username
	_ = j_username

	a.Cookies = res.Header.Get("Set-Cookie")

	return "keypadID", nil
}
