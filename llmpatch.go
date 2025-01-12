package llmpatch

import (
	_ "embed"
	"strings"
)

//go:embed prompt.md
var Prompt string

type Edit struct {
	Search  string
	Replace string
}

func Extract(s string) []Edit {
	var edits []Edit
	for {
		var edit Edit
		var ok bool
		if _, s, ok = strings.Cut(s, "<SEARCH>"); !ok {
			break
		}
		if edit.Search, s, ok = strings.Cut(s, "</SEARCH>"); !ok {
			break
		}
		if _, s, ok = strings.Cut(s, "<REPLACE>"); !ok {
			break
		}
		if edit.Replace, s, ok = strings.Cut(s, "</REPLACE>"); !ok {
			break
		}
		if edit.Search != "" {
			edits = append(edits, edit)
		}
	}
	return edits
}

func Apply(s string, edits []Edit) string {
	for _, e := range edits {
		s = strings.ReplaceAll(s, e.Search, e.Replace)
	}
	return s
}
