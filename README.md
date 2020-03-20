# Autoconf

A simple tool to load config from yaml, environment variables and flags. 

## Install

```bash
go get github.com/chanjarster/autoconf
```

## Example

```go
package main

import (
  "fmt"
  "github.com/chanjarster/autoconf"
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

func main() {
  conf := &Conf{}
  autoconf.Load(conf, "c")
  fmt.Println(conf)
}
```

conf.yaml:

```yaml
mysql:
  host: localhost
  port: 3306
  database: test
redis:
  host: localhost
  port: 6379
  password: test
```

Run:

```bash
# Load from nothing
./example
# Load from yaml
./example -c conf.yaml
# Load from env
MYSQL_HOST=localhost ./example
# Load from flag
./example -mysql-host=localhost
# Or put them together
MYSQL_PORT=3306 ./example -c conf.yaml -redis-host=localhost
```

## Supported types

Complete list:

* `string`
* `bool`
* `int`
* `int64`
* `uint`
* `uint64`
* `float64`
* `string`
* struct
* pointers to above types

Only the exported fields will be loaded.

## Name convention

 Autoconf load fields by convention. Example, for field path `Foo.Bar.BazBoom`:

* flag name will be `-foo-bar-baz-boom`
* env name will be `FOO_BAR_BAZ_BOOM`

As Autoconf use `gopkg.in/yaml.v3`, so the field name appear in yaml file should be lowercase, unless you customized it the "yaml" name in the field tag:

```yaml
foo:
  bar:
    bazroom: ...
```

