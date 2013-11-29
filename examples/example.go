package main

import (
	"fmt"
	"github.com/jingweno/conf"
	"os"
)

func main() {
	os.Setenv("GO_ENV", "development")
	c, err := conf.NewLoader().
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

	printConf(c, "GO_ENV")
	printConf(c, "DATABASE")
	printConf(c, "DATABASE_HOST")
	printConf(c, "DATABASE_PORT")
}

func printConf(c *conf.Conf, k string) {
	fmt.Printf("%s: %v\n", k, c.Get(k))
}
