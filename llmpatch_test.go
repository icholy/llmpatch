package llmpatch

import (
	"slices"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		input string
		edits []Edit
	}{
		{
			input: "<SEARCH>a</SEARCH><REPLACE>b</REPLACE>",
			edits: []Edit{
				{
					Search:  "a",
					Replace: "b",
				},
			},
		},
		{
			input: "<SEARCH>\na\n</SEARCH>\n<REPLACE>\nb\n</REPLACE>\n",
			edits: []Edit{
				{
					Search:  "a",
					Replace: "b",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Extract(tt.input)
			if !slices.Equal(got, tt.edits) {
				t.Fatalf("got %v, want %v", got, tt.edits)
			}
		})
	}
}
