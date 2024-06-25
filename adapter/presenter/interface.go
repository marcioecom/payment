package presenter

type Presenter interface {
	Show() ([]byte, error)
	Bind(any) error
}
