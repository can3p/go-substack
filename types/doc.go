package types

type NodeType string

const (
	NTDoc            NodeType = "doc"
	NTParagraph      NodeType = "paragraph"
	NTHeading        NodeType = "heading"
	NTCaptionedImage NodeType = "captionedImage"
	NTImage2         NodeType = "image2"
	NTCaption        NodeType = "caption"
	NTFootnote       NodeType = "footnote"
	NTText           NodeType = "text"
	NTFootnoteAnchor NodeType = "footnoteAnchor"
)

type MarkType string

const (
	MarkStrong = "strong"
)

type Mark struct {
	Type MarkType `json:"type"`
}

type Node struct {
	Type NodeType `json:"type"`

	// for container nodes
	Content []*Node `json:"content,omitempty"`

	// some types have args
	Args map[string]string `json:"args,omitempty"`

	// for inline nodes
	Text  string  `json:"text,omitempty"`
	Marks []*Mark `json:"marks,omitempty"`
}
