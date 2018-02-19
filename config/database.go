package config

var DbConfig  = map[string]interface{}{
	"default":         "mysql_dev", // 默认数据库配置
	"SetMaxOpenConns": 300,         // (连接池)最大打开的连接数，默认值为0表示不限制
	"SetMaxIdleConns": 10,          // (连接池)闲置的连接数, 默认-1
	"mysql_dev": map[string]string{ // 定义名为 mysql_dev 的数据库配置
		"host":     "192.168.10.10", // 数据库地址
		"username": "homestead",           // 数据库用户名
		"password": "secret",           // 数据库密码
		"port":     "3306",            // 端口
		"database": "xz_local",            // 链接的数据库名字
		"charset":  "utf8",            // 字符集
		"protocol": "tcp",             // 链接协议
		"prefix":   "",                // 表前缀
		"driver":   "mysql",           // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
	},
}
