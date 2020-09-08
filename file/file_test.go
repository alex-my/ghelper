package file_test

import (
	"testing"

	"github.com/alex-my/ghelper/file"
)

func TestFile(t *testing.T) {
	filePath := "./file_test.go"
	dirPath := "."

	if !file.IsFile(filePath) {
		t.Error("IsFile error")
	}

	if file.IsFile(dirPath) {
		t.Error("IsFile error, dirPath")
	}

	if !file.IsDir(dirPath) {
		t.Error("IsDir error")
	}

	if file.IsExist("/hi/ROS9cRYp") {
		t.Error("IsExist error")
	}

	if !file.IsExist(dirPath) {
		t.Error("IsExist error, dirPath")
	}
}

func TestFileName(t *testing.T) {
	filePath := "./file_test.go"
	name := file.Name(filePath)
	if name != "file_test.go" {
		t.Errorf("Name error, name: %s", name)
	}
}

func TestFileNameRandom(t *testing.T) {
	name := file.NameRand("score.txt")
	if name == "" {
		t.Error("name is empty")
	} else {
		t.Logf("name: %s", name)
	}
}

func TestBaseName(t *testing.T) {
	filePath := "./file_test.go"
	name := file.BaseName(filePath)
	if name != "file_test" {
		t.Errorf("BaseName error, name: %s", name)
	}
}

func TestExtensionName(t *testing.T) {
	filePath := "./file_test.go"
	name := file.ExtensionName(filePath)
	if name != ".go" {
		t.Errorf("ExtensionName error, name: %s", name)
	}
}
