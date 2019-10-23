package main

import (
	"fmt"

	"github.com/alex-my/ghelper/cache"
)

func main() {
	c := cache.NewCache(
		cache.WithPort(11234),
		cache.WithPassword("uio876..."),
	)

	err := c.Open()
	if err != nil {
		fmt.Printf("open failed: %s\n", err.Error())
		return
	}

	r1, err := c.SetString("a", "a0", "3000")
	if err != nil {
		fmt.Printf("setstring failed: %s\n", err.Error())
		return
	}
	fmt.Printf("setstring result, r1: %v\n", r1)

	r2, err := c.String("a")
	if err != nil {
		fmt.Printf("getstring failed: %s\n", err.Error())
		return
	}
	fmt.Printf("getstring result, r2: %v\n", r2)
}
