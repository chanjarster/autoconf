package autoconf

import (
	"fmt"
)

type Conf struct {
	Mysql *MysqlConf
	Redis *RedisConf
}

type MysqlConf struct {
	Host string
	Port int
	Database string
}

type RedisConf struct {
	Host string
	Port int
	Password string
}

func Example() {
	conf := &Conf{}
	Load(conf, "c")
	fmt.Println(conf)
}
