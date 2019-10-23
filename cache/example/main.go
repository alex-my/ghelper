package main

import (
	"errors"

	"github.com/alex-my/ghelper/cache"
	"github.com/alex-my/ghelper/logger"
)

var (
	// RedisHost ...
	RedisHost = "127.0.0.1"
	// RedisPort ...
	RedisPort = 11234
	// RedisPassword ...
	RedisPassword = "uio876..."
)

var log = logger.NewLogger()

func main() {

	c := cache.NewCache(
		cache.WithHost(RedisHost),
		cache.WithPort(RedisPort),
		cache.WithPassword(RedisPassword),
	)

	err := c.Open()
	if err != nil {
		log.Error("cache open failed: %s", err.Error())
		return
	}

	if err := testString(c); err != nil {
		log.Errorf("test1 failed: %s", err.Error())
		return
	}

	log.Info("test cache success")
}

func testString(c cache.Cache) error {
	const (
		key   = "hello"
		value = "world"
	)

	c.Set(key, value)
	ttl, err := c.TTL(key)
	if err != nil {
		return err
	}
	if ttl != -1 {
		return errors.New("error 1")
	}

	v, err := c.Get(key)
	if err != nil {
		return err
	}
	if v != value {
		return errors.New("error 2")
	}

	c.SetEx(key, value, "10")
	ttl, err = c.TTL(key)
	if err != nil {
		return err
	}
	if ttl <= 0 || ttl > 10 {
		return errors.New("error 3")
	}

	c.PSetEx(key, value, "10000")
	ttl, err = c.PTTL(key)
	if err != nil {
		return err
	}
	if ttl <= 0 || ttl > 10000 {
		return errors.New("error 4")
	}

	c.MSet("key-1", "value-1", "key-2", "value-2")
	vs := c.MGet("key-1", "key-2", "key-3")
	if len(vs) != 3 {
		return errors.New("error 5")
	}

	c.Set(key, value)
	n, err := c.Strlen(key)
	if err != nil {
		return err
	}
	if n != len(key) {
		return errors.New("error 6")
	}
	n, err = c.Append(key, "haha")
	if err != nil {
		return err
	}
	n2, _ := c.Strlen(key)
	if n != n2 {
		return errors.New("error 7")
	}

	c.Set(key, value)
	newValue := "gogogo"
	v, err = c.GetSet(key, newValue)
	if err != nil {
		return err
	}
	if v != value {
		return errors.New("error 8")
	}
	v, err = c.Get(key)
	if err != nil {
		return err
	}
	if v != newValue {
		return errors.New("error 9")
	}

	var (
		nKey   = "num"
		nValue = "3"
	)

	c.Set(nKey, nValue)
	n64, err := c.Incr(nKey)
	if err != nil {
		return err
	}
	if n64 != 4 {
		return errors.New("error 10")
	}
	n64, err = c.Incrby(nKey, 6)
	if err != nil {
		return err
	}
	if n64 != 10 {
		return errors.New("error 11")
	}

	c.Set(key, value)
	exist, err := c.Exists(key)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("error 12")
	}

	delSize, err := c.Del(key, nKey)
	if err != nil {
		return err
	}
	if delSize != 2 {
		return errors.New("error 13")
	}

	exist, err = c.Exists(key)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("error 14")
	}

	c.DO("FLUSHDB")

	return nil
}
