package main

import (
	"fmt"
	"log"
	"ntoken"
	"time"
)

func main() {
	// 测试用的密钥和用户ID
	secretKey := "my_secret_key"
	uid := "user12345"

	// 创建一个额外负载
	payloads := map[string]interface{}{
		"user_role": "admin",
	}

	// 当前时间戳
	iat := time.Now().Unix()

	// 令牌的有效期（以秒为单位）
	expirationSeconds := int64(60 * 60) // 1小时有效期

	// 生成JWT
	tokenString, err := ntoken.GenerateJwt(iat, uid, secretKey, payloads, expirationSeconds)
	if err != nil {
		log.Fatalf("Failed to generate JWT: %v", err)
	}
	fmt.Printf("Generated JWT: %s\n", tokenString)

	// 解析JWT
	parsedClaims, err := ntoken.ParseJwt(tokenString, secretKey)
	if err != nil {
		log.Fatalf("Failed to parse JWT: %v", err)
	}
	fmt.Printf("Parsed JWT Claims: %+v\n", parsedClaims)

	// 检查解析后的 Claims 是否与原始值匹配
	if parsedClaims.Uid == uid && parsedClaims.Payloads["user_role"] == "admin" {
		fmt.Println("Token verification succeeded.")
	} else {
		log.Println("Token verification failed.")
	}

	// 测试使用错误的密钥解析 JWT
	errSecretKey := "wrong_secret_key"
	_, err = ntoken.ParseJwt(tokenString, errSecretKey)
	if err != nil {
		fmt.Printf("Expected error with wrong secret key: %v\n", err)
	} else {
		log.Fatal("Unexpected success with wrong secret key")
	}
}
