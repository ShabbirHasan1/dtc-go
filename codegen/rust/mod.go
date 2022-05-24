package rust

import "strings"

func (g *Generator) generateMod() error {
	writer := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		writer.Line("// %s", g.config.GeneratedComment)
		writer.Line("")
	}

	writer.Line("mod enums;")
	writer.Line("mod alias;")
	writer.Line("mod constants;")
	writer.Line("mod factory;")
	writer.Line("mod handler;")
	for _, m := range g.messages {
		writer.Line("mod %s;", strings.ToLower(toSnakeCase(m.Name())))
	}
	writer.Line("")

	writer.Line("pub use enums::*;")
	writer.Line("pub use alias::*;")
	writer.Line("pub use constants::*;")
	writer.Line("pub use factory::*;")
	writer.Line("pub use handler::*;")
	for _, m := range g.messages {
		writer.Line("pub use %s::*;", strings.ToLower(toSnakeCase(m.Name())))
	}
	writer.Line("")

	return g.writeFile("mod.rs", writer.b)
}
