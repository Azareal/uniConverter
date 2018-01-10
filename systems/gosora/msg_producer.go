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
func (lex *GosoraMsgProducer) Run(msg string) common.Tree {
	var runes = []rune(msg)
	var ast = &common.Ast{}
	parse(ast, runes, 0, len(runes))
	return ast
}

func parse(tree common.Tree, runes []rune, i int, maxLen int) {
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
					continue
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

	for ; i < maxLen; i++ {
		char := runes[i]
		if char == '[' {
			switch {
			case peekMatch(i, "b]", runes):
				i += 2
				locMax := zoomToMatch(i, []rune("[/b]"), runes, maxLen)
				boldNode := &common.Bold{}
				parse(boldNode,runes,i,locMax)
				tree.AddChild(boldNode)
			case peekMatch(i, "i]", runes):
				i += 2
				locMax := zoomToMatch(i, []rune("[/i]"), runes, maxLen)
				italicNode := &common.Italic{}
				parse(italicNode,runes,i,locMax)
				tree.AddChild(italicNode)
			case peekMatch(i, "u]", runes):
				i += 2
				locMax := zoomToMatch(i, []rune("[/u]"), runes, maxLen)
				underNode := &common.Underline{}
				parse(underNode,runes,i,locMax)
				tree.AddChild(underNode)
			case peekMatch(i, "s]", runes):
				i += 2
				locMax := zoomToMatch(i, []rune("[/s]"), runes, maxLen)
				strikeNode := &common.Strikethrough{}
				parse(strikeNode,runes,i,locMax)
				tree.AddChild(strikeNode)
			}
		}
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
