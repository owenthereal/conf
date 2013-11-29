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
