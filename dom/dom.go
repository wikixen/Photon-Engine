package dom

type (
	// Node represents a node on the DOM tree; A Node can be of either of 2 types: Text or Element
	Node struct {
		children []Node
		nodeType NodeType
	}

	NodeType struct {
		Text    string
		Element ElementData
	}

	ElementData struct {
		tagName string
		attrs   AttrMap
	}

	// AttrMap stores all attributes
	AttrMap map[string]string
)

// Text creates a text node
func Text(data string) Node {
	var c []Node
	return Node{children: c, nodeType: NodeType{Text: data}}
}

// Element creates an element node
func Element(tag string, attr AttrMap, children []Node) Node {
	return Node{
		children: children,
		nodeType: NodeType{Element: ElementData{tagName: tag, attrs: attr}},
	}
}
