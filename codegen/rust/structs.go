package rust

import (
	"fmt"
	"strings"

	"github.com/moontrade/dtc-go/codegen/schema"
)

func (g *Generator) generateStructDebug(msg *Struct) error {
	return nil
}

func (g *Generator) generateStructs() error {
	for _, m := range g.messages {
		writer := &Writer{}
		if len(g.config.GeneratedComment) > 0 {
			writer.Line("// %s", g.config.GeneratedComment)
		}

		writer.Line("use super::*;")
		writer.Line("")

		// writeImports := func(w *Writer) {
		// 	w.Line("package %s", g.packageName)
		// 	w.Line("")
		// 	w.Line("import (")
		// 	w.IndentLine(1, "\"github.com/moontrade/dtc-go/message\"")
		// 	w.Line(")")
		// 	w.Line("")
		// }
		// writeImports(writer)

		writeMessageDefault := func(w *Writer, msg *Struct) {
			if msg == nil {
				return
			}
			g.writeClearComment(w, msg, "")
			w.Write("pub(crate) const %s_DEFAULT: [u8; %d] = ", strings.ToUpper(toSnakeCase(msg.Name)), len(msg.Init))
			w.WriteByteSlice(msg.Init)
			w.Write(";\n\n")
		}
		writeMessageSize := func(w *Writer, msg *Struct) {
			if msg == nil {
				return
			}
			w.Line("pub(crate) const %s_SIZE: usize = %d;", strings.ToUpper(toSnakeCase(msg.Name)), msg.Size)
			w.Line("")
		}

		writeMessageSize(writer, m.VLS)
		writeMessageSize(writer, m.Fixed)
		writeMessageDefault(writer, m.VLS)
		writeMessageDefault(writer, m.Fixed)

		// Trait
		{
			write := func(msg *Struct) {
				if msg.Doc != nil {
					_ = g.writeComments(writer, 0, msg.Name, msg.Doc.Description)
				}
				writer.Line("pub trait %s: Message {", m.Name())
				writer.IndentLine(1, "type Safe: %s;", m.Name())
				writer.IndentLine(1, "type Unsafe: %s;", m.Name())
				writer.Line("")
				for _, field := range msg.Fields {
					if !field.IsSuper() {
						continue
					}
					if len(field.Fields) > 0 {
						for _, f := range field.Fields {
							if f.Doc != nil {
								g.writeComments(writer, 1, f.Name, f.Doc.Description)
							}
							writer.IndentLine(1, "fn %s(&self) -> %s;", f.Name, g.PublicType(&f.Type))
							writer.Line("")
						}
					} else {
						if field.Doc != nil {
							g.writeComments(writer, 1, field.Name, field.Doc.Description)
						}
						writer.IndentLine(1, "fn %s(&self) -> %s;", field.Name, g.PublicType(&field.Type))
						writer.Line("")
					}
				}
				for i, field := range msg.Fields {
					if !field.IsSuper() {
						continue
					}
					if len(field.Fields) > 0 {
						for _, f := range field.Fields {
							if f.Doc != nil {
								g.writeComments(writer, 1, f.Name, f.Doc.Description)
							}
							writer.IndentLine(1, "fn %s(&mut self, value: %s) -> &mut Self;", f.SetterName(), g.PublicType(&f.Type))
						}
					} else {
						if field.Doc != nil {
							g.writeComments(writer, 1, field.Name, field.Doc.Description)
						}
						writer.IndentLine(1, "fn %s(&mut self, value: %s) -> &mut Self;", field.SetterName(), g.PublicType(&field.Type))
					}
					if i < len(msg.Fields)-1 {
						writer.Line("")
					}
				}
				writer.Line("")
				writer.IndentLine(1, "fn clone_safe(&self) -> Self::Safe;")
				writer.Line("")
				writer.IndentLine(1, "fn to_safe(self) -> Self::Safe;")
				writer.Line("")

				writer.IndentLine(1, "fn copy_to(&self, to: &mut impl %s) {", cleanStructName(msg.Struct.Name))
				for _, field := range msg.Fields {
					if !field.IsSuper() {
						continue
					}
					if len(field.Fields) > 0 {
						largest := field.Fields[0]
						for _, f := range field.Fields {
							if f.Type.Size > largest.Type.Size {
								largest = f
							}
						}
						writer.IndentLine(2, "to.%s(self.%s());", largest.SetterName(), largest.Name)
					} else {
						writer.IndentLine(2, "to.%s(self.%s());", field.SetterName(), field.Name)
					}
				}
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			} else {
				write(m.Fixed)
			}
		}

		// Type declarations
		{
			write := func(msg *Struct) {
				if msg.Doc != nil {
					_ = g.writeComments(writer, 0, msg.Name, msg.Doc.Description)
				}
				writer.Line("pub struct %s {", msg.Name)
				if msg.VLS {
					writer.IndentLine(1, "data: *const %s,", msg.DataName)
					writer.IndentLine(1, "capacity: usize")
				} else {
					writer.IndentLine(1, "data: *const %s", msg.DataName)
				}
				writer.Line("}")
				writer.Line("")

				writer.Line("pub struct %s {", msg.UnsafeName)
				if msg.VLS {
					writer.IndentLine(1, "data: *const %s,", msg.DataName)
					writer.IndentLine(1, "capacity: usize")
				} else {
					writer.IndentLine(1, "data: *const %s", msg.DataName)
				}
				writer.Line("}")
				writer.Line("")

				writer.Line("#[repr(packed(%d), C)]", msg.Struct.MaxAlign)
				writer.Line("pub struct %s {", msg.DataName)
				for _, field := range msg.Fields {
					if len(field.Fields) > 0 {
						largest := field.Fields[0]
						for _, f := range field.Fields {
							if f.Type.Size > largest.Type.Size {
								largest = f
							}
						}
						writer.IndentLine(1, "%s: %s,", largest.Name, g.RustTypeName(&largest.Type))
					} else {
						writer.IndentLine(1, "%s: %s,", field.Name, g.RustTypeName(&field.Type))
					}
				}
				writer.Line("}")
				writer.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Data Constructors
		{
			write := func(msg *Struct) {
				writer.Line("impl %s {", msg.DataName)
				writer.IndentLine(1, "pub fn new() -> Self {")
				writer.IndentLine(2, "Self {")
				for _, field := range msg.Fields {
					if len(field.Fields) > 0 {
						largest := field.Fields[0]
						for _, f := range field.Fields {
							if f.Type.Size > largest.Type.Size {
								largest = f
							}
						}
						field = largest
					}
					writer.IndentLine(3, "%s: %s,", field.Name, initValue(field.Field))
				}
				writer.IndentLine(2, "}")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			if m.Fixed != nil {
				write(m.Fixed)
			}
		}

		// Send impl
		{
			write := func(msg *Struct) {
				writer.Line("unsafe impl Send for %s {}", msg.Name)
				writer.Line("unsafe impl Send for %s {}", msg.UnsafeName)
				writer.Line("unsafe impl Send for %s {}", msg.DataName)
			}

			if m.Fixed != nil {
				write(m.Fixed)
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			writer.Line("")
		}

		// Drop
		{
			write := func(msg *Struct) {
				writer.Line("impl Drop for %s {", msg.Name)
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn drop(&mut self) {")
				writer.IndentLine(2, "crate::deallocate(self.data as *mut u8, self.capacity() as usize);")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
				writer.Line("impl Drop for %s {", msg.UnsafeName)
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn drop(&mut self) {")
				writer.IndentLine(2, "crate::deallocate(self.data as *mut u8, self.capacity() as usize);")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
			}

			if m.Fixed != nil {
				write(m.Fixed)
			}
			if m.VLS != nil {
				write(m.VLS)
			}
		}

		// Clone
		{
			write := func(msg *Struct) {
				writer.Line("impl Clone for %s {", msg.Name)
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn clone(&self) -> Self {")
				writer.IndentLine(2, "let mut c = Self::new();")
				writer.IndentLine(2, "self.copy_to(&mut c);")
				writer.IndentLine(2, "c")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
				writer.Line("impl Clone for %s {", msg.UnsafeName)
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn clone(&self) -> Self {")
				writer.IndentLine(2, "let mut c = Self::new();")
				writer.IndentLine(2, "self.copy_to(&mut c);")
				writer.IndentLine(2, "c")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
			}

			if m.Fixed != nil {
				write(m.Fixed)
			}
			if m.VLS != nil {
				write(m.VLS)
			}
		}

		// Into<Vec<u8>>
		{
			write := func(msg *Struct) {
				writer.Line("impl Into<Vec<u8>> for %s {", msg.Name)
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn into(self) -> Vec<u8> {")
				writer.IndentLine(2, "self.into_vec()")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
				writer.Line("impl Into<Vec<u8>> for %s {", msg.UnsafeName)
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn into(self) -> Vec<u8> {")
				writer.IndentLine(2, "self.into_vec()")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
			}

			if m.Fixed != nil {
				write(m.Fixed)
			}
			if m.VLS != nil {
				write(m.VLS)
			}
		}

		// Deref
		{
			write := func(msg *Struct) {
				writer.Line("impl core::ops::Deref for %s {", msg.Name)
				writer.IndentLine(1, "type Target = %s;", msg.DataName)
				writer.Line("")
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn deref(&self) -> &Self::Target {")
				writer.IndentLine(2, "unsafe { &*self.data }")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
				writer.Line("impl core::ops::DerefMut for %s {", msg.Name)
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn deref_mut(&mut self) -> &mut Self::Target {")
				writer.IndentLine(2, "unsafe { &mut *(self.data as *mut Self::Target) }")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
				writer.Line("impl core::ops::Deref for %s {", msg.UnsafeName)
				writer.IndentLine(1, "type Target = %s;", msg.DataName)
				writer.Line("")
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn deref(&self) -> &Self::Target {")
				writer.IndentLine(2, "unsafe { &*self.data }")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
				writer.Line("impl core::ops::DerefMut for %s {", msg.UnsafeName)
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn deref_mut(&mut self) -> &mut Self::Target {")
				writer.IndentLine(2, "unsafe { &mut *(self.data as *mut Self::Target) }")
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
			}

			if m.Fixed != nil {
				write(m.Fixed)
			}
			if m.VLS != nil {
				write(m.VLS)
			}
		}

		// impl Message
		{
			write := func(msg *Struct, isUnsafe bool) {
				if isUnsafe {
					writer.Line("impl crate::Message for %s {", msg.UnsafeName)
				} else {
					writer.Line("impl crate::Message for %s {", msg.Name)
				}
				writer.IndentLine(1, "type Data = %s;", msg.DataName)
				writer.Line("")
				writer.IndentLine(1, "const BASE_SIZE: usize = %d;", msg.Struct.Size)
				if msg.VLS {
					writer.IndentLine(1, "const BASE_SIZE_OFFSET: isize = 4;")
					writer.IndentLine(1, "const DEFAULT_CAPACITY: usize = Self::BASE_SIZE * 2;")
				} else {
					writer.IndentLine(1, "const BASE_SIZE_OFFSET: isize = 0;")
				}
				writer.Line("")
				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn new() -> Self {")
				writer.IndentLine(2, "Self {")
				if msg.VLS {
					writer.IndentLine(3, "data: crate::allocate(Self::BASE_SIZE, %s::new()),", msg.DataName)
					writer.IndentLine(3, "capacity: Self::DEFAULT_CAPACITY")
				} else {
					writer.IndentLine(3, "data: crate::allocate(Self::BASE_SIZE, %s::new())", msg.DataName)
				}
				writer.IndentLine(2, "}")
				writer.IndentLine(1, "}")
				writer.Line("")

				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn size(&self) -> u16 {")
				writer.IndentLine(2, "u16::from_le(self.size)")
				writer.IndentLine(1, "}")
				writer.Line("")

				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn r#type(&self) -> u16 {")
				writer.IndentLine(2, "u16::from_le(self.r#type)")
				writer.IndentLine(1, "}")
				writer.Line("")

				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn base_size(&self) -> u16 {")
				if msg.VLS {
					writer.IndentLine(2, "u16::from_le(self.base_size)")
				} else {
					writer.IndentLine(2, "u16::from_le(self.size)")
				}
				writer.IndentLine(1, "}")
				writer.Line("")

				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn capacity(&self) -> u16 {")
				if msg.VLS {
					writer.IndentLine(2, "self.capacity as u16")
				} else {
					writer.IndentLine(2, "u16::from_le(self.size)")
				}
				writer.IndentLine(1, "}")
				writer.Line("")

				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "unsafe fn as_ptr(&self) -> *const Self::Data {")
				writer.IndentLine(2, "self.data")
				writer.IndentLine(1, "}")
				writer.Line("")

				writer.IndentLine(1, "#[inline]")
				if msg.VLS {
					writer.IndentLine(1, "unsafe fn from_raw_parts(data: *const u8, capacity: usize) -> Self {")
				} else {
					writer.IndentLine(1, "unsafe fn from_raw_parts(data: *const u8, _: usize) -> Self {")
				}
				writer.IndentLine(2, "Self {")
				if msg.VLS {
					writer.IndentLine(3, "data: data as *const %s,", msg.DataName)
					writer.IndentLine(3, "capacity")
				} else {
					writer.IndentLine(3, "data: data as *const %s", msg.DataName)
				}
				writer.IndentLine(2, "}")
				writer.IndentLine(1, "}")
				writer.Line("")

				writer.Line("}")

				if msg.VLS {
					writer.Line("")
					if isUnsafe {
						writer.Line("impl crate::VLSMessage for %s {", msg.UnsafeName)
					} else {
						writer.Line("impl crate::VLSMessage for %s {", msg.Name)
					}

					writer.IndentLine(1, "#[inline]")
					writer.IndentLine(1, "unsafe fn set_ptr(&mut self, value: *const u8) {")
					writer.IndentLine(2, "self.data = value as *const %s;", msg.DataName)
					writer.IndentLine(1, "}")
					writer.Line("")

					writer.IndentLine(1, "#[inline]")
					writer.IndentLine(1, "fn set_capacity(&mut self, capacity: u16) {")
					writer.IndentLine(2, "self.capacity = capacity as usize;")
					writer.IndentLine(1, "}")
					writer.Line("")

					writer.IndentLine(1, "#[inline]")
					writer.IndentLine(1, "fn set_size(&mut self, size: u16) {")
					writer.IndentLine(2, "self.size = size.to_le();")
					writer.IndentLine(1, "}")
					writer.Line("")

					writer.Write("}")
				}
			}

			if m.Fixed != nil {
				write(m.Fixed, false)
				write(m.Fixed, true)
			}
			if m.VLS != nil {
				write(m.VLS, false)
				write(m.VLS, true)
			}
		}

		// impl
		{
			write := func(msg *Struct, isUnsafe bool) {
				if msg.Doc != nil {
					_ = g.writeComments(writer, 0, msg.Name, msg.Doc.Description)
				}
				if isUnsafe {
					writer.Line("impl %s for %s {", cleanStructName(msg.Struct.Name), msg.UnsafeName)
				} else {
					writer.Line("impl %s for %s {", cleanStructName(msg.Struct.Name), msg.Name)
				}

				writer.IndentLine(1, "type Safe = %s;", msg.Name)
				writer.IndentLine(1, "type Unsafe = %s;", msg.UnsafeName)
				writer.Line("")

				for _, field := range msg.Fields {
					if !field.IsSuper() {
						continue
					}
					if len(field.Fields) > 0 {
						largest := field.Fields[0]
						for _, f := range field.Fields {
							if f.Type.Size > largest.Type.Size {
								largest = f
							}
						}
						for _, f := range field.Fields {
							if f.Doc != nil {
								g.writeComments(writer, 1, f.Name, f.Doc.Description)
							}
							writer.IndentLine(1, "fn %s(&self) -> %s {", f.Name, g.PublicType(&f.Type))
							g.getAccessor(writer, 2, largest, isUnsafe)
							writer.IndentLine(1, "}")
							writer.Line("")
						}
					} else {
						if field.Doc != nil {
							g.writeComments(writer, 1, field.Name, field.Doc.Description)
						}
						writer.IndentLine(1, "fn %s(&self) -> %s {", field.Name, g.PublicType(&field.Type))
						g.getAccessor(writer, 2, field, isUnsafe)
						writer.IndentLine(1, "}")
						writer.Line("")
					}
				}
				for _, field := range msg.Fields {
					if !field.IsSuper() {
						continue
					}
					if len(field.Fields) > 0 {
						largest := field.Fields[0]
						for _, f := range field.Fields {
							if f.Type.Size > largest.Type.Size {
								largest = f
							}
						}
						for _, f := range field.Fields {
							if f.Doc != nil {
								g.writeComments(writer, 1, f.Name, f.Doc.Description)
							}
							writer.IndentLine(1, "fn %s(&mut self, value: %s) -> &mut Self {", f.SetterName(), g.PublicType(&f.Type))
							g.setAccessor(writer, 2, largest, isUnsafe)
							writer.IndentLine(1, "}")
							writer.Line("")
						}
					} else {
						if field.Doc != nil {
							g.writeComments(writer, 1, field.Name, field.Doc.Description)
						}
						writer.IndentLine(1, "fn %s(&mut self, value: %s) -> &mut Self {", field.SetterName(), g.PublicType(&field.Type))
						g.setAccessor(writer, 2, field, isUnsafe)
						writer.IndentLine(1, "}")
						writer.Line("")
					}
					writer.Line("")
				}

				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn clone_safe(&self) -> Self::Safe {")
				writer.IndentLine(2, "let mut s = Self::Safe::new();")
				writer.IndentLine(2, "self.copy_to(&mut s);")
				writer.IndentLine(2, "s")
				writer.IndentLine(1, "}")
				writer.Line("")

				writer.IndentLine(1, "#[inline]")
				writer.IndentLine(1, "fn to_safe(self) -> Self::Safe {")
				if isUnsafe {
					writer.IndentLine(2, "let mut s = Self::Safe::new();")
					writer.IndentLine(2, "self.copy_to(&mut s);")
					writer.IndentLine(2, "s")
				} else {
					writer.IndentLine(2, "self")
				}
				writer.IndentLine(1, "}")
				writer.Line("}")
				writer.Line("")
			}
			if m.VLS != nil {
				write(m.VLS, false)
				write(m.VLS, true)
			}
			if m.Fixed != nil {
				write(m.Fixed, false)
				write(m.Fixed, true)
			}
		}

		{
			writer.Line("#[cfg(test)]")
			writer.Line("pub(crate) mod tests {")
			writer.IndentLine(1, "use super::*;")
			writer.Line("")
			writer.IndentLine(1, "#[test]")
			writer.IndentLine(1, "pub(crate) fn layout() {")
			write := func(s *Struct) {
				writer.IndentLine(2, "unsafe {")
				writer.IndentLine(3, "assert_eq!(")
				writer.IndentLine(4, "%dusize,", s.Size)
				writer.IndentLine(4, "core::mem::size_of::<%s>(),", s.DataName)
				writer.IndentLine(4, "\"%s sizeof expected {:} but was {:}\",", s.DataName)
				writer.IndentLine(4, "%dusize,", s.Size)
				writer.IndentLine(4, "core::mem::size_of::<%s>()", s.DataName)
				writer.IndentLine(3, ");")
				writer.IndentLine(3, "assert_eq!(")
				writer.IndentLine(4, "%du16,", s.Size)
				writer.IndentLine(4, "%s::new().size(),", s.Name)
				writer.IndentLine(4, "\"%s sizeof expected {:} but was {:}\",", s.Name)
				writer.IndentLine(4, "%du16,", s.Size)
				writer.IndentLine(4, "%s::new().size(),", s.Name)
				writer.IndentLine(3, ");")
				writer.IndentLine(3, "assert_eq!(")
				writer.IndentLine(4, "%s,", s.TypeConst.Name)
				writer.IndentLine(4, "%s::new().r#type(),", s.Name)
				writer.IndentLine(4, "\"%s type expected {:} but was {:}\",", s.Name)
				writer.IndentLine(4, "%s,", s.TypeConst.Name)
				writer.IndentLine(4, "%s::new().r#type(),", s.Name)
				writer.IndentLine(3, ");")
				writer.IndentLine(3, "assert_eq!(")
				writer.IndentLine(4, "%du16,", s.Type)
				writer.IndentLine(4, "%s::new().r#type(),", s.Name)
				writer.IndentLine(4, "\"%s type expected {:} but was {:}\",", s.Name)
				writer.IndentLine(4, "%du16,", s.Type)
				writer.IndentLine(4, "%s::new().r#type(),", s.Name)
				writer.IndentLine(3, ");")

				writer.IndentLine(3, "let d = %s::new();", s.DataName)
				writer.IndentLine(3, "let p = (&d as *const _ as *const u8).offset(0) as usize;")
				for _, field := range s.Fields {
					if len(field.Fields) > 0 {
						largest := field.Fields[0]
						for _, f := range field.Fields {
							if f.Type.Size > largest.Type.Size {
								largest = f
							}
						}
						field = largest
					}
					name := field.Name
					if name == "r#type" {
						name = "type"
					}
					writer.IndentLine(3, "assert_eq!(")
					writer.IndentLine(4, "%dusize,", field.Type.Offset)
					writer.IndentLine(4, "(core::ptr::addr_of!(d.%s) as usize) - p,", field.Name)
					writer.IndentLine(4, "\"%s offset expected {:} but was {:}\",", name)
					writer.IndentLine(4, "%dusize,", field.Type.Offset)
					writer.IndentLine(4, "(core::ptr::addr_of!(d.%s) as usize) - p,", field.Name)
					writer.IndentLine(3, ");")
				}
				writer.IndentLine(2, "}")
			}

			if m.Fixed != nil {
				write(m.Fixed)
			}
			if m.VLS != nil {
				write(m.VLS)
			}
			writer.IndentLine(1, "}")
			writer.Line("}")
			writer.Line("")
		}

		if err := g.writeFile(fmt.Sprintf("%s.rs", toSnakeCase(m.Name())), writer.b); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) getAccessor(w *Writer, indent int, f *Field, isUnsafe bool) {
	if isUnsafe {
		w.IndentLine(indent, "if self.is_out_of_bounds(%d) {", f.Field.Type.Offset+f.Field.Type.Size)
		if f.Type.Kind == schema.KindStringFixed || f.Type.Kind == schema.KindStringVLS {
			w.IndentLine(indent+1, "\"\"")
		} else {
			w.IndentLine(indent+1, initValue(f.Field))
		}
		w.IndentLine(indent, "} else {")
		indent++
	}
	t := &f.Type
	if t.Kind == schema.KindAlias {
		t = &t.Alias.Type
	}
	switch t.Kind {
	case schema.KindBool:
		w.IndentLine(indent, "self.%s", f.Name)
	case schema.KindInt8:
		w.IndentLine(indent, "self.%s", f.Name)
	case schema.KindInt16:
		w.IndentLine(indent, "i16::from_le(self.%s)", f.Name)
	case schema.KindInt32:
		w.IndentLine(indent, "i32::from_le(self.%s)", f.Name)
	case schema.KindInt64:
		w.IndentLine(indent, "i64::from_le(self.%s)", f.Name)
	case schema.KindUint8:
		w.IndentLine(indent, "self.%s", f.Name)
	case schema.KindUint16:
		w.IndentLine(indent, "u16::from_le(self.%s)", f.Name)
	case schema.KindUint32:
		w.IndentLine(indent, "u32::from_le(self.%s)", f.Name)
	case schema.KindUint64:
		w.IndentLine(indent, "u64::from_le(self.%s)", f.Name)
	case schema.KindFloat32:
		w.IndentLine(indent, "f32_le(self.%s)", f.Name)
	case schema.KindFloat64:
		w.IndentLine(indent, "f64_le(self.%s)", f.Name)
	case schema.KindStringFixed:
		w.IndentLine(indent, "get_fixed(&self.%s[..])", f.Name)
	case schema.KindStringVLS:
		w.IndentLine(indent, "get_vls(self, self.%s)", f.Name)
	case schema.KindEnum:
		w.IndentLine(indent, "%s::from_le(self.%s)",
			cleanStructName(f.Type.Enum.Name), f.Name)
	}
	if isUnsafe {
		indent--
		w.IndentLine(indent, "}")
	}
}

func (g *Generator) setAccessor(w *Writer, indent int, f *Field, isUnsafe bool) {
	if isUnsafe {
		w.IndentLine(indent, "if !self.is_out_of_bounds(%d) {", f.Field.Type.Offset+f.Field.Type.Size)
		indent++
	}
	t := &f.Type
	if t.Kind == schema.KindAlias {
		t = &t.Alias.Type
	}
	switch t.Kind {
	case schema.KindBool:
		w.IndentLine(indent, "self.%s = value;", f.Name)
	case schema.KindInt8:
		w.IndentLine(indent, "self.%s = value;", f.Name)
	case schema.KindInt16:
		w.IndentLine(indent, "self.%s = value.to_le();", f.Name)
	case schema.KindInt32:
		w.IndentLine(indent, "self.%s = value.to_le();", f.Name)
	case schema.KindInt64:
		w.IndentLine(indent, "self.%s = value.to_le();", f.Name)
	case schema.KindUint8:
		w.IndentLine(indent, "self.%s = value;", f.Name)
	case schema.KindUint16:
		w.IndentLine(indent, "self.%s = value.to_le();", f.Name)
	case schema.KindUint32:
		w.IndentLine(indent, "self.%s = value.to_le();", f.Name)
	case schema.KindUint64:
		w.IndentLine(indent, "self.%s = value.to_le();", f.Name)
	case schema.KindFloat32:
		w.IndentLine(indent, "self.%s = f32_le(value);", f.Name)
	case schema.KindFloat64:
		w.IndentLine(indent, "self.%s = f64_le(value);", f.Name)
	case schema.KindStringFixed:
		w.IndentLine(indent, "set_fixed(&mut self.%s[..], value);", f.Name)
	case schema.KindStringVLS:
		w.IndentLine(indent, "self.%s = set_vls(self, self.%s, value);", f.Name, f.Name)
	case schema.KindEnum:
		w.IndentLine(indent, "self.%s = unsafe { core::mem::transmute((value as %s).to_le()) };",
			f.Name, primitiveTypeName(f.Type.Enum.Type))
	}
	if isUnsafe {
		indent--
		w.IndentLine(indent, "}")
	}
	w.IndentLine(indent, "self")
}
