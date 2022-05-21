package rust

func (g *Generator) generateConstants() error {
	w := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		w.Line("/// %s", g.config.GeneratedComment)
		w.Line("")
	}
	// w.Line("package %s", g.packageName)
	w.Line("")

	// importMath := false
	// for _, constant := range g.constants {
	// 	switch constant.Const.Value.Type {
	// 	case schema.ValueTypeFloat32Max, schema.ValueTypeFloat64Max:
	// 		importMath = true
	// 	}
	// }
	// if importMath {
	// 	w.Line("import \"math\"")
	// 	w.Line("")
	// }

	for _, constant := range g.constants {
		w.Line("pub const %s: %s = %s;", constant.Name, g.RustTypeName(&constant.Type), valueToRust(constant.Const.Value))
	}
	w.Line("")

	return g.writeFile("constants.rs", w.b)
}
