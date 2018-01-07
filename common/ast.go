package common

type Ast struct {
	Nodes []Node
}

type Node interface {
}

type Text struct {
	Body string
}

type DecoratedText struct {
	Body string

	Bold          bool
	Italic        bool
	Strikethrough bool
	Underline     bool
}

type List struct {
}
