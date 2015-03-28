package confluence

type Confluence interface {
	Foo()
}

type DefaultConfluence struct {
	BaseURL  string
	username string
	password string
	Confluence
}

func (d DefaultConfluence) Foo() {

}
