package schema

import (
	"errors"
	"fmt"
	"strings"
)

type Schema struct {
	Docs               *Docs
	Fixed              *Namespace
	VLS                *Namespace
	NonStandard        *Namespace
	Unknown            *Namespace
	Aliases            []*Alias
	AliasesByMame      map[string]*Alias
	Constants          []*Const
	ConstantsByName    map[string]*Const
	DuplicateConstants []*Const
	Enums              []*Enum
	EnumsByName        map[string]*Enum
	EnumOptionsByName  map[string]*EnumOption
	Messages           []*Message
	MessagesByName     map[string]*Message
	MessagesByType     map[uint16]*Message
	DuplicateEnums     []*Enum
}

func NewSchema() *Schema {
	return &Schema{
		AliasesByMame:     make(map[string]*Alias),
		ConstantsByName:   make(map[string]*Const),
		EnumsByName:       make(map[string]*Enum),
		EnumOptionsByName: make(map[string]*EnumOption),
		MessagesByName:    make(map[string]*Message),
		MessagesByType:    make(map[uint16]*Message),
	}
}

func (schema *Schema) IsNonStandard() bool {
	return schema.NonStandard != nil
}

func (schema *Schema) GetNamespaces() []*Namespace {
	n := make([]*Namespace, 0, 3)
	if schema.Fixed != nil {
		n = append(n, schema.Fixed)
	}
	if schema.VLS != nil {
		n = append(n, schema.VLS)
	}
	if schema.NonStandard != nil {
		n = append(n, schema.NonStandard)
	}
	return n
}

func (schema *Schema) GetNamespace(name string) *Namespace {
	if strings.HasPrefix(name, "n_") {
		name = name[2:]
	}
	switch strings.ToUpper(name) {
	case "DTC", "FIXED":
		if schema.Fixed == nil {
			schema.Fixed = NewNamespace(NamespaceKindFixed, schema)
		}
		return schema.Fixed
	case "DTC_VLS", "VLS":
		if schema.VLS == nil {
			schema.VLS = NewNamespace(NamespaceKindVLS, schema)
		}
		return schema.VLS
	case "DTCNONSTANDARD", "DTC_NONSTANDARD", "NONSTANDARD", "NS", "EXTEND", "EXTENSION":
		if schema.NonStandard == nil {
			schema.NonStandard = NewNamespace(NamespaceKindNonStandard, schema)
		}
		return schema.NonStandard
	default:
		if schema.Unknown == nil {
			schema.Unknown = NewNamespace(NamespaceKindUnknown, schema)
		}
		return schema.Unknown
	}
}

func (schema *Schema) Validate() error {
	var (
		constants = make(map[string]*Const)
		enums     = make(map[string]*Enum)
	)
	namespaces := schema.GetNamespaces()
	for _, namespace := range namespaces {
		for _, constant := range namespace.Constants {
			if constants[constant.Name] != nil {
				dup := constants[constant.Name]
				if dup != constant {
					return errors.New(fmt.Sprintln("duplicate constant named: "+constant.Name, " in namespace:", constant.Namespace.Name, " and namespace:", constants[constant.Name].Namespace.Name))
				}
			}
			constants[constant.Name] = constant
		}
		for _, enum := range namespace.Enums {
			if enums[enum.Name] != nil {
				dup := enums[enum.Name]
				if dup != enum {
					return errors.New(fmt.Sprintln("duplicate constant named: "+enum.Name, " in namespace:", enum.Namespace.String(), " and namespace:", enums[enum.Name].Namespace.String()))
				}
			}
			enums[enum.Name] = enum
		}

		for _, s := range namespace.Structs {
			if err := s.Validate(); err != nil {
				return err
			}
		}

		for _, s := range namespace.Structs {
			s.Layout()
		}
	}

	for _, msg := range schema.Messages {
		if msg.Fixed == nil || msg.VLS == nil {
			continue
		}

		existing := schema.MessagesByType[msg.Type()]
		if existing != nil {
			existing.Extension = msg
			msg.Extends = existing
		} else {
			schema.MessagesByType[msg.Type()] = msg
		}

		if len(msg.Fixed.Fields) < 3 {
			continue
		}

		if len(msg.Fixed.Fields) != len(msg.VLS.Fields)-1 {
			return fmt.Errorf("message '%s' has field count mismatch between Fixed and VLS structs: %d vs %d", msg.Fixed.Name, len(msg.Fixed.Fields), len(msg.VLS.Fields))
		}

		if msg.Fixed.VLS {
			return fmt.Errorf("message '%s' was flagged as VLS when expecting fixed", msg.Fixed.Name)
		}
		if !msg.VLS.VLS {
			return fmt.Errorf("message '%s' was not flagged as VLS when expecting VLS", msg.VLS.Name)
		}

		var (
			fixedIndex = 2
			vlsIndex   = 3
		)

		for fixedIndex < len(msg.Fixed.Fields) {
			fixedField := msg.Fixed.Fields[fixedIndex]
			vlsField := msg.VLS.Fields[vlsIndex]

			if fixedField.Type.Union != nil {
				fixedIndex++
				vlsIndex++
				continue
			}

			if fixedField.Name != vlsField.Name {
				return fmt.Errorf("message '%s' has field Name mismatch between Fixed and VLS at index %d and Fixed name '%s' vs VLS name '%s'", msg.Fixed.Name, fixedIndex, fixedField.Name, vlsField.Name)
			}
			if fixedField.Type.Kind != vlsField.Type.Kind {
				if fixedField.Type.Kind == KindStringFixed && vlsField.Type.Kind == KindStringVLS {

				} else {
					return fmt.Errorf("message '%s' has field type mismatch between Fixed and VLS at index %d named '%s' and Fixed type %d vs VLS type %d", msg.Fixed.Name, fixedIndex, fixedField.Name, fixedField.Type.Kind, vlsField.Type.Kind)
				}
			}

			fixedIndex++
			vlsIndex++
		}
	}

	return nil
}
