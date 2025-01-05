package ntoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims 是自定义的JWT claims结构
type CustomClaims struct {
	jwt.RegisteredClaims                        // 包含标准的注册声明，如 exp 和 iat
	Uid                  string                 `json:"uid"`
	Payloads             map[string]interface{} `json:"payloads,omitempty"` // 可选的额外负载
}

// GenerateJwt 生成JWT
func GenerateJwt(iat int64, uid string, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	// 创建 Claims 对象
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Unix(iat, 0)),
			ExpiresAt: jwt.NewNumericDate(time.Unix(iat+seconds, 0)),
		},
		Uid:      uid,
		Payloads: payloads,
	}

	// 创建 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并返回 token 字符串
	return token.SignedString([]byte(secretKey))
}

// ParseJwt TODO 解析和验证JWT
func ParseJwt(tokenString string, secretKey string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	// 使用自定义 Claims 类型进行解析
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法是否匹配
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
