package basicbuilder

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
