package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func EncryptPassword(password string) (string, error) {
	// 检查密码长度
	if len(password) == 0 {
		return "", fmt.Errorf("密码不能为空")
	}

	// 使用 bcrypt 加密，成本因子设为 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("密码加密失败: %v", err)
	}
	return string(hashedPassword), nil
}

// 验证密码
func ComparePassword(hashedPassword, password string) bool {
	// 检查输入
	if len(hashedPassword) == 0 || len(password) == 0 {
		fmt.Printf("密码验证失败: 输入为空 (hash长度=%d, 密码长度=%d)\n",
			len(hashedPassword), len(password))
		return false
	}

	// 尝试解析哈希
	if !isValidBcryptHash(hashedPassword) {
		fmt.Printf("密码验证失败: 无效的哈希格式: %s\n", hashedPassword)
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Printf("密码验证失败: %v\n", err)
		return false
	}
	return true
}

// 检查是否是有效的 bcrypt 哈希
func isValidBcryptHash(hash string) bool {
	// bcrypt 哈希应该以 $2a$ 开头
	if len(hash) < 4 || hash[:4] != "$2a$" {
		return false
	}
	return true
}
