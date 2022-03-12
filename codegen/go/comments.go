package golang

import (
	"fmt"
	"strings"
)

func (g *Generator) writeComments(w *Writer, indent int, name string, comments []string) error {
	if len(comments) == 0 {
		return nil
	}
	comments = g.normalizeComments(comments)
	if !strings.HasPrefix(comments[0], name) {
		comments[0] = fmt.Sprintf("%s %s", name, strings.TrimSpace(comments[0]))
	}
	for _, comment := range comments {
		comment = strings.TrimSpace(comment)
		if len(comment) == 0 {
			w.IndentLine(indent, "//")
		} else {
			w.IndentLine(indent, "// %s", comment)
		}
	}
	return nil
}

func (g *Generator) normalizeComments(comments []string) []string {
	to := make([]string, 0, len(comments))
	for _, comment := range comments {
		for _, alias := range g.aliases {
			comment = strings.ReplaceAll(comment, alias.Alias.Name, alias.Name)
		}
		for _, enum := range g.enums {
			comment = strings.ReplaceAll(comment, enum.Enum.Name, enum.Name)
		}
		for _, msg := range g.messages {
			m := msg.VLS
			if m == nil {
				m = msg.Fixed
			}
			comment = strings.ReplaceAll(comment, m.Struct.Name, m.Name)
		}

		for len(comment) > 70 {
			index := 69
		LOOP:
			for ; index < len(comment); index++ {
				c := comment[index]
				switch c {
				case ' ', '\r', '\t':
					break LOOP

				case '.':
					index += 1

					break LOOP
				}
			}

			if index >= len(comment) {
				comment = strings.TrimSpace(comment)
				if len(comment) > 0 {
					to = append(to, comment)
				}
				break
			}

			cutoff := strings.TrimSpace(comment[0:index])
			to = append(to, cutoff)
			comment = strings.TrimSpace(comment[index:])
		}

		to = append(to, comment)
	}
	return to
}
