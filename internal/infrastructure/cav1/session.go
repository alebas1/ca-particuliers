package cav1

import (
	"errors"
	"net/http"
	"time"
)

type Keypad struct {
	Id     string
	Layout []string
}

type Session struct {
	Authenticated     bool
	Referer           string
	Cookies           []*http.Cookie
	RegionalBankAlias string
	Keypad            Keypad
	CreationDate      time.Time
	UpdateDate        time.Time
}

func NewSession() *Session {
	now := time.Now()
	return &Session{
		Authenticated:     false,
		Referer:           "",
		RegionalBankAlias: "",
		Keypad:            Keypad{},
		Cookies:           []*http.Cookie{},
		CreationDate:      now,
		UpdateDate:        now,
	}
}

func (s *Session) RefreshUpdateDate() {
	s.UpdateDate = time.Now()
}

func (s *Session) SetAuthenticated() {
	s.RefreshUpdateDate()
	s.Authenticated = true
}

func (s *Session) AppendCookies(cookies []*http.Cookie) {
	s.RefreshUpdateDate()
	s.Cookies = append(s.Cookies, cookies...)
}

func (s Session) Validate() error {
	if !s.Authenticated {
		return errors.New("session is not authenticated")
	}
	if s.Referer == "" {
		return errors.New("session referer is empty")
	}
	if len(s.Cookies) == 0 {
		return errors.New("session has no cookies")
	}
	if s.UpdateDate.Before(s.CreationDate) || s.UpdateDate.Equal(s.CreationDate) {
		return errors.New("session update date is inferior or equal to creation date")
	}
	return nil
}
