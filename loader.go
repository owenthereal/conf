package conf

import (
	"os"
)

type adapterList struct {
	list []Adapter
}

func (s *adapterList) Add(a Adapter) bool {
	for _, adapter := range s.list {
		if adapter == a {
			return false
		}
	}

	s.list = append(s.list, a)
	return true
}

type Loader struct {
	adapters     *adapterList
	defaultConf  *Conf
	overrideConf *Conf
}

func (l *Loader) Defaults(defs map[string]interface{}) *Loader {
	l.defaultConf.Merge(defs)
	return l
}

func (l *Loader) Overrides(defs map[string]interface{}) *Loader {
	l.overrideConf.Merge(defs)
	return l
}

func (l *Loader) Argv() *Loader {
	l.adapters.Add(argv{os.Args})
	return l
}

func (l *Loader) Env() *Loader {
	l.adapters.Add(env{})
	return l
}

func (l *Loader) File(path string) *Loader {
	l.adapters.Add(file{path})
	return l
}

func (l *Loader) Register(a Adapter) *Loader {
	l.adapters.Add(a)
	return l
}

func (l *Loader) Load() (*Conf, error) {
	c := &Conf{make(map[string]interface{})}
	c.Merge(l.defaultConf.conf)
	for _, a := range l.adapters.list {
		err := a.Apply(c)
		if err != nil {
			return nil, err
		}
	}
	c.Merge(l.overrideConf.conf)

	return c, nil
}

func NewLoader() *Loader {
	return &Loader{
		adapters:     &adapterList{make([]Adapter, 0)},
		defaultConf:  &Conf{make(map[string]interface{})},
		overrideConf: &Conf{make(map[string]interface{})},
	}
}
