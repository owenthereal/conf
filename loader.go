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

// Define default configuration
func (l *Loader) Defaults(m map[string]interface{}) *Loader {
	l.defaultConf.Merge(m)
	return l
}

// Override configuration loaded from any Adapter
func (l *Loader) Overrides(m map[string]interface{}) *Loader {
	l.overrideConf.Merge(m)
	return l
}

// Load configuration from command line arguments
func (l *Loader) Argv() *Loader {
	l.adapters.Add(argv{os.Args})
	return l
}

// Load configuration from environment variables
func (l *Loader) Env() *Loader {
	l.adapters.Add(env{})
	return l
}

// Load configuration from a JSON file
func (l *Loader) File(path string) *Loader {
	l.adapters.Add(file{path})
	return l
}

// Register a custom Adapter for loading configuration
func (l *Loader) Register(a Adapter) *Loader {
	l.adapters.Add(a)
	return l
}

// Load configuration with registered Adapters
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

// Create a Loader
func NewLoader() *Loader {
	return &Loader{
		adapters:     &adapterList{make([]Adapter, 0)},
		defaultConf:  &Conf{make(map[string]interface{})},
		overrideConf: &Conf{make(map[string]interface{})},
	}
}
