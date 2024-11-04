package dom

type (
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

	AttrMap map[string]string
)

func Text(data string) Node {
	var c []Node
	return Node{children: c, nodeType: NodeType{Text: data}}
}

func Element(tag string, attr AttrMap, children []Node) Node {
	return Node{
		children: children,
		nodeType: NodeType{Element: ElementData{tagName: tag, attrs: attr}},
	}
}
