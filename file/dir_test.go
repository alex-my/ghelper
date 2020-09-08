package file_test

import (
	"testing"

	"github.com/alex-my/ghelper/file"
)

func TestListDirs(t *testing.T) {
	dirs, err := file.ListDirs("../")
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(len(dirs))
}

func TestDirContains(t *testing.T) {
	if !file.DirContains("..", ".git") {
		t.Fatal("contains .git")
	}

	if file.DirContains("..", ".unknwon") {
		t.Fatal("unknon exist")
	}
}
