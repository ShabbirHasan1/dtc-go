package rust

func (g *Generator) generateAliases() error {
	w := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		w.Line("// %s", g.config.GeneratedComment)
		w.Line("")
	}

	for _, alias := range g.aliases {
		if alias.Doc != nil {
			_ = g.writeComments(w, 0, alias.Name, alias.Doc.Description)
		}
		w.Line("pub type %s = %s;", alias.Name, g.RustTypeName(&alias.Type))

		if alias != g.aliases[len(g.aliases)-1] {
			w.Line("")
		}
	}
	w.Line("")

	return g.writeFile("alias.rs", w.b)
}
