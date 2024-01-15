package constant

import "time"

const (
	JwtKey      = "zust-acm"
	TokenExpire = 7 * 24 * time.Hour

	RedisTokenPrefix = "AuthToken"

	DateTimeLayout = "2006-01-02 15:04:05"

	SuperAdmin = "SUPER_ADMIN"
)
