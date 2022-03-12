package golang

import (
	"github.com/moontrade/dtc-go/codegen/schema"
)

func (g *Generator) generateConstants() error {
	w := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		w.Line("// %s", g.config.GeneratedComment)
		w.Line("")
	}
	w.Line("package %s", g.packageName)
	w.Line("")

	importMath := false
	for _, constant := range g.constants {
		switch constant.Const.Value.Type {
		case schema.ValueTypeFloat32Max, schema.ValueTypeFloat64Max:
			importMath = true
		}
	}
	if importMath {
		w.Line("import \"math\"")
		w.Line("")
	}

	w.Line("const (")
	for _, constant := range g.constants {
		w.IndentLine(1, "%s %s = %s", constant.Name, g.GoTypeName(&constant.Type), valueToGo(constant.Const.Value))
	}
	w.Line(")")
	w.Line("")

	return g.writeFile("constants.go", w.b)
}
