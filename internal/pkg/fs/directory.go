package fs

import "os"

// PathExists 文件目录是否存在
func PathExists(path string) bool {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true
		}
		return false
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}
