package file

import (
	"testing"
)

func TestFile(t *testing.T) {
	filePath := "./file_test.go"
	dirPath := "."

	if !IsFile(filePath) {
		t.Error("IsFile error")
	}

	if IsFile(dirPath) {
		t.Error("IsFile error, dirPath")
	}

	if !IsDir(dirPath) {
		t.Error("IsDir error")
	}

	if IsExist("/hi/ROS9cRYp") {
		t.Error("IsExist error")
	}

	if !IsExist(dirPath) {
		t.Error("IsExist error, dirPath")
	}
}

func TestFileName(t *testing.T) {
	filePath := "./file_test.go"
	name := Name(filePath)
	if name != "file_test.go" {
		t.Errorf("Name error, name: %s", name)
	}
}

func TestFileNameRandom(t *testing.T) {
	name := NameRand("score.txt")
	if name == "" {
		t.Error("name is empty")
	} else {
		t.Logf("name: %s", name)
	}
}

func TestBaseName(t *testing.T) {
	filePath := "./file_test.go"
	name := BaseName(filePath)
	if name != "file_test" {
		t.Errorf("BaseName error, name: %s", name)
	}
}

func TestExtensionName(t *testing.T) {
	filePath := "./file_test.go"
	name := ExtensionName(filePath)
	if name != ".go" {
		t.Errorf("ExtensionName error, name: %s", name)
	}
}
