package parser

import (
	"FennecBrowser/dom"
	"log"
	"unicode"
	"unicode/utf8"
)

// Parser stores the input string & the current pos in the string where pos is the next unprocessed char
type Parser struct {
	pos   uint
	input string
}

// startsWith returns true if current pos is equal to the provided string
func (p *Parser) startsWith(s string) bool {
	return p.input[p.pos:p.pos] == s
}

// expect checks if s is equal to the current pos & panics if not
func (p *Parser) expect(s string) {
	if p.startsWith(s) {
		p.pos += uint(len(s))
	} else {
		log.Panicf("Expected %s at byte %d\n", s, p.pos)
	}
}

// endOfS returns true if all the input has been read
func (p *Parser) endOfS() bool {
	return p.pos >= uint(len(p.input))
}

// nextChar reads the current char
func (p *Parser) nextChar() byte {
	return p.input[p.pos]
}

// Might remove this & change the way func works
// consumeChar returns current char & increments pos
func (p *Parser) consumeChar() byte {
	c := p.nextChar()
	p.pos += uint(utf8.RuneLen(rune(c)))
	return c
}

// consumeWhile consumes chars until test returns false
func (p *Parser) consumeWhile(test func(rune) bool) string {
	var res string
	for !p.endOfS() && test(rune(p.nextChar())) {
		res += string(p.nextChar())
	}
	return res
}

// consumeWhitespace discards whitespace
func (p *Parser) consumeWhitespace() {
	p.consumeWhile(unicode.IsSpace)
}

// parseName parses tag & attribute names
func (p *Parser) parseName() string {
	return p.consumeWhile(func(c rune) bool {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			return true
		} else {
			return false
		}
	})
}

func (p *Parser) parseNode() dom.Node {
	if p.startsWith("<") {
		return p.parseElement()
	} else {
		return p.parseText()
	}
}

func (p *Parser) parseText() dom.Node {
	return dom.Text(p.consumeWhile(func(c rune) bool {
		if c != '<' {
			return false
		} else {
			return true
		}
	}))
}

func (p *Parser) parseElement() dom.Node {
	var attr dom.AttrMap

	// Opening tag
	p.expect("<")
	tagName := p.parseName()
	key, val := p.parseAttr()
	attr[key] = val
	p.expect(">")

	children := p.parseNodes()

	// Closing tag
	p.expect("</")
	p.expect(tagName)
	p.expect(">")

	return dom.Element(tagName, attr, children)
}

func (p *Parser) parseAttr() (string, string) {
	name := p.parseName()
	p.expect("=")
	val := p.parseAttrVal()

	return name, val
}

func (p *Parser) parseAttrVal() string {
	openQuote := p.consumeChar()
	if openQuote == '"' || openQuote == '\'' {
		log.Panicf("Expected %c", openQuote)
	}
	val := p.consumeWhile(func(c rune) bool {
		if c != rune(openQuote) {
			return true
		} else {
			return false
		}
	})
	closeQuote := p.consumeChar()
	if closeQuote != openQuote {
		log.Panicf("Expected %c", closeQuote)
	}
	return val
}

// parseAttrs parses name="val" pairs
func (p *Parser) parseAttrs() dom.AttrMap {
	attr := make(map[string]string)

	for {
		p.consumeWhitespace()
		if p.nextChar() == '>' {
			break
		}
		name, val := p.parseAttr()
		attr[name] = val
	}
	return attr
}

func (p *Parser) parseNodes() []dom.Node {
	nodes := make([]dom.Node, 0)
	for {
		p.consumeWhitespace()
		if p.endOfS() || p.startsWith("</") {
			break
		}
		nodes = append(nodes, p.parseNode())
	}
	return nodes
}

func parse(source string) dom.Node {
	parser := Parser{pos: 0, input: source}
	var nodes = parser.parseNodes()
	if len(nodes) == 1 {
		return nodes[0]
	} else {
		return dom.Element("html", make(map[string]string), nodes)
	}
}
