package gzip_test

import (
	"testing"

	"github.com/alex-my/ghelper/compress/gzip"
)

func TestGZip(t *testing.T) {
	s := "test gzip"
	c := gzip.NewGZip()

	compressed, err := c.Compress([]byte(s))
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	uncompressed, err := c.Uncompress(compressed)
	if err != nil {
		t.Fatal(err.Error())
	}

	s2 := string(uncompressed)
	if s != s2 {
		t.Fatalf("s: %s is not the same as s2: %s", s, s2)
	}
}
