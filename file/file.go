package file

import (
	"bufio"
	"bytes"
	"io"
	"os"
	_path "path"
	"path/filepath"
	"strings"

	"github.com/alex-my/ghelper/random"
)

// IsDir 判断路径是否是文件夹
// eg:
//
// /tmp 为实际上存在的文件夹
// IsDir("/tmp") 					-> true
//
// /tmp2 为实际上不存在的文件夹
// IsDir("/tmp2") 					-> false
//
// /tmp/test.txt 为实际存在的文件
// IsDir("/tmp/test.txt") 			-> false
func IsDir(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return s.IsDir()
	}

	return false
}

// IsFile 判断路径是否是文件
// eg:
//
// /tmp 为实际存在的文件夹
// IsFile("/tmp") 					-> false
//
// /tmp/test.txt 为实际存在的文件
// IsFile("/tmp/test.txt") 			-> true
//
// /tmp/test2.txt 为实际不存在的文件
// IsFile("/tmp/test2.txt") 		-> false
//
// /tmp/test_link.txt 为有效的链接文件 ln -s src dst
// IsFile("/tmp/test_link.txt") 	-> true
func IsFile(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return !s.IsDir()
	}

	return false
}

// IsExist 判断路径上的 文件或文件夹 是否存在
// eg:
//
// /tmp 为实际存在的文件夹
// IsExist("/tmp") 					-> true
//
// /tmp/test.txt 为实际存在的文件
// IsExist("/tmp/test.txt") 		-> true
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// Name 获取文件名
// eg: Name("/tmp/test.txt") 		-> test.txt
func Name(path string) string {
	_, name := filepath.Split(path)
	return name
}

// NameRand 获取一个随机文件名称，默认以 _ 连接
// eg: NameRand("test.txt") 		-> test_869mUEfWXOaB.txt
// eg: NameRand("test.txt", "@") 	-> test@jQ0EQkDQ285x.txt
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
// eg: BaseName("/tmp/test.txt") 		-> test
func BaseName(path string) string {
	return strings.TrimSuffix(filepath.Base(path), _path.Ext(path))
}

// ExtensionName 获取文件拓展名
// eg: ExtensionName("/tmp/test.txt") 	-> .txt
func ExtensionName(path string) string {
	return _path.Ext(path)
}

// ReadLine 逐行读取文件、读取大文件
// path: 文件路径
// handle: 每读取一行的处理函数，返回的error为非nil时，不再继续向后读取
func ReadLine(path string, handle func([]byte) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, isPrefix, err := reader.ReadLine()

		// isPrefix 一行太长了，需要多次调用，默认一行为 4096 字节
		// 也可以修改 bufio.NewReader 为 bufio.NewReaderSize 自行修改默认一行的长度
		for isPrefix && err == nil {
			var bs []byte
			bs, isPrefix, err = reader.ReadLine()
			line = append(line, bs...)
		}

		if err := handle(line); err != nil {
			return err
		}

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
