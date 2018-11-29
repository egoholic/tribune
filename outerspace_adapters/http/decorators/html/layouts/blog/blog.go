package blog

import (
	"strings"
)

type Layout struct {
	builder *strings.Builder
	title   *func(*strings.Builder) error
	content *func(*strings.Builder) error
}

func Build(title, content *func(*strings.Builder) error) *Layout {
	sb := &strings.Builder{}
	return &Layout{sb, title, content}
}

func (l *Layout) Render() string {
	l.builder.WriteString(`<!DOCTYPE html>`)
	l.builder.WriteString(`<html><head><title>`)
	_ = (*l.title)(l.builder) // TODO: handle error
	l.builder.WriteString(`</title></head><body>`)
	l.builder.WriteString(`<body>`)
	_ = (*l.content)(l.builder) // TODO: handle error
	l.builder.WriteString(`</body></html>`)
	return l.builder.String()
}
