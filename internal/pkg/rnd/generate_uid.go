package rnd

import (
	"strconv"
	"time"
)

// GenerateUID返回一个唯一的id, 前缀为字符串
func GenerateUID(prefix byte) string {
	result := make([]byte, 0, 16)
	result = append(result, prefix)
	result = append(result, strconv.FormatInt(time.Now().UTC().Unix(), 36)[0:6]...)
	result = append(result, GenerateToken(9)...)

	return string(result)
}
