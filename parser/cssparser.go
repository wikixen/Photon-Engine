package parser

import "unicode"

type (
	Stylesheet struct {
		rules []Rules
	}

	Rules struct {
		selectors    []Selector
		declarations []Declaration
	}

	Selector interface {
		Selector(SimpleSelector)
	}

	SimpleSelector struct {
		tagName, id string
		class       []string
	}

	Declaration struct {
		name  string
		value Value
	}

	Value interface {
		Keyword(string)
		Length(float32, Unit)
		ColorVal(Color)
	}

	Unit struct {
		Px, Em, Rem int8
	}

	Color struct {
		r, g, b, a uint8
	}
)

func (p *Parser) parseSimpleSelector() SimpleSelector {
	selector := SimpleSelector{
		tagName: "",
		id:      "",
		class:   []string{},
	}

	for !p.endOfS() {
		switch p.nextChar() {
		case '#':
			p.consumeChar()
			selector.id = ""
		case '.':
			p.consumeChar()
			selector.class = append(selector.class, "")
		case '*':
			p.consumeChar()
		default:
			break
		}
	}
	return selector
}

func validIDChar(c byte) bool {
	if unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) || c == '_' {
		return true
	} else {
		return false
	}
}
