package gosora

import "../../common"

type PrealphaSoftware struct {
}

func NewPrealphaSoftware() *PrealphaSoftware {
	return &PrealphaSoftware{}
}

func (soft *PrealphaSoftware) Version() string {
	return "0.0.0"
}

func (soft *PrealphaSoftware) Producer() common.Producer {
	return NewGosoraMsgProducer()
}

// Bypass the parser for Gosora to Gosora migrations otherwise flatten it into HTML to make it more portable across Gosora installs
func (soft *PrealphaSoftware) Consumer() common.Consumer {
	return common.NewSimpleHTMLConsumer()
}
