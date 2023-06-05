package types

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID       uint
	UserUID  string
	Username string
	NickName string
	Phone    string
}
