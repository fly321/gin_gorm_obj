package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
)

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

type UserClaims struct {
	Name     string `json:"name"`
	Identity string `json:"identity"`
	jwt.StandardClaims
}

var myKey = []byte("gin-gorm-obj-key")

// GenerateToken 生成token
func GenerateToken(name, identity string) (string, error) {
	// 生成Token
	UserClaim := UserClaims{
		Name:           name,
		Identity:       identity,
		StandardClaims: jwt.StandardClaims{},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	token, err := claims.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return token, nil

}

// AnalyseToken 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	UserClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, UserClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("token is invalid : %s", err)
	}
	return UserClaim, nil
}
