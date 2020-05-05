package template

import "strings"

type MessageRetriever interface {
	Message() string
}

type Templater interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type Template struct{}

func (t *Template) first() string {
	return "hello"
}

func (t *Template) third() string {
	return "template"
}

func (t *Template) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}

type AnonymousTemplate struct{}

func (t *AnonymousTemplate) first() string {
	return "hello"
}

func (t *AnonymousTemplate) third() string {
	return "template"
}

func (t *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{t.first(), f(), t.third()}, " ")
}

type adapter struct {
	myFunc func() string
}

func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}
