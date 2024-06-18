package main

import "github.com/alebas1/ca-particuliers/pkg/authenticator"

func main() {
	username := "XXXXXXXXXXX"
	passcode := []string{"0", "0", "0", "0", "0", "0"}
	session, err := authenticator.CreateSession(username, passcode, "62")
	if err != nil {
		panic(err)
	}
	_ = session
}
