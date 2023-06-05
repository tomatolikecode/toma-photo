package format

import "fmt"

// 传入字节类型的数据, 返回对应B,KB,MB,GB,TB格式的字符串, 小数点保留一位小数点
func ByteCountIEC(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
