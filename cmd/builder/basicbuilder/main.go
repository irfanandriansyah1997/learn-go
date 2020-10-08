package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))
	return sb.String()
}

type HTMLBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHTMLBuilder(rootName string) *HTMLBuilder {
	b := HTMLBuilder{
		rootName,
		HtmlElement{
			rootName,
			"",
			[]HtmlElement{},
		},
	}

	return &b
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

func (b *HTMLBuilder) AddChild(
	childName, childText string,
) *HTMLBuilder {
	e := HtmlElement{
		childName,
		childText,
		[]HtmlElement{},
	}
	b.root.elements = append(b.root.elements, e)

	return b
}

func withoutBuilder() {
	sb := strings.Builder{}
	words := []string{"hello", "world", "Irfan"}

	sb.WriteString("<ul>")
	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())
}

func withBuilder() {
	div := NewHTMLBuilder("div").
		AddChild("div", "header").
		AddChild("div", "content").
		AddChild("div", "footer")

	fmt.Println(div.String())
}

func main() {
	withoutBuilder()
	withBuilder()
}
