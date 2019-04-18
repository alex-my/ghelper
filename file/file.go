package file

import (
	"bytes"
	"os"
	_path "path"
	"path/filepath"
	"strings"

	"github.com/alex-my/ghelper/random"
)

// IsDir 判断路径是否是文件夹
func IsDir(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return s.IsDir()
	}

	return false
}

// IsFile 判断路径是否是文件
func IsFile(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return !s.IsDir()
	}

	return false
}

// IsExist 判断路径上的 文件/文件夹 是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// Name 获取文件名
func Name(path string) string {
	_, name := filepath.Split(path)
	return name
}

// NameRand 获取一个随机文件名称
func NameRand(path string, sep ...string) string {
	s := "_"
	if len(sep) > 0 && sep[0] != "" {
		s = sep[0]
	}

	buf := &bytes.Buffer{}
	name := BaseName(path)
	rand := random.String(12)
	ext := ExtensionName(path)

	buf.WriteString(name)
	buf.WriteString(s)
	buf.WriteString(rand)
	buf.WriteString(ext)

	return buf.String()
}

// BaseName 获取文件名，不带后缀
func BaseName(path string) string {
	return strings.TrimSuffix(filepath.Base(path), _path.Ext(path))
}

// ExtensionName 获取文件拓展名
// some_file.go -> .go
// some_file.py -> .py
func ExtensionName(path string) string {
	return _path.Ext(path)
}
