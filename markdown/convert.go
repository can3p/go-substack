package markdown

import (
	"fmt"

	"github.com/can3p/go-substack/types"
	"github.com/yuin/goldmark"
	goldmarkAst "github.com/yuin/goldmark/ast"
	goldmarkText "github.com/yuin/goldmark/text"
)

type parsedText struct {
	s      []byte
	node   goldmarkAst.Node
	parser goldmark.Markdown
}

func parse(s string) *parsedText {
	r := goldmarkText.NewReader([]byte(s))
	parser := goldmark.New()
	ast := parser.Parser().Parse(r)

	return &parsedText{
		s:      []byte(s),
		node:   ast,
		parser: parser,
	}
}

func ToDoc(mdText string) (*types.Node, error) {
	ast := parse(mdText)

	doc := types.NewNode(types.NTDoc)

	if err := translateAst(ast.s, doc, ast.node, 0, translateNode); err != nil {
		return nil, err
	}

	return doc, nil
}

type translater func(s []byte, parent *types.Node, n goldmarkAst.Node) (*types.Node, error)

func translateAst(source []byte, parent *types.Node, n goldmarkAst.Node, level int, translate translater) error {
	if n.ChildCount() > 0 {
		ch := n.FirstChild()

		for ch != nil {
			newNode, err := translate(source, parent, ch)

			if err != nil {
				return err
			}

			// no node means we don't need to go deeper,
			// the inspected node either has no meaning or has modified the parrent
			if newNode != nil {
				parent.Content = append(parent.Content, newNode)

				if err := translateAst(source, newNode, ch, level+1, translate); err != nil {
					return err
				}
			}

			ch = ch.NextSibling()
		}
	}

	return nil
}

func translateNode(s []byte, parent *types.Node, n goldmarkAst.Node) (*types.Node, error) {
	switch n.Kind().String() {
	case "Paragraph":
		return types.NewNode(types.NTParagraph), nil
	case "Text":
		newNode := types.NewNode(types.NTText)
		newNode.Text = string(n.Text(s))
		return newNode, nil
	}

	return nil, fmt.Errorf("Unknown node kind: %s, source text: [%s]", n.Kind().String(), string(n.Text(s)))
}
