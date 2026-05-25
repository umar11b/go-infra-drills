package healthcheck
type Checker interface {
	Check() error
	Name() string
}

type HTTPCheck struct {
	Check() error
}

type DiskCHeck struct {
	Checker interface{}
}