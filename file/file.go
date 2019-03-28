package file

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/alex-my/ghelper/random"
)

// IsDir 判断路径是否是文件夹
func IsDir(_path string) bool {
	if s, err := os.Stat(_path); err == nil {
		return s.IsDir()
	}
	return false
}

// IsFile 判断路径是否是文件
func IsFile(_path string) bool {
	return !IsDir(_path)
}

// IsExist 判断路径上的 文件/文件夹 是否存在
func IsExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// Name 获取文件名
func Name(_path string) string {
	_, name := filepath.Split(_path)
	return name
}

// NameRand 获取一个随机文件名称
func NameRand(_path string) string {
	name := BaseName(_path)
	return name + "-" + random.RandomString(12)
}

// BaseName 获取文件名，不带后缀
func BaseName(_path string) string {
	return strings.TrimSuffix(filepath.Base(_path), path.Ext(_path))
}

// ExtensionName 获取文件拓展名
// some_file.go -> .go
// some_file.py -> .py
func ExtensionName(_path string) string {
	return path.Ext(_path)
}
