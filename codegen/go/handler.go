package golang

func (g *Generator) generateHandler() error {
	w := &Writer{}
	if len(g.config.GeneratedComment) > 0 {
		w.Line("// %s", g.config.GeneratedComment)
		w.Line("")
	}
	w.Line("package %s", g.packageName)
	w.Line("")
	w.Line("import (")
	w.IndentLine(1, "\"github.com/moontrade/dtc-go/message\"")
	w.IndentLine(1, "\"github.com/moontrade/log\"")
	//w.IndentLine(1, "\"net\"")
	w.Line(")")
	w.Line("")

	w.Line("type Handler interface {")
	for _, msg := range g.messages {
		if msg.Extends != nil {
			continue
		}
		if msg.Doc() != nil {
			if err := g.writeComments(w, 0, "On"+msg.Name(), msg.Doc().Description); err != nil {
				return err
			}
		}
		w.IndentLine(1, "On%s(msg %s) error", msg.Name(), msg.Name())

		if msg.Extension != nil {
			w.Line("")
			doc := msg.Extension.Doc()
			if doc == nil {
				doc = msg.Doc()
			}
			if doc != nil {
				if err := g.writeComments(w, 0, "On"+msg.Extension.Name(), doc.Description); err != nil {
					return err
				}
			}
			w.IndentLine(1, "On%s(msg %s) error", msg.Extension.Name(), msg.Extension.Name())
		}
		w.Line("")
	}

	w.IndentLine(1, "OnUnknown(b []byte, t uint16) error")
	w.Line("")
	w.IndentLine(1, "OnError(b []byte, t uint16, err error)")
	w.Line("}")
	w.Line("")

	w.Line("type Stub struct {}")
	w.Line("")
	for _, msg := range g.messages {
		if msg.Extends != nil {
			continue
		}
		if msg.Doc() != nil {
			if err := g.writeComments(w, 0, "On"+msg.Name(), msg.Doc().Description); err != nil {
				return err
			}
		}
		w.Line("func (s *Stub) On%s(msg %s) error {", msg.Name(), msg.Name())
		w.IndentLine(1, "log.Info(\"On%s: %s\", message.JsonString(msg))", msg.Name(), "%s")
		w.IndentLine(1, "return nil")
		w.Line("}")

		if msg.Extension != nil {
			w.Line("")
			doc := msg.Extension.Doc()
			if doc == nil {
				doc = msg.Doc()
			}
			if doc != nil {
				if err := g.writeComments(w, 0, "On"+msg.Extension.Name(), doc.Description); err != nil {
					return err
				}
			}
			w.Line("func (s *Stub) On%s(msg %s) error {", msg.Extension.Name(), msg.Extension.Name())
			w.IndentLine(1, "log.Info(\"On%s: %s\", message.JsonString(msg))", msg.Extension.Name(), "%s")
			w.IndentLine(1, "return nil")
			w.Line("}")
		}
		w.Line("")
	}
	w.Line("func (s *Stub) OnUnknown(b []byte, t uint16) error {")
	w.IndentLine(1, "log.Info(\"OnUnknown: size:%s type:%s\", len(b), t)", "%s", "%d")
	w.IndentLine(1, "return nil")
	w.Line("}")
	w.Line("")

	w.Line("func (s *Stub) OnError(b []byte, t uint16, err error) {")
	w.IndentLine(1, "log.Info(\"OnError: size:%s type:%s error:%s\", len(b), t, err.Error())", "%s", "%d", "%s")
	w.Line("}")
	w.Line("")

	//w.Line("type Sender struct {")
	//w.IndentLine(1, "c net.TCPConn")
	//w.Line("}")
	//w.Line("")
	//w.Line("func NewSender(c net.TCPConn) Sender {")
	//w.IndentLine(1, "return Sender{c}")
	//w.Line("}")
	//
	//for _, msg := range g.messages {
	//	s := msg.VLS
	//	if s == nil {
	//		s = msg.Fixed
	//	}
	//	if s.Doc != nil {
	//		if err := g.writeComments(w, 0, s.Name, s.Doc.Description); err != nil {
	//			return err
	//		}
	//	}
	//	w.Line("func (s *Sender) %s(msg %s) error {", s.Name, s.Name)
	//	w.IndentLine(1, "return msg.WriteToConn(s.c)")
	//	w.Line("}")
	//	w.Line("")
	//}

	w.Line("func Handle(b []byte, h Handler) error {")
	w.IndentLine(1, "if len(b) < 6 {")
	w.IndentLine(2, "return message.ErrShortBuffer")
	w.IndentLine(1, "}")
	w.IndentLine(1, "p := message.PointerOf(b)")
	w.IndentLine(1, "if int(p.AsUint16LE()) != len(b) {")
	w.IndentLine(2, "return message.ErrFraming")
	w.IndentLine(1, "}")
	w.IndentLine(1, "t := p.Uint16LE(2)")
	w.IndentLine(1, "switch t {")
	for _, msg := range g.messages {
		if msg.Extends != nil {
			continue
		}
		s := msg.VLS
		if s == nil {
			s = msg.Fixed
		}
		w.IndentLine(1, "case %d:", s.Type)

		if msg.Extension != nil {
			extension := msg.Extension.VLS
			if extension == nil {
				extension = msg.Extension.Fixed
			}
			w.IndentLine(2, "if len(b) > %d {", s.Size)
			w.IndentLine(3, "m, e := Parse%s(b)", extension.Name)
			w.IndentLine(3, "if e != nil {")
			w.IndentLine(4, "h.OnError(b, t, e)")
			w.IndentLine(4, "return e")
			w.IndentLine(3, "}")
			w.IndentLine(3, "return h.On%s(m)", extension.Name)
			w.IndentLine(2, "}")
		}
		w.IndentLine(2, "m, e := Parse%s(b)", s.Name)
		w.IndentLine(2, "if e != nil {")
		w.IndentLine(3, "h.OnError(b, t, e)")
		w.IndentLine(3, "return e")
		w.IndentLine(2, "}")
		w.IndentLine(2, "return h.On%s(m)", s.Name)
	}
	w.IndentLine(1, "default:")
	w.IndentLine(2, "return h.OnUnknown(b, t)")
	w.IndentLine(1, "}")
	w.Line("}")
	w.Line("")

	//w.Line("func FixedHandler(b []byte, h Handler) error {")
	//w.IndentLine(1, "if len(b) < 4 {")
	//w.IndentLine(2, "return message.ErrShortBuffer")
	//w.IndentLine(1, "}")
	//w.IndentLine(1, "p := message.PointerOf(b)")
	//w.IndentLine(1, "if int(p.AsUint16LE()) != len(b) {")
	//w.IndentLine(2, "return message.ErrFraming")
	//w.IndentLine(1, "}")
	//w.IndentLine(1, "t := p.Uint16LE(2)")
	//w.IndentLine(1, "switch t {")
	//for _, msg := range g.messages {
	//	s := msg.Fixed
	//	if s == nil {
	//		s = msg.VLS
	//	}
	//	w.IndentLine(1, "case %d:", s.Type)
	//	if s.VLS {
	//		w.IndentLine(2, "m, e := Parse%s(b)", s.Name)
	//		w.IndentLine(2, "if e != nil {")
	//		w.IndentLine(3, "return e")
	//		w.IndentLine(2, "}")
	//		w.IndentLine(2, "return h.On%s(m)")
	//	} else {
	//		w.IndentLine(2, "m, e := Parse%s(b)", s.Name)
	//		w.IndentLine(2, "if e != nil {")
	//		w.IndentLine(3, "return e")
	//		w.IndentLine(2, "}")
	//		w.IndentLine(2, "v := New%s()")
	//		w.IndentLine(2, "m.CopyTo(v)")
	//		w.IndentLine(2, "return h.On%s(v)")
	//	}
	//}
	//w.IndentLine(1, "}")
	//w.Line("}")
	//w.Line("")

	return g.writeFile("handler.go", w.b)
}
