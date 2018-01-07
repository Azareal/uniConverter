package common

type Software interface {
	Version() string
	Producer() Producer
	Consumer() Consumer
}

// Producer consumes a string and spits out an AST
type Producer interface {
	Run(msg string) *Ast
}

// Consumer consumes an AST and spits out a string
type Consumer interface {
	Run(tree *Ast) string
}
