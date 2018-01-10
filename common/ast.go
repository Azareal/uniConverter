package common

type Ast struct {
	Nodes []Node
}

type Node interface {
	AddChild(node Node)
}

type Text struct {
	Body string
}

func (text *Text) AddChild(_ Node) {
}

type Bold struct {
	Nodes []Node
}

func (bold *Bold) AddChild(node Node) {
	bold.Nodes = append(bold.Nodes, node)
}

type Italic struct {
	Nodes []Node
}

func (italic *Italic) AddChild(node Node) {
	italic.Nodes = append(italic.Nodes, node)
}

type Strikethrough struct {
	Nodes []Node
}

func (strike *Strikethrough) AddChild(node Node) {
	strike.Nodes = append(strike.Nodes, node)
}

type Underline struct {
	Nodes []Node
}

func (under *Underline) AddChild(node Node) {
	under.Nodes = append(under.Nodes, node)
}

type List struct {
	Nodes []Node
}

func (list *List) AddChild(node Node) {
	list.Nodes = append(list.Nodes, node)
}
