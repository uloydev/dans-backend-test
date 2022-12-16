package model

import "time"

type JWTMetadata struct {
	Expires        time.Time
	UserId         uint
	Username       string
	IsRefreshToken bool
}
