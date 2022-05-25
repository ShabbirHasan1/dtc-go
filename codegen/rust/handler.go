package rust

import "strings"

func (g *Generator) generateHandler() error {
	writer := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		writer.Line("// %s", g.config.GeneratedComment)
		writer.Line("")
	}

	writer.Line("use async_trait::async_trait;")
	writer.Line("use super::*;")
	writer.Line("use crate::{Connection, Error, Parsed, message::parse};")

	writer.Line("")
	writer.Line("#[async_trait]")
	writer.Line("pub trait Handler: Send + 'static {")

	for _, m := range g.messages {
		s := m.Fixed
		if s == nil {
			s = m.VLS
		}
		if m.IsExtension {
			continue
		}
		if s.Doc != nil {
			g.writeComments(writer, 1, s.Name, s.Doc.Description)
		}
		writer.IndentLine(1, "fn on_%s(&self, conn: &Connection, msg: &impl %s) -> Result<(), Error>;",
			strings.ToLower(toSnakeCase(m.Name())), m.Name())

		if m.Extension != nil {
			writer.Line("")

			if s.Doc != nil {
				g.writeComments(writer, 1, m.Extension.Name(), s.Doc.Description)
			}
			writer.IndentLine(1, "fn on_%s(&self, conn: &Connection, msg: &impl %s) -> Result<(), Error>;",
				strings.ToLower(toSnakeCase(m.Extension.Name())), m.Extension.Name())
		}

		writer.Line("")
	}

	writer.IndentLine(1, "fn on_unknown_message(&self, code: u16, data: &[u8]) -> Result<(), Error>;")

	writer.Line("}")
	writer.Line("")

	writer.Line("#[inline]")
	writer.Line("pub fn handle<F: Factory, H: Handler>(")
	writer.IndentLine(1, "handler: &H,")
	writer.IndentLine(1, "data: &[u8],")
	writer.IndentLine(1, "conn: &crate::Connection,")
	writer.Line(") -> Result<(), Error> {")
	writer.IndentLine(1, "if data.len() < 4 || (unsafe { u16::from_le(*(data[0..2].as_ptr() as *const u16)) } as usize) != data.len() {")
	writer.IndentLine(2, "return Err(Error::Malformed(\"\"));")
	writer.IndentLine(1, "}")
	writer.IndentLine(1, "let r#type = unsafe { u16::from_le(*(data[2..].as_ptr() as *const u16)) };")
	writer.IndentLine(1, "match r#type {")
	for _, m := range g.messages {
		s := m.Fixed
		if s == nil {
			s = m.VLS
		}
		if s.TypeConst == nil {
			println("nil")
		}
		if m.IsExtension {
			continue
		}
		if m.Extension != nil {
			writer.IndentLine(2, "%s => if data.len() > %d {", s.TypeConst.Name, s.Size)
			writer.IndentLine(3, "match parse::<F::%s, F::%sUnsafe>(data)? {", m.Extension.Name(), m.Extension.Name())
			writer.IndentLine(4, "Parsed::Left(m) => handler.on_%s(conn, m),", strings.ToLower(toSnakeCase(m.Extension.Name())))
			writer.IndentLine(4, "Parsed::Right(m) => handler.on_%s(conn, m),", strings.ToLower(toSnakeCase(m.Extension.Name())))
			writer.IndentLine(3, "}")
			writer.IndentLine(2, "} else {")
			writer.IndentLine(3, "match parse::<F::%s, F::%sUnsafe>(data)? {", m.Name(), m.Name())
			writer.IndentLine(4, "Parsed::Left(m) => handler.on_%s(conn, m),", strings.ToLower(toSnakeCase(m.Name())))
			writer.IndentLine(4, "Parsed::Right(m) => handler.on_%s(conn, m),", strings.ToLower(toSnakeCase(m.Name())))
			writer.IndentLine(3, "}")
			writer.IndentLine(2, "},")
		} else {
			writer.IndentLine(2, "%s => match parse::<F::%s, F::%sUnsafe>(data)? {", s.TypeConst.Name, m.Name(), m.Name())
			writer.IndentLine(3, "Parsed::Left(m) => handler.on_%s(conn, m),", strings.ToLower(toSnakeCase(m.Name())))
			writer.IndentLine(3, "Parsed::Right(m) => handler.on_%s(conn, m),", strings.ToLower(toSnakeCase(m.Name())))
			writer.IndentLine(2, "},")
		}
	}
	writer.IndentLine(2, "code => handler.on_unknown_message(code, data),")
	writer.IndentLine(1, "}")
	writer.Line("}")
	writer.Line("")

	return g.writeFile("handler.rs", writer.b)
}
