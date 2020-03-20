package autoconf

type ConfResolver interface {
	Resolve(conf interface{}) error
}
