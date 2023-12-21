package types

import (
	"encoding/json"
	"testing"
)

func TestDraftBodyMarshalling(t *testing.T) {
	doc := NewNode(
		NTDoc,
		NewNode(NTParagraph,
			&Node{
				Type: NTText,
				Text: "first",
			}),
	)

	draft := Draft{
		DraftBody: DraftBody(*doc),
	}

	b, err := json.Marshal(draft)

	if err != nil {
		t.Fatalf("Failed to encode draft: %v", err.Error())
	}

	var decoded Draft

	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("Failed to decode draft: %v", err.Error())
	}

	expected, _ := json.Marshal(doc)
	got, _ := json.Marshal(Node(decoded.DraftBody))

	if string(expected) != string(got) {
		t.Errorf("expected and received output do not match")
		t.Errorf("expected: %s", string(expected))
		t.Errorf("got     : %s", string(got))
	}
}
