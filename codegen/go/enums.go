package golang

func (g *Generator) generateEnums() error {
	w := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		w.Line("// %s", g.config.GeneratedComment)
		w.Line("")
	}
	w.Line("package %s", g.packageName)
	w.Line("")

	for _, enum := range g.enums {
		if enum.Doc != nil {
			_ = g.writeComments(w, 0, enum.Name, enum.Doc.Description)
		}
		w.Line("type %s %s", enum.Name, primitiveKindName(enum.Type))
		w.Line("")
		w.Line("const (")
		for _, option := range enum.Options {
			if option.Doc != nil {
				_ = g.writeComments(w, 1, option.Name, option.Doc.Description)
			}
			w.IndentLine(1, "%s %s = %d", option.Name, enum.Name, option.Value)
		}
		w.Line(")")
		w.Line("")
	}

	return g.writeFile("enums.go", w.b)
}
