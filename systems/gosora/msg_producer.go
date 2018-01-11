package gosora

import "../../common"

type GosoraMsgProducer struct {
}

func NewGosoraMsgProducer() *GosoraMsgProducer {
	return &GosoraMsgProducer{}
}

// TODO: BBCode
// TODO: HTML
// TODO: Markdown. Let's forget this exists for now, I'm thinking of crunching this down to HTML on Gosora's side anyway.
func (lex *GosoraMsgProducer) Run(msg string) common.Node {
	var runes = []rune(msg)
	var ast = &common.Ast{}
	parse(ast, runes, 0, len(runes))
	return ast
}

func parse(tree common.Node, runes []rune, i int, maxLen int) {
	var buffer string
	
	var stepForward = func(i int, step int, maxLen int) int {
		i += step
		if i < maxLen {
			return i
		}
		return i - step
	}

	var zoomToMatch = func(i int, phrase []rune, runes []rune, maxLen int) int {
		OuterZoom:
		for ; i < maxLen; i++ {
			if phrase[0] == runes[i] {
				i++
				if (i + len(phrase) >= maxLen {
					continue OuterZoom
				}
				for pi := 1; pi < len(phrase); pi++ {
					if phrase[pi] != runes[i] {
						continue OuterZoom
					}
					stepForward(i,1,maxLen)
				}
				break
			}
		}
		return i
	}

	var addChild = func(node Node, skip int, maxLen int) {
		tree.AddChild(&common.Text{Body:buffer})
		buffer = ""
		i += skip
		parse(node,runes,i,maxLen)
		tree.AddChild(node)
	}
	
	OuterParse:
	for ; i < maxLen; i++ {
		char := runes[i]
		if char == '[' {
			switch {
			case peekMatch(i, "b]", runes):
				locMax := zoomToMatch(i, []rune("[/b]"), runes, maxLen)
				addChild(&common.Bold{},2,locMax)
				continue OuterParse
			case peekMatch(i, "i]", runes):
				locMax := zoomToMatch(i, []rune("[/i]"), runes, maxLen)
				addChild(&common.Italic{},2,locMax)
				continue OuterParse
			case peekMatch(i, "u]", runes):
				locMax := zoomToMatch(i, []rune("[/u]"), runes, maxLen)
				addChild(&common.Underline{},2,locMax)
				continue OuterParse
			case peekMatch(i, "s]", runes):
				locMax := zoomToMatch(i, []rune("[/s]"), runes, maxLen)
				addChild(&common.Strikethrough{},2,locMax)
				continue OuterParse
			}
		}
		buffer += string(char)
	}
}

// TODO: Test this
func peek(cur int, skip int, runes []rune) rune {
	if (cur + skip) < len(runes) {
		return runes[cur+skip]
	}
	return 0 // null byte
}

// TODO: Test this
func peekMatch(cur int, phrase string, runes []rune) bool {
	if cur+len(phrase) > len(runes) {
		return false
	}
	for i, char := range phrase {
		if runes[cur+i+1] != char {
			return false
		}
	}
	return true
}
