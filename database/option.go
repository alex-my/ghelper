package database

// Config 配置文件
type config struct {
	username     string
	password     string
	host         string
	port         int
	dbname       string
	dialect      string
	maxIdleConns int
	maxOpenConns int
	logDebug     bool
}

func defaultConfig() *config {
	return &config{
		username:     "root",
		password:     "123456",
		host:         "127.0.0.1",
		port:         3306,
		dialect:      "mysql",
		maxIdleConns: 20,
		maxOpenConns: 50,
		logDebug:     false,
	}
}

// Option ..
type Option func(Database)

// WithUsername 数据库用户名
func WithUsername(username string) Option {
	return func(d Database) {
		d.config().username = username
	}
}

// WithPassword 数据库密码
func WithPassword(password string) Option {
	return func(d Database) {
		d.config().password = password
	}
}

// WithHost 数据库地址
func WithHost(host string) Option {
	return func(d Database) {
		d.config().host = host
	}
}

// WithPort 数据库端口号
func WithPort(port int) Option {
	return func(d Database) {
		d.config().port = port
	}
}

// WithDBName 数据库名称
func WithDBName(dbname string) Option {
	return func(d Database) {
		d.config().dbname = dbname
	}
}

// WithDialect 数据库 dialect
func WithDialect(dialect string) Option {
	return func(d Database) {
		d.config().dialect = dialect
	}
}

// WithMaxIdleConns 数据库 maxIdleConns
func WithMaxIdleConns(maxIdleConns int) Option {
	return func(d Database) {
		d.config().maxIdleConns = maxIdleConns
	}
}

// WithMaxOpenConns 数据库 maxOpenConns
func WithMaxOpenConns(maxOpenConns int) Option {
	return func(d Database) {
		d.config().maxOpenConns = maxOpenConns
	}
}

// WithLogDebug 是否打印debug日志
func WithLogDebug(logDebug bool) Option {
	return func(d Database) {
		d.config().logDebug = logDebug
	}
}
