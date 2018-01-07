package gosora

import "../../common"

type GosoraMsgProducer struct {
}

func NewGosoraMsgProducer() *GosoraMsgProducer {
	return &GosoraMsgProducer{}
}

func (lex *GosoraMsgProducer) Run(msg string) *common.Ast {
	return nil
}
