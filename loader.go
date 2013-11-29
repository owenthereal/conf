package conf

type Loader struct {
	adapters    []Adapter
	defaultConf *Conf
}

func (l *Loader) Defaults(defs map[string]interface{}) *Loader {
	l.defaultConf.Merge(defs)
	return l
}

func (l *Loader) Env() *Loader {
	l.adapters = append(l.adapters, env{})
	return l
}

func (l *Loader) File(path string) *Loader {
	l.adapters = append(l.adapters, file{path})
	return l
}

func (l *Loader) Register(a Adapter) *Loader {
	l.adapters = append(l.adapters, a)
	return l
}

func (l *Loader) Load() (*Conf, error) {
	c := &Conf{make(map[string]interface{})}
	c.Merge(l.defaultConf.conf)
	for _, adapter := range l.adapters {
		err := adapter.Apply(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func NewLoader() *Loader {
	return &Loader{adapters: make([]Adapter, 0), defaultConf: &Conf{make(map[string]interface{})}}
}
