package utils

import "golang.org/x/crypto/bcrypt"

const _bcryptCost = 11

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	if password == "" {
		return ""
	}
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), _bcryptCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	if password == "" || hash == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
