package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	secretKey = "your-secret-key"
)

// GenerateToken 生成 JWT token
func GenerateToken(userID string) (string, error) {
	// 定义 JWT 的有效期限
	expirationTime := time.Now().Add(24 * time.Hour) // 设置为 24 小时有效期，可根据需求调整

	// 创建 token 的声明部分
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	}

	// 使用 HS256 算法进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对 token 进行签名，生成字符串格式的 token
	tokenString, err := token.SignedString([]byte(secretKey)) // 使用与验证时相同的密钥
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateToken 生成 JWT token
func GenerateWxUserToken(userID string) (string, error) {
	// 定义 JWT 的有效期限
	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 设置为 24 小时有效期，可根据需求调整

	// 创建 token 的声明部分
	claims := jwt.MapClaims{
		"wx_user_id": userID,
		"exp":        expirationTime.Unix(),
	}

	// 使用 HS256 算法进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对 token 进行签名，生成字符串格式的 token
	tokenString, err := token.SignedString([]byte(secretKey)) // 使用与验证时相同的密钥
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析 JWT token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil // 使用与生成 token 时相同的密钥
	})
	if err != nil {
		return nil, err
	}

	// 获取 token 中的声明部分
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

// 解析token返回user_id
func ParseTokenGetUserID(tokenString string) (string, string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", "", err
	}

	userID, ok := claims["user_id"].(string)
	if ok {
		return userID, "", err
	}

	wxUserId, ok := claims["wx_user_id"].(string)
	if ok {
		return "", wxUserId, err
	}

	return "", "", errors.New("user_id not found")
}

// 生成验证码
func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(999999)
	return fmt.Sprintf("%06d", code) // 格式化为6位数，不足的前面补零
}

// SignBody 签名
func SignBody(body, secretKey []byte) string {
	mac := hmac.New(sha256.New, secretKey)
	mac.Write(body)
	return hex.EncodeToString(mac.Sum(nil))
}

// GenerateOrderID generates a unique order ID based on the current date and time including nanoseconds.
func GenerateOrderID() string {
	// Set the seed for random number generation
	rand.Seed(time.Now().UnixNano())

	// Get the current date and time including nanoseconds
	now := time.Now()
	dateStr := now.Format("20060102150405") // Format as YYYYMMDDHHMMSS
	nanoStr := fmt.Sprintf("%09d", now.Nanosecond())
	fmt.Println("nanoStr:", nanoStr)

	// Generate a random 4-digit number
	randomNum := rand.Intn(10000)
	randomStr := fmt.Sprintf("%04d", randomNum)

	// Combine the date, time, nanoseconds, and random number to form the order ID
	orderID := dateStr + nanoStr + randomStr
	return orderID
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateInviteCode 生成一个6位包含字母和数字的邀请码
func GenerateInviteCode(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
