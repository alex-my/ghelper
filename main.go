package main

import (
	"fmt"

	"github.com/alex-my/ghelper/sign"
)

func main() {
	testValue()
	testValues()
}

func testValue() {
	params := map[string]string{
		"timestamp": "1571747084",
		"username":  "root",
		"password":  "123456",
	}

	calcSign, _ := sign.Sign(params)
	fmt.Printf("calcSign: %s\n", calcSign)
}

func testValues() {

}
