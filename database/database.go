package database

import (
	"database/sql"
	"errors"

	"fmt"

	"github.com/jinzhu/gorm"

	// 导入数据接口
	_ "github.com/go-sql-driver/mysql"
)

// Database 数据库
type Database interface {
	config() *config
	Open() error
	DB() *sql.DB
	Close()
}

type database struct {
	c  *config
	db *gorm.DB
}

// NewDatabase ..
func NewDatabase(opts ...Option) Database {
	return newDatabase(opts...)
}

func newDatabase(opts ...Option) Database {
	d := &database{
		c: defaultConfig(),
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

// config ..
func (d *database) config() *config {
	return d.c
}

// Open ..
func (d *database) Open() error {
	url, err := d.url()
	if err != nil {
		return err
	}

	db, err := gorm.Open(d.c.dialect, url)
	if err != nil {
		return err
	}

	if d.c.logDebug {
		db.LogMode(true)
	}

	if d.c.maxIdleConns > 0 {
		db.DB().SetMaxIdleConns(d.c.maxIdleConns)
	}
	if d.c.maxOpenConns > 0 {
		db.DB().SetMaxOpenConns(d.c.maxOpenConns)
	}

	// 表名默认不使用复用形式，比如表名使用 user 而不是 users
	// 使用 TableName 设置的除外
	db.SingularTable(true)

	return nil
}

// DB 获取当前与数据库连接中的 DB，如果未连接，返回 nil
func (d *database) DB() *sql.DB {
	return d.db.DB()
}

// Close ..
func (d *database) Close() {
	d.db.Close()
}

func (d *database) url() (string, error) {
	c := d.c
	s := ""

	switch d.c.dialect {
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
