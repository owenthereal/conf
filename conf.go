package conf

import (
	"strconv"
	"strings"
)

type Conf struct {
	conf map[string]interface{}
}

func (c *Conf) Size() int {
	return len(c.conf)
}

func (c *Conf) Get(key string) interface{} {
	return c.conf[key]
}

func (c *Conf) Bool(key string) bool {
	val := c.Get(key)
	if val != nil {
		switch val := val.(type) {
		default:
			return false
		case string:
			return strings.ToLower(strings.TrimSpace(val)) == "true"
		case bool:
			return val
		}
	}

	return false
}

func (c *Conf) Int(key string) int {
	val := c.Get(key)
	if val != nil {
		switch val := val.(type) {
		default:
			return 0
		case string:
			i, err := strconv.Atoi(val)
			if err != nil {
				return 0
			}

			return i
		case float64:
			return int(val)
		case int:
			return val
		}
	}

	return 0
}

func (c *Conf) String(key string) string {
	val := c.Get(key)
	if val != nil {
		switch val := val.(type) {
		default:
			return ""
		case string:
			return val
		case int:
			return strconv.Itoa(val)
		case bool:
			if val {
				return "true"
			} else {
				return "false"
			}
		}
	}

	return ""
}

func (c *Conf) Set(key string, val interface{}) interface{} {
	c.conf[key] = val
	return val
}

func (c *Conf) Merge(m map[string]interface{}) {
	for k, v := range m {
		c.Set(k, v)
	}
}
