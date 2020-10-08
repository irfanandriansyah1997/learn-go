package basicbuilder

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
