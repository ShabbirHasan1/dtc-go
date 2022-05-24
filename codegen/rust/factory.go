package rust

func (g *Generator) generateFactory() error {
	writer := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		writer.Line("// %s", g.config.GeneratedComment)
		writer.Line("")
	}

	writer.Line("use super::*;")
	writer.Line("use crate::Message;")

	writer.Line("")
	writer.Line("pub trait Factory: Send + 'static {")
	for _, m := range g.messages {
		name := m.Name()
		writer.IndentLine(1, "type %s: %s + Message + Send;", name, name)
		writer.IndentLine(1, "type %sUnsafe: %s + Message + Send;", name, name)
	}
	writer.Line("}")
	writer.Line("")

	writer.Line("pub struct FactoryFixed;")
	writer.Line("")
	writer.Line("impl Factory for FactoryFixed {")
	for _, m := range g.messages {
		s := m.Fixed
		if s == nil {
			s = m.VLS
		}
		name := m.Name()
		writer.IndentLine(1, "type %s = %s;", name, s.Name)
		writer.IndentLine(1, "type %sUnsafe = %s;", name, s.UnsafeName)
	}
	writer.Line("}")
	writer.Line("")

	writer.Line("pub struct FactoryVLS;")
	writer.Line("")
	writer.Line("impl Factory for FactoryVLS {")
	for _, m := range g.messages {
		s := m.VLS
		if s == nil {
			s = m.Fixed
		}
		name := m.Name()
		writer.IndentLine(1, "type %s = %s;", name, s.Name)
		writer.IndentLine(1, "type %sUnsafe = %s;", name, s.UnsafeName)
	}
	writer.Line("}")
	writer.Line("")

	return g.writeFile("factory.rs", writer.b)
}
