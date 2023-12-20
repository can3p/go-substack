package markdown

import (
	"encoding/json"
	"testing"

	"github.com/can3p/go-substack/types"
)

func TestToDoc(t *testing.T) {
	type testCase struct {
		description string
		input       string
		output      *types.Node
		err         error
	}

	examples := []testCase{
		{
			description: "simplest case",
			input:       `test`,
			output: &types.Node{
				Type: types.NTDoc,
				Content: []*types.Node{
					{
						Type: types.NTParagraph,
						Content: []*types.Node{
							{
								Type: types.NTText,
								Text: "test",
							},
						},
					},
				},
			},
		},
	}

	for idx, ex := range examples {
		out, err := ToDoc(ex.input)

		if (err == nil) != (ex.err == nil) {
			t.Errorf("Example [%d, %s]: expected err [%v] but got [%v]", idx+1, ex.description, ex.err, err)
			continue
		}

		if err != nil {
			if err.Error() != ex.err.Error() {
				t.Errorf("Example [%d, %s]: error mismatch - expected [%v] but got [%v]", idx+1, ex.description, ex.err, err)
			}
			continue
		}

		if out == nil {
			t.Errorf("Example [%d, %s]: no error got returned, out put is nil as well", idx+1, ex.description)
		}

		expected, _ := json.Marshal(ex.output)
		got, _ := json.Marshal(out)

		if string(expected) != string(got) {
			t.Errorf("Example [%d, %s]: expected and received output do not match", idx+1, ex.description)
			t.Errorf("Example [%d, %s]: expected: %s", idx+1, ex.description, string(expected))
			t.Errorf("Example [%d, %s]: got     : %s", idx+1, ex.description, string(got))
		}
	}
}
