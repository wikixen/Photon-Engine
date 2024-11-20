package style

import (
	"FennecBrowser/dom"
	"FennecBrowser/parser"
)

type propertyMap = map[string]parser.Value

type StyledNode struct {
	node     *dom.Node
	specVal  propertyMap
	children []*StyledNode
}

func matches(elem *dom.ElementData, selector *parser.Selector) bool {
	switch selector {
	case s:

	}
}
