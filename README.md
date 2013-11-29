Conf
====

A configuration loader in Go. Load configuration with files, environment variables and command-line arguments.

## Example

Consider the following Go code in a file called `example.go`:

```go
import (
	"fmt"
	"github.com/jingweno/conf"
	"os"
)

func main() {
	os.Setenv("GO_ENV", "development")
	conf, err := conf.NewLoader().
		Env().
		File("./config.json").
		Defaults(
		map[string]interface{}{
			"DATABASE_HOST": "127.0.0.1",
			"DATABASE_PORT": "1234",
		}).
		Load()

	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}

	printConf(conf, "GO_ENV")
	printConf(conf, "DATABASE")
	printConf(conf, "DATABASE_HOST")
	printConf(conf, "DATABASE_PORT")
}

func printConf(conf *conf.Conf, key string) {
	fmt.Printf("%s: %v\n", key, conf.Get(key))
}
```

If you run the above code:

```plain
$ DATABASE_PORT=5678 go run example.go
```

The output will be:

```plain
GO_ENV: development
DATABASE: postgres
DATABASE_HOST: foo
DATABASE_PORT: 1234
```
