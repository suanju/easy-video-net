package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Hotkey 密钥
var Hotkey = []byte("G0-store")

// SaltStr  密码盐的随机字符串
var SaltStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//Claims  TOKEN 的结构体
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// NextToken 登录以后签发jwt
func NextToken(userId uint) string {
	fmt.Printf("传入JWT的id:%v/n", userId)
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "root",       // 签名颁发者
			Subject:   "user token", //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(Hotkey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}

// ParseToken 解析 Token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Hotkey, nil
	})
	if err != nil {
		fmt.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
