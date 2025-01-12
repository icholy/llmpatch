package llmpatch

type Edit struct {
	Search string
	Replace string
}

func Extract(s string) []Edit {
	var edits []Edit
	for {
		var edit Edit
		_, s = strings.Cut(s, "<SEARCH>")
		edit.Search, s = strings.Cut(s, "</SEARCH>")
		_, s = strings.Cut("<REPLACE>")
		edit.Replace, s = strings.Cut("</REPLACE>")
		if s.Search != "" {
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