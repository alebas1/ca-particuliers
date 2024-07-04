package entities

import "errors"

var (
	ErrInvalidLengthUsername   = errors.New("username must be 11 characters long")
	ErrInvalidLengthPassword   = errors.New("password must be 6 characters long")
	ErrInvalidLengthRegionCode = errors.New("region code must be 2 characters long")
)

type User struct {
	Username   string
	Password   []string
	RegionCode string
}

func NewUser(username string, password []string, regionCode string) *User {
	return &User{
		Username:   username,
		Password:   password,
		RegionCode: regionCode,
	}
}

func (u *User) Validate() error {
	if len(u.Username) != 11 {
		return ErrInvalidLengthUsername
	}

	if len(u.Password) != 6 {
		return ErrInvalidLengthUsername
	}

	if len(u.RegionCode) != 2 {
		return ErrInvalidLengthRegionCode
	}

	return nil
}
