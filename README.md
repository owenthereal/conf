Conf
====

A configuration loader in Go. Load configuration with files, environment variables and command-line arguments.

## Example

Consider the following Go code:

```go
package main

import (
	"fmt"
	"github.com/jingweno/conf"
	"os"
)

func main() {
	d := map[string]interface{}{
		"GO_ENV":        "development",
		"DATABASE_NAME": "example_development",
		"DATABASE_POOL": 5,
	}
	c, err := conf.NewLoader().Argv().Env().File("./config.json").Defaults(d).Load()

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %s\n", err)
		return
	}

	printConf(c, "GO_ENV")
	printConf(c, "DATABASE")
	printConf(c, "DATABASE_NAME")
	printConf(c, "DATABASE_HOST")
	printConf(c, "DATABASE_PORT")
	printConf(c, "DATABASE_POOL")
}

func printConf(c *conf.Conf, k string) {
	fmt.Printf("%s: %v\n", k, c.Get(k))
}
```
and a `config.json`:

```json
{
  "DATABASE": "postgres",
  "DATABASE_HOST": "127.0.0.1",
  "DATABASE_PORT": "1234"
}
```

If you run the above code:

```plain
$ GO_ENV=production go run example.go --DATABASE_POOL 10
```

The output will be:

```plain
GO_ENV: production
DATABASE: postgres
DATABASE_NAME: example_development
DATABASE_HOST: 127.0.0.1
DATABASE_PORT: 1234
DATABASE_POOL: 10
```

## Load Configuration from Multiple Sources

Configuration management can get complicated very quickly for even trivial applications running in production.
`conf` addresses this problem by enabling you to load configuration from different sources.
**The order in which you attach these configuration sources determines their priority.**
These are the supported sources:

* [Loader.Argv](http://godoc.org/github.com/jingweno/conf#Loader.Argv) loads configuration from command line arguments
* [Loader.Env](http://godoc.org/github.com/jingweno/conf#Loader.Env) loads configuration from environment variables
* [Loader.File](http://godoc.org/github.com/jingweno/conf#Loader.File) loads configuration from a JSON file
* [Loader.Defaults](http://godoc.org/github.com/jingweno/conf#Loader.Defaults) defines default configuration
* [Loader.Overrides](http://godoc.org/github.com/jingweno/conf#Loader.Overrides) overrides configuration loaded from any sources
* [Loader.Register](http://godoc.org/github.com/jingweno/conf#Loader.Register) takes a custom [Adapter](http://godoc.org/github.com/jingweno/conf#Adapter) for loading configuration
