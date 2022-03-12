package golang

func (g *Generator) generateAliases() error {
	w := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		w.Line("// %s", g.config.GeneratedComment)
		w.Line("")
	}
	w.Line("package %s", g.packageName)
	w.Line("")

	w.Line("type (")
	for _, alias := range g.aliases {
		if alias.Doc != nil {
			_ = g.writeComments(w, 1, alias.Name, alias.Doc.Description)
		}
		w.IndentLine(1, "%s %s", alias.Name, g.GoTypeName(&alias.Type))

		if alias != g.aliases[len(g.aliases)-1] {
			w.Line("")
		}
	}
	w.Line(")")
	w.Line("")

	return g.writeFile("alias.go", w.b)
}
