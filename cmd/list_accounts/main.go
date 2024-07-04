package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alebas1/ca-particuliers/pkg/accounts"
)

func main() {
	regionCode := os.Getenv("CA_REGION_CODE")
	username := os.Getenv("CA_USERNAME")
	passcode := strings.Split(os.Getenv("CA_PASSCODE"), "")

	accountList, err := accounts.ListAccounts(username, passcode, regionCode)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", accountList)
}
