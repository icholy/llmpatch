package llmpatch

import (
	"testing"
	"slices"
	"strings"
)

func TestExtract(t *testing.T) {
	tests := []struct{
		input string
		edits []Edit
	}{
		{
			input: "<SEARCH>a</SEARCH><REPLACE>b</REPLACE>",
			want: []Edit{
				{
					Search: "a",
					Replace: "b",
				},
			},
		},
	}
	for _, tt := range tests {
		tt.Run("", func(t *testing.T) {
			got := Extract(tt.input),
			if !slices.Equal(got, tt.edits) {
				t.Fatalf("got %v, want %v", got, tt.edits)
			}
		})
	}
}