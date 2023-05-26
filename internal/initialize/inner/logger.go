package inner

import "gorm.io/gorm/logger"

type writer struct {
	logger.Writer
}

// writer 的构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}
