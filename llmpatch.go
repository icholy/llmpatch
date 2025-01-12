package llmpatch

import (
	"bufio"
	_ "embed"
	"fmt"
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
	scanner := bufio.NewScanner(strings.NewReader(s))
	for {
		var ok bool
		var edit Edit
		if _, ok = scanUntil(scanner, "<SEARCH>"); !ok {
			break
		}
		if edit.Search, ok = scanUntil(scanner, "</SEARCH>"); !ok {
			break
		}
		if _, ok = scanUntil(scanner, "<REPLACE>"); !ok {
			break
		}
		if edit.Replace, ok = scanUntil(scanner, "</REPLACE>"); !ok {
			break
		}
		if edit.Search != "" {
			edits = append(edits, edit)
		}
	}
	return edits
}

func scanUntil(scanner *bufio.Scanner, stop string) (string, bool) {
	var text strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == stop {
			return text.String(), true
		}
		if text.Len() > 0 {
			_ = text.WriteByte('\n')
		}
		_, _ = text.WriteString(line)
	}
	return "", false
}

func Apply(s string, edits []Edit) string {
	for _, e := range edits {
		s = strings.ReplaceAll(s, e.Search, e.Replace)
	}
	return s
}
