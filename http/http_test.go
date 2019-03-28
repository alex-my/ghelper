package http

import (
	"fmt"
	"testing"

	"github.com/alex-my/ghelper/random"
)

func TestGet(t *testing.T) {
	url := "https://www.keylala.cn"
	_, err := Get(url, nil, nil, 0)
	if err != nil {
		t.Errorf("Get failed: %s", err.Error())
	}
}

func TestGetTimeout(t *testing.T) {
	url := fmt.Sprintf("http://%s", random.RandomString(16))
	t.Logf("url: %s", url)
	_, err := Get(url, nil, nil, 3)
	if err == nil {
		t.Errorf("Get timeout failed, url: %s", url)
	}
}

func TestGetRetry(t *testing.T) {
	url := fmt.Sprintf("http://%s", random.RandomString(16))
	t.Logf("url: %s", url)
	_, err := GetRetry(url, nil, nil, 2, 3, 1000)
	if err == nil {
		t.Errorf("Get retry failed, url: %s", url)
	}
}

func TestPost(t *testing.T) {
	url := fmt.Sprintf("http://%s", random.RandomString(16))
	_, err := Post(url, nil, nil, nil, 0)
	if err == nil {
		t.Errorf("Get failed: %s", err.Error())
	}
}
