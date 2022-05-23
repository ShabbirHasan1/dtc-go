package rust

func (g *Generator) generateEnums() error {
	w := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		w.Line("/// %s", g.config.GeneratedComment)
		w.Line("")
	}
	// w.Line("package %s", g.packageName)
	w.Line("")

	for _, enum := range g.enums {
		if enum.Doc != nil {
			_ = g.writeComments(w, 0, enum.Name, enum.Doc.Description)
		}
		w.Line("#[derive(Clone, Copy)]")
		w.Line("#[repr(%s)]", primitiveKindName(enum.Type))
		w.Line("pub enum %s {", enum.Name)
		for _, option := range enum.Options {
			if option.Doc != nil {
				_ = g.writeComments(w, 1, option.Name, option.Doc.Description)
			}
			w.IndentLine(1, "%s = %d,", option.Name, option.Value)
		}
		w.Line("}")
		w.Line("")

		/*
					#[inline]
			    pub fn to_le(self) -> Self {
			        unsafe { core::mem::transmute((self as u32).to_le()) }
			    }

			    #[inline]
			    pub fn from_le(value: Self) -> Self {
			        unsafe { core::mem::transmute(u32::from_le(value as u32)) }
			    }
		*/
		w.Line("impl %s {", enum.Name)
		w.IndentLine(1, "#[inline]")
		w.IndentLine(1, "pub fn to_le(self) -> Self {")
		w.IndentLine(2, "#[cfg(target_endian = \"little\")]")
		w.IndentLine(2, "{")
		w.IndentLine(3, "self")
		w.IndentLine(2, "}")
		w.IndentLine(2, "#[cfg(not(target_endian = \"little\"))]")
		w.IndentLine(2, "{")
		w.IndentLine(3, "unsafe { core::mem::transmute((self as %s).to_le()) }", primitiveTypeName(enum.Type))
		w.IndentLine(2, "}")
		w.IndentLine(1, "}")
		w.Line("")
		w.IndentLine(1, "#[inline]")
		w.IndentLine(1, "pub fn from_le(value: Self) -> Self {")
		w.IndentLine(2, "#[cfg(target_endian = \"little\")]")
		w.IndentLine(2, "{")
		w.IndentLine(3, "value")
		w.IndentLine(2, "}")
		w.IndentLine(2, "#[cfg(not(target_endian = \"little\"))]")
		w.IndentLine(2, "{")
		w.IndentLine(3, "unsafe { core::mem::transmute(%s::from_le(value as %s)) }",
			primitiveTypeName(enum.Type), primitiveTypeName(enum.Type))
		w.IndentLine(2, "}")

		w.IndentLine(1, "}")
		w.Line("}")
		w.Line("")

	}

	return g.writeFile("enums.rs", w.b)
}
