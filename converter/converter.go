package converter

import "../common"
import "../systems/gosora"

func Lookup(software string, version string) (soft common.Software, exists bool) {
	switch software {
	case "gosora":
		return gosora.Lookup(version)
	}
	return soft, false
}

type Converter struct {
	From common.Software
	To   common.Software
}

func NewConverter() *Converter {
	return &Converter{}
}

func (conv *Converter) From(soft common.Software) {
	conv.From = soft
}

func (conv *Converter) To(soft common.Software) {
	conv.To = soft
}
