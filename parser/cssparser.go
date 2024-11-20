package parser

import "unicode"

type (
	// Stylesheet respresents a CSS stylesheet
	Stylesheet struct {
		rules []Rules
	}

	// Rules are CSS code brackets written like so: Selector { Declaration; }
	Rules struct {
		selectors    []Selector
		declarations []Declaration
	}

	// Selector can be either '#', '.', or 'tagName' and represent an attribute or element from HTML
	Selector interface {
		Simple(SimpleSelector)
	}

	SimpleSelector struct {
		tagName, id string
		class       []string
	}

	// A Declaration is CSS code within brackets
	Declaration struct {
		name  string
		value Value
	}

	// Value are possible values for CSS declarations
	Value interface {
		Keyword(string)
		Length(float32, Unit)
		ColorVal(Color)
	}

	// Unit for length values
	Unit struct {
		Px int8
	}

	// Color represented by RGBA
	Color struct {
		r, g, b, a uint8
	}

	Specificity [3]*uint
)

// parseSimpleSelector parses a single selector
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

// validIDChar checks that input is a valid char
func validIDChar(c byte) bool {
	if unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) || c == '_' {
		return true
	} else {
		return false
	}
}
