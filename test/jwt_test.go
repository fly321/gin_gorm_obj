package test

import (
	"github.com/dgrijalva/jwt-go/v4"
	"testing"
)

type UserClaims struct {
	Name     string `json:"name"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}

var myKey = []byte("gin-gorm-obj-key")

func TestGenerateToken(t *testing.T) {
	// 生成Token
	UserClaim := UserClaims{
		Name:           "user1",
		Identity:       "123456",
		StandardClaims: jwt.StandardClaims{},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	token, err := claims.SignedString(myKey)
	if err != nil {
		t.Error(err)
	}
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcjEiLCJpZGVudGl0eSI6IjEyMzQ1NiJ9.jtFUAyfzIUboDgUuSvXhONi667Xxr-kJaqCP_WsJ5Dk
	t.Log(token)

}

func TestAnalyseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlcjEiLCJpZGVudGl0eSI6IjEyMzQ1NiJ9.jtFUAyfzIUboDgUuSvXhONi667Xxr-kJaqCP_WsJ5Dk"
	UserClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, UserClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		t.Error(err)
	}
	if claims.Valid {
		t.Log(UserClaim)
	}
}
