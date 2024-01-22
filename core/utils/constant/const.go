package constant

import "time"

// common
const (
	JwtKey      = "zust-acm"
	TokenExpire = 7 * 24 * time.Hour

	RedisTokenPrefix = "AuthToken"

	DateTimeLayout = "2006-01-02 15:04:05"
)

// team user type
const (
	TeamMember  = "member"
	TeamTeacher = "teacher"
	TeamLeader  = "leader"
)

// role type
const (
	SuperAdmin = "SUPER_ADMIN"
)
