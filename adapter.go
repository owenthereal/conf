package conf

type Adapter interface {
	Apply(conf *Conf) error
}
