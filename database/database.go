package database

import (
	"errors"

	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	// 导入数据接口
	_ "github.com/go-sql-driver/mysql"
)

// Database 数据库
type Database interface {
	config() *config
	Open() error
	DB() *gorm.DB
	AutoMigrate(values ...interface{}) Database
	Close()
}

type database struct {
	conf *config
	db   *gorm.DB
}

// NewDatabase ..
func NewDatabase(opts ...Option) Database {
	return newDatabase(opts...)
}

func newDatabase(opts ...Option) Database {
	d := &database{
		conf: defaultConfig(),
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

// config ..
func (d *database) config() *config {
	return d.conf
}

// Open ..
func (d *database) Open() error {
	url, err := d.url()
	if err != nil {
		return err
	}

	db, err := gorm.Open(d.conf.dialect, url)
	if err != nil {
		return err
	}

	db.LogMode(d.conf.logDebug)

	if d.conf.maxIdleConns > 0 {
		db.DB().SetMaxIdleConns(d.conf.maxIdleConns)
	}
	if d.conf.maxOpenConns > 0 {
		db.DB().SetMaxOpenConns(d.conf.maxOpenConns)
	}
	if d.conf.maxConnLifeTime > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(d.conf.maxConnLifeTime) * time.Second)
	}

	// 表名默认不使用复用形式，比如表名使用 user 而不是 users
	// 使用 TableName 设置的除外
	db.SingularTable(true)

	d.db = db

	return nil
}

// DB ..
func (d *database) DB() *gorm.DB {
	return d.db
}

// AutoMigrate 迁移
func (d *database) AutoMigrate(values ...interface{}) Database {
	d.db.AutoMigrate(values...)
	return d
}

// Close ..
func (d *database) Close() {
	d.db.Close()
}

func (d *database) url() (string, error) {
	c := d.conf
	s := ""

	switch d.conf.dialect {
	case "mysql":
		s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			c.username, c.password, c.host, c.port, c.dbname)
	case "postgres":
		s = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			c.host, c.username, c.dbname, c.password)
	case "sqlite3":
		s = fmt.Sprintf("%s/%s", c.host, c.dbname)
	}

	if s == "" {
		return "", errors.New("invalid dialect")
	}

	return s, nil
}
