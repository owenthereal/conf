package conf

import (
	"os"
	"strings"
)

type env struct {
}

func (e env) Apply(conf *Conf) error {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		conf.Set(pair[0], pair[1])
	}

	return nil
}
