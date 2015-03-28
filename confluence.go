package confluence

import (
	"fmt"
	"net/url"
)

type Confluence interface {
	Foo()
}

type DefaultConfluence struct {
	baseURL  string
	username string
	password string
	Confluence
}

func NewConfluenceClient(baseURL string, username, password string) (Confluence, error) {
	if _, err := url.Parse(baseURL); err != nil {
		return DefaultConfluence{}, err
	}
	return DefaultConfluence{baseURL: baseURL, username: username, password: password}, nil
}

func (d DefaultConfluence) Foo() {
	fmt.Printf("hello, foo\n")
}
