package file

import (
	"io/ioutil"
	"path"
)

// ListDirs 列出指定目录的文件夹
func ListDirs(dirname string) ([]string, error) {
	dirs, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(dirs))

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		newdirname := path.Join(dirname, dir.Name())

		out = append(out, newdirname)
	}

	return out, nil
}

// DirContains 文件夹下是否包含指定名称的文件或者文件夹
func DirContains(dirname string, name string) bool {
	return IsExist(path.Join(dirname, name))
}
