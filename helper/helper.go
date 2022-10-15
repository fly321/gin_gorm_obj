package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/jordan-wright/email"
	"net/smtp"
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

func SendEmail(toEmail, code string) error {
	e := email.NewEmail()
	e.From = "fly321 <944219401@qq.com>"
	e.To = []string{toEmail}
	e.Subject = "主题：验证码发送测试"
	//e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>我是测试验证码：" + code + "</h1>")
	//err := e.Send("smtp.qq.com:465", smtp.PlainAuth("", "944219401@qq.com", "密钥", "smtp.qq.com"))
	// eof 使用下面的tls
	return e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "944219401@qq.com", "密钥", "smtp.qq.com"),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         "smtp.qq.com",
		})
}
