package conf

type adapterSet struct {
	set map[Adapter]bool
}

func (s *adapterSet) Add(a Adapter) bool {
	_, found := s.set[a]
	s.set[a] = true
	return !found
}

type Loader struct {
	adapters    *adapterSet
	defaultConf *Conf
}

func (l *Loader) Defaults(defs map[string]interface{}) *Loader {
	l.defaultConf.Merge(defs)
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
	for a, _ := range l.adapters.set {
		err := a.Apply(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func NewLoader() *Loader {
	return &Loader{
		adapters:    &adapterSet{make(map[Adapter]bool)},
		defaultConf: &Conf{make(map[string]interface{})},
	}
}
