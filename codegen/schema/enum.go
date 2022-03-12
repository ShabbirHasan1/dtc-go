package schema

import (
	"fmt"
	"github.com/yoheimuta/go-protoparser/v4/parser"
	"strconv"
)

type Enum struct {
	Namespace     *Namespace
	Name          string
	Type          Kind
	Options       []*EnumOption
	OptionsByName map[string]*EnumOption
	Proto         *parser.Enum
	Doc           *TypeDoc
}

type EnumOption struct {
	Enum    *Enum
	Name    string
	Value   int64
	Comment string
	Proto   *parser.EnumField
	Doc     *EnumOptionDoc
}

func (enum *Enum) OptionByValue(value int64) *EnumOption {
	for _, o := range enum.Options {
		if o.Value == value {
			return o
		}
	}
	return nil
}

func (enum *Enum) Bind(proto *parser.Enum) error {
	for _, ef := range proto.EnumBody {
		switch ef := ef.(type) {
		case *parser.EnumField:
			value, err := strconv.ParseInt(ef.Number, 10, 64)
			if err != nil {
				return fmt.Errorf("proto: enum %s.%s has invalid number: %s", proto.EnumName, ef.Ident, err.Error())
			}
			opt := enum.OptionsByName[ef.Ident]
			if opt == nil {
				opt = enum.OptionByValue(value)
			}
			if opt == nil {
				if ef.Number == "0" {
					continue
				}
				return fmt.Errorf("proto: enum %s.%s does not match any options for %s", proto.EnumName, ef.Ident, enum.Name)
			}
			opt.Proto = ef
		}
	}
	enum.Proto = proto
	return nil
}
