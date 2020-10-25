package conf

type DBConf struct {
	Host string
	Port int
	User string
	Password string
	DBName string
}

var MySQLConf = DBConf{
	Host: "127.0.0.1",
	Port: 3306,
	User: "root",
	Password: "12345678",
	DBName: "tagtalk",
}

