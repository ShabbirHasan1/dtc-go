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

	writer.Line("pub use crate::message::*;")
	writer.Line("pub use enums::*;")
	writer.Line("pub use alias::*;")
	writer.Line("pub use constants::*;")
	writer.Line("pub use factory::*;")
	writer.Line("pub use handler::*;")
	for _, m := range g.messages {
		writer.Line("pub use %s::*;", strings.ToLower(toSnakeCase(m.Name())))
	}
	writer.Line("")

	writer.Line("#[cfg(test)]")
	writer.Line("mod tests {")
	writer.IndentLine(1, "use super::*;")
	writer.Line("")
	writer.IndentLine(1, "#[test]")
	writer.IndentLine(1, "fn layout() {")
	for _, m := range g.messages {
		writer.IndentLine(2, "%s::tests::layout();", strings.ToLower(toSnakeCase(m.Name())))
		// write := func(s *Struct) {
		// 	writer.IndentLine(2, "assert_eq!(")
		// 	writer.IndentLine(3, "%dusize,", s.Size)
		// 	writer.IndentLine(3, "core::mem::size_of::<%s>(),", s.DataName)
		// 	writer.IndentLine(3, "\"%s sizeof expected {:} but was {:}\",", s.DataName)
		// 	writer.IndentLine(3, "%dusize,", s.Size)
		// 	writer.IndentLine(3, "core::mem::size_of::<%s>()", s.DataName)
		// 	writer.IndentLine(2, ");")
		// 	writer.IndentLine(2, "assert_eq!(")
		// 	writer.IndentLine(3, "%du16,", s.Size)
		// 	writer.IndentLine(3, "%s::new().size(),", s.Name)
		// 	writer.IndentLine(3, "\"%s sizeof expected {:} but was {:}\",", s.Name)
		// 	writer.IndentLine(3, "%du16,", s.Size)
		// 	writer.IndentLine(3, "%s::new().size(),", s.Name)
		// 	writer.IndentLine(2, ");")
		// 	writer.IndentLine(2, "assert_eq!(")
		// 	writer.IndentLine(3, "%s,", s.TypeConst.Name)
		// 	writer.IndentLine(3, "%s::new().r#type(),", s.Name)
		// 	writer.IndentLine(3, "\"%s type expected {:} but was {:}\",", s.Name)
		// 	writer.IndentLine(3, "%s,", s.TypeConst.Name)
		// 	writer.IndentLine(3, "%s::new().r#type(),", s.Name)
		// 	writer.IndentLine(2, ");")
		// 	writer.IndentLine(2, "assert_eq!(")
		// 	writer.IndentLine(3, "%du16,", s.Type)
		// 	writer.IndentLine(3, "%s::new().r#type(),", s.Name)
		// 	writer.IndentLine(3, "\"%s type expected {:} but was {:}\",", s.Name)
		// 	writer.IndentLine(3, "%du16,", s.Type)
		// 	writer.IndentLine(3, "%s::new().r#type(),", s.Name)
		// 	writer.IndentLine(2, ");")
		// }

		// if m.Fixed != nil {
		// 	write(m.Fixed)
		// }
		// if m.VLS != nil {
		// 	write(m.VLS)
		// }
	}
	writer.IndentLine(1, "}")
	writer.Line("}")
	writer.Line("")

	return g.writeFile("mod.rs", writer.b)
}
