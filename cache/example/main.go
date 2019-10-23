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
		log.Errorf("testString failed: %s", err.Error())
		return
	}

	if err := testHash(c); err != nil {
		log.Errorf("testHash failed: %s", err.Error())
		return
	}

	if err := testList(c); err != nil {
		log.Errorf("testList failed: %s", err.Error())
		return
	}

	if err := testSet(c); err != nil {
		log.Errorf("testList failed: %s", err.Error())
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
		return errors.New("testString error 1")
	}

	v, err := c.Get(key)
	if err != nil {
		return err
	}
	if v != value {
		return errors.New("testString error 2")
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
	vs, err := c.MGet("key-1", "key-2", "key-3")
	if err != nil {
		return err
	}
	if len(vs) != 3 {
		return errors.New("testString error 5")
	}

	c.Set(key, value)
	n, err := c.Strlen(key)
	if err != nil {
		return err
	}
	if n != len(key) {
		return errors.New("testString error 6")
	}
	n, err = c.Append(key, "haha")
	if err != nil {
		return err
	}
	n2, _ := c.Strlen(key)
	if n != n2 {
		return errors.New("testString error 7")
	}

	c.Set(key, value)
	newValue := "gogogo"
	v, err = c.GetSet(key, newValue)
	if err != nil {
		return err
	}
	if v != value {
		return errors.New("testString error 8")
	}
	v, err = c.Get(key)
	if err != nil {
		return err
	}
	if v != newValue {
		return errors.New("testString error 9")
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
		return errors.New("testString error 10")
	}
	n64, err = c.Incrby(nKey, 6)
	if err != nil {
		return err
	}
	if n64 != 10 {
		return errors.New("testString error 11")
	}

	c.Set(key, value)
	exist, err := c.Exists(key)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("testString error 12")
	}

	delSize, err := c.Del(key, nKey)
	if err != nil {
		return err
	}
	if delSize != 2 {
		return errors.New("testString error 13")
	}

	exist, err = c.Exists(key)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("testString error 14")
	}

	c.DO("FLUSHDB")

	return nil
}

func testHash(c cache.Cache) error {
	var (
		key                    = "hello"
		field1, field2, field3 = "field1", "field2", "field3"
		value1, value2, value3 = "value1", "value2", "value3"
	)

	c.HMSet(key, field1, value1, field2, value2)
	c.HSet(key, field3, value3)
	n, err := c.HLen(key)
	if err != nil {
		return err
	}
	if n != 3 {
		return errors.New("testHash error 1")
	}

	n, err = c.HDel(key, field1, field2, "field4")
	if err != nil {
		return err
	}
	if n != 2 {
		return errors.New("testHash error 2")
	}

	vs, err := c.HGetAll(key)
	if err != nil {
		return err
	}
	// field-value
	if len(vs) != 2 {
		return errors.New("testHash error 3")
	}

	v, err := c.HGet(key, field3)
	if err != nil {
		return err
	}
	if v != value3 {
		return errors.New("testHash error 4")
	}

	c.DO("FLUSHDB")

	return nil
}

func testList(c cache.Cache) error {
	var (
		key            = "key"
		value1, value2 = "value1", "value2"
	)

	n, _ := c.LPush(key, value1)
	if n != 1 {
		return errors.New("testList error 1")
	}
	n, _ = c.RPush(key, value2)
	if n != 2 {
		return errors.New("testList error 2")
	}

	n, _ = c.LLen(key)
	if n != 2 {
		return errors.New("testList error 3")
	}

	v, _ := c.LIndex(key, 0)
	if v != value1 {
		return errors.New("testList error 4")
	}

	vs, _ := c.LRange(key, 0, 100)
	if len(vs) != 2 || vs[0] != value1 || vs[1] != value2 {
		return errors.New("testList error 5")
	}

	c.DO("FLUSHDB")
	return nil
}

func testSet(c cache.Cache) error {
	var (
		key        = "key"
		m1, m2, m3 = "m1", "m2", "m3"
	)

	n, _ := c.SAdd(key, m1, m2, m3)
	n2, _ := c.SCard(key)
	if n != n2 {
		return errors.New("testSet error 1")
	}

	r, _ := c.SIsMember(key, m3)
	if !r {
		return errors.New("testSet error 2")
	}
	r, _ = c.SIsMember(key, "m4")
	if r {
		return errors.New("testSet error 3")
	}

	c.DO("FLUSHDB")
	return nil
}
