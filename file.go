package conf

import (
	"encoding/json"
	"io"
	"os"
)

type file struct {
	Path string
}

func (f file) Apply(conf *Conf) error {
	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	c, err := loadJSON(file)
	if err != nil {
		return err
	}

	conf.Merge(c)
	return nil
}

func loadJSON(r io.Reader) (map[string]interface{}, error) {
	var c map[string]interface{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&c)

	return c, err
}
