package conf

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
		v, ok := val.(bool)
		if ok {
			return v
		}

		return false
	}

	return false
}

func (c *Conf) Int(key string) int {
	val := c.Get(key)
	if val != nil {
		v, ok := val.(int)
		if ok {
			return v
		}

		return 0
	}

	return 0
}

func (c *Conf) String(key string) string {
	val := c.Get(key)
	if val != nil {
		v, ok := val.(string)
		if ok {
			return v
		}

		return ""
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
