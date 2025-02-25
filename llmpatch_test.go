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

func TestSliceIndex(t *testing.T) {
	tests := []struct {
		s      []string
		search []string
		index  int
	}{
		{
			s:      []string{"a"},
			search: []string{"a"},
			index:  0,
		},
		{
			s:      []string{"b", "b", "c", "c"},
			search: []string{"c", "c"},
			index:  2,
		},
		{
			s:      []string{"b", "b", "c", "c"},
			search: []string{"a", "a"},
			index:  -1,
		},
		{
			s:      []string{"b", "b", "c", "c"},
			search: []string{"b", "b"},
			index:  0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := sliceIndex(tt.s, tt.search)
			if got != tt.index {
				t.Fatalf("got %v, want %v", got, tt.index)
			}
		})
	}
}

func TestApply(t *testing.T) {
	tests := []struct {
		input string
		edits []Edit
		want  string
	}{
		{
			input: "a",
			edits: []Edit{
				{
					Search:  "a",
					Replace: "b",
				},
			},
			want: "b",
		},
		{
			input: "1\na\na",
			edits: []Edit{
				{
					Search:  "a\na",
					Replace: "b\nb",
				},
			},
			want: "1\nb\nb",
		},
		{
			input: "a\n",
			want:  "a\n",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Apply(tt.input, tt.edits); got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
