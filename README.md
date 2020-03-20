# Autoconf

[![Build Status](https://travis-ci.com/chanjarster/autoconf.svg?branch=master)](https://travis-ci.com/chanjarster/autoconf.svg?branch=master)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

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

Noop:

```bash
./example
```

Load from yaml:

```bash
./example
```

Load from environment:

```bash
./example -c conf.yaml
```

Load from flag:

```bash
MYSQL_HOST=localhost ./example
```

Mix them together (flag takes the highest precedence, then environtment, then yaml file) :

```bash
MYSQL_PORT=3306 ./example -c conf.yaml -redis-host=localhost
```

## Supported types

Only the exported fields will be loaded. Support field types are:

* `string`
* `bool`
* `int`
* `int64`
* `uint`
* `uint64`
* `float64`
* `string`
* struct
* pointer to above types

## Name convention

 Autoconf load fields by convention. Example, field `BazBoom` path is `Foo.Bar.BazBoom`:

```go
type bar struct {
  BazBoom string
}
type foo struct {
  Bar *bar
}
type Conf struct {
  Foo *foo
}
```

So the corresponding:

* flag name is `-foo-bar-baz-boom`
* env name is `FOO_BAR_BAZ_BOOM`

As Autoconf use `gopkg.in/yaml.v3` to parse yaml, so the field name appear in yaml file should be lowercase, unless you customized it the "yaml" name in the field tag:

```yaml
foo:
  bar:
    bazroom: ...
```

