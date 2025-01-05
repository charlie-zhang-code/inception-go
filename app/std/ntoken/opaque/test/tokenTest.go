package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	accessTokenStore  = make(map[string]string)
	refreshTokenStore = make(map[string]string)
	accessTokenMutex  sync.Mutex
	refreshTokenMutex sync.Mutex
	tokenTTL          time.Duration = 15 * time.Minute   // 访问令牌的有效期
	refreshTokenTTL   time.Duration = 7 * 24 * time.Hour // 刷新令牌的有效期
	cleanupInterval   time.Duration = 30 * time.Minute   // 清理过期令牌的时间间隔
)

// 生成随机字符串作为令牌
func generateRandomToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// 生成访问令牌
func generateAccessToken(username string) (string, error) {
	token, err := generateRandomToken()
	if err != nil {
		return "", err
	}
	expiration := time.Now().Add(tokenTTL).Unix()

	accessTokenMutex.Lock()
	defer accessTokenMutex.Unlock()
	accessTokenStore[token] = username + ":" + strconv.FormatInt(expiration, 10)
	return token, nil
}

// 生成刷新令牌
func generateRefreshToken(username string) (string, error) {
	token, err := generateRandomToken()
	if err != nil {
		return "", err
	}
	expiration := time.Now().Add(refreshTokenTTL).Unix()

	refreshTokenMutex.Lock()
	defer refreshTokenMutex.Unlock()
	refreshTokenStore[token] = username + ":" + strconv.FormatInt(expiration, 10)
	return token, nil
}

// 校验访问令牌
func validateAccessToken(token string) (bool, string, error) {
	accessTokenMutex.Lock()
	defer accessTokenMutex.Unlock()

	data, exists := accessTokenStore[token]
	if !exists {
		return false, "", fmt.Errorf("access token not found")
	}

	parts := strings.SplitN(data, ":", 2)
	if len(parts) != 2 {
		return false, "", fmt.Errorf("invalid access token format")
	}

	username := parts[0]
	expirationStr := parts[1]
	expiration, err := strconv.ParseInt(expirationStr, 10, 64)
	if err != nil {
		return false, "", fmt.Errorf("invalid expiration format in access token: %v", err)
	}

	if time.Now().Unix() > expiration {
		delete(accessTokenStore, token)
		return false, "", fmt.Errorf("access token expired")
	}

	return true, username, nil
}

// 校验刷新令牌
func validateRefreshToken(token string) (bool, string, error) {
	refreshTokenMutex.Lock()
	defer refreshTokenMutex.Unlock()

	data, exists := refreshTokenStore[token]
	if !exists {
		return false, "", fmt.Errorf("refresh token not found")
	}

	parts := strings.SplitN(data, ":", 2)
	if len(parts) != 2 {
		return false, "", fmt.Errorf("invalid refresh token format")
	}

	username := parts[0]
	expirationStr := parts[1]
	expiration, err := strconv.ParseInt(expirationStr, 10, 64)
	if err != nil {
		return false, "", fmt.Errorf("invalid expiration format in refresh token: %v", err)
	}

	if time.Now().Unix() > expiration {
		delete(refreshTokenStore, token)
		return false, "", fmt.Errorf("refresh token expired")
	}

	return true, username, nil
}

// 使用刷新令牌获取新的访问令牌
func refreshToken(oldRefreshToken string) (string, error) {
	valid, username, err := validateRefreshToken(oldRefreshToken)
	if err != nil {
		fmt.Println("Failed to validate refresh token:", err)
		return "", err
	}
	if !valid {
		fmt.Println("Failed to validate refresh token:", err)
		return "", fmt.Errorf("refresh token validation failed")
	}

	newAccessToken, err := generateAccessToken(username)
	if err != nil {
		fmt.Println("Failed to validate refresh token:", err)
		return "", err
	}

	fmt.Printf("Validated Refresh Token for user: %s\n", username)

	return newAccessToken, nil
}

// 后台任务：清理过期的访问令牌
func cleanupExpiredAccessTokens() {
	for {
		time.Sleep(cleanupInterval)
		currentTime := time.Now().Unix()

		accessTokenMutex.Lock()
		for token, data := range accessTokenStore {
			parts := strings.SplitN(data, ":", 2)
			if len(parts) == 2 {
				expiration, err := strconv.ParseInt(parts[1], 10, 64)
				if err == nil && currentTime > expiration {
					delete(accessTokenStore, token)
				}
			}
		}
		accessTokenMutex.Unlock()
	}
}

// 后台任务：清理过期的刷新令牌
func cleanupExpiredRefreshTokens() {
	for {
		time.Sleep(cleanupInterval)
		currentTime := time.Now().Unix()

		refreshTokenMutex.Lock()
		for token, data := range refreshTokenStore {
			parts := strings.SplitN(data, ":", 2)
			if len(parts) == 2 {
				expiration, err := strconv.ParseInt(parts[1], 10, 64)
				if err == nil && currentTime > expiration {
					delete(refreshTokenStore, token)
				}
			}
		}
		refreshTokenMutex.Unlock()
	}
}

func main() {
	go cleanupExpiredAccessTokens()
	go cleanupExpiredRefreshTokens()

	username := "test_user_rand_id"

	accessToken, _ := generateAccessToken(username)
	fmt.Println("Access Token:", accessToken)

	refreshToken1, _ := generateRefreshToken(username)
	fmt.Println("Refresh Token:", refreshToken1)

	valid, user, err := validateAccessToken(accessToken)
	if err != nil {
		fmt.Println("Failed to validate access token:", err)
	} else if valid {
		fmt.Printf("Validated Access Token for user: %s\n", user)
	}

	valid, user, err = validateAccessToken("accessToken")
	if err != nil {
		fmt.Println("Failed to validate access token1:", err)
	} else if valid {
		fmt.Printf("Validated Access Token for user1: %s\n", user)
	}

	newAccessToken, err := refreshToken(refreshToken1)
	if err != nil {
		fmt.Println("Failed to refresh token:", err)
	} else {
		fmt.Println("New Access Token:", newAccessToken)
	}

	valid, user, err = validateAccessToken(newAccessToken)
	if err != nil {
		fmt.Println("Failed to validate access token1:", err)
	} else if valid {
		fmt.Printf("Validated Access Token for user1: %s\n", user)
	}

	newAccessToken, err = refreshToken("refreshToken1")
	if err != nil {
		fmt.Println("Failed to refresh token:", err)
	} else {
		fmt.Println("New Access Token:", newAccessToken)
	}

}
