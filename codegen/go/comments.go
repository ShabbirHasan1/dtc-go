package golang

import (
	"fmt"
	"github.com/moontrade/dtc-go/codegen/schema"
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

func (g *Generator) writeClearComment(w *Writer, msg *Struct, name string) {
	if len(name) > 0 {
		w.Line("// %s", name)
	}
	//w.Line("// {")

	var (
		maxWidth     = 0
		maxTypeWidth = 0
	)
	for _, f := range msg.Fields {
		if len(f.Fields) > 0 {
			continue
		}
		if maxWidth < len(f.Name) {
			maxWidth = len(f.Name)
		}
		tn := g.GoTypeName(&f.Type)
		if f.Type.Kind == schema.KindStringFixed {
			tn = fmt.Sprintf("%s[%d]", tn, f.Type.Size)
		}
		if maxTypeWidth < len(tn) {
			maxTypeWidth = len(tn)
		}
	}

	maxWidth += 2
	maxTypeWidth += 2
	pad := func(s string, width int) string {
		for len(s) < width {
			s = s + " "
		}
		return s
	}
	for _, f := range msg.Fields {
		if len(f.Fields) > 0 {
			continue
		}

		name := pad(f.Name, maxWidth)
		tn := g.GoTypeName(&f.Type)
		if f.Type.Kind == schema.KindStringFixed {
			tn = fmt.Sprintf("%s[%d]", tn, f.Type.Size)
		}

		tn = pad(tn, maxTypeWidth)

		if f.Initial == nil {
			switch f.Type.Kind {
			case schema.KindStringFixed, schema.KindStringVLS:
				w.Line("//     %s%s= \"\"", name, tn)
			case schema.KindBool:
				w.Line("//     %s%s= false", name, tn)
			case schema.KindEnum:
				option := f.Type.Enum.OptionByValue(0)
				if option != nil {
					w.Line("//     %s%s= %s  (%d)", name, option.Name, tn, option.Value)
				} else {
					w.Line("//     %s%s= 0", name, tn)
				}
			default:
				w.Line("//     %s%s= 0", name, tn)
			}
			continue
		}

		switch f.Initial.Type {
		case schema.ValueTypeSizeof:
			w.Line("//     %s%s= %sSize  (%d)", name, tn, msg.Name, msg.Size)
		case schema.ValueTypeString:
			w.Line("//     %s%s= \"%s\"", name, tn, f.Initial.Str)
		case schema.ValueTypeBool:
			if f.Initial.Int == 0 {
				w.Line("//     %s%s= false", name, tn)
			} else {
				w.Line("//     %s%s= true", name, tn)
			}
		case schema.ValueTypeInt:
			if f.Type.Kind == schema.KindBool {
				if f.Initial.Int == 0 {
					w.Line("//     %s%s= false", name, tn)
				} else {
					w.Line("//     %s%s= true", name, tn)
				}
			} else {
				w.Line("//     %s%s= %d", name, tn, f.Initial.Int)
			}
		case schema.ValueTypeUint:
			if f.Type.Kind == schema.KindBool {
				if f.Initial.Uint == 0 {
					w.Line("//     %s%s= false", tn, name)
				} else {
					w.Line("//     %s%s= true", tn, name)
				}
			} else {
				w.Line("//     %s%s= %d", name, tn, f.Initial.Int)
			}
		case schema.ValueTypeFloat:
			w.Line("//     %s%s= %f", name, tn, f.Initial.Float)
		case schema.ValueTypeFloat32Max:
			w.Line("//     %s%s= math.MaxFloat32", name, tn)
		case schema.ValueTypeFloat64Max:
			w.Line("//     %s%s= math.MaxFloat64", name, tn)
		case schema.ValueTypeConst:
			w.Line("//     %s%s= %s  (%d)", name, tn, f.Initial.Const.Name, f.Initial.Int)
		case schema.ValueTypeEnumOption:
			w.Line("//     %s%s= %s  (%d)", name, tn, f.Initial.EnumOption.Name, f.Initial.Int)
		}
	}
	//w.Line("// }")
}
