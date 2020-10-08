package main

import (
	"fmt"
	builder "learning-go/pkg/builder/basicbuilder"
	"strings"
)

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
	div := builder.NewHTMLBuilder("div").
		AddChild("div", "header").
		AddChild("div", "content").
		AddChild("div", "footer")

	fmt.Println(div.String())
}

func main() {
	withoutBuilder()
	withBuilder()
}
