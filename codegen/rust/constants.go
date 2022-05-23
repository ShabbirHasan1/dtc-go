package rust

func (g *Generator) generateConstants() error {
	w := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		w.Line("// %s", g.config.GeneratedComment)
		w.Line("")
	}

	for _, constant := range g.constants {
		w.Line("pub const %s: %s = %s;", constant.Name, g.RustTypeName(&constant.Type), valueToRust(constant.Const.Value))
	}
	w.Line("")

	return g.writeFile("constants.rs", w.b)
}
