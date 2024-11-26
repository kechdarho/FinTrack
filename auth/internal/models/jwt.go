package models

import "time"

type AccessToken struct {
	Token string        `json:"access_token"`
	TTL   time.Duration `json:"access_token_ttl"`
}

type RefreshToken struct {
	Token string        `json:"refresh_token"`
	TTL   time.Duration `json:"refresh_token_ttl"`
}
