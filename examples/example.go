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
