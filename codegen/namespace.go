package codegen

type NamespaceKind byte

const (
	NamespaceKindUnknown     NamespaceKind = 0
	NamespaceKindFixed       NamespaceKind = 1
	NamespaceKindVLS         NamespaceKind = 2
	NamespaceKindNonStandard NamespaceKind = 3
)

func (nk NamespaceKind) String() string {
	switch nk {
	case NamespaceKindFixed:
		return "Fixed"
	case NamespaceKindVLS:
		return "VLS"
	case NamespaceKindNonStandard:
		return "NonStandard"
	default:
		return "Unknown"
	}
}

type Namespace struct {
	Header          *CHeaderFile
	Kind            NamespaceKind
	Schema          *Schema
	Name            string
	Alias           []*Alias
	AliasByName     map[string]*Alias
	Constants       []*Const
	ConstantsByName map[string]*Const
	Enums           []*Enum
	EnumsByName     map[string]*Enum
	EnumOptions     map[string]*EnumOption
	Structs         []*Struct
	StructsByName   map[string]*Struct
}

func NewNamespace(kind NamespaceKind, schema *Schema) *Namespace {
	return &Namespace{
		Kind:            kind,
		Schema:          schema,
		Name:            kind.String(),
		AliasByName:     make(map[string]*Alias),
		ConstantsByName: make(map[string]*Const),
		EnumsByName:     make(map[string]*Enum),
		EnumOptions:     make(map[string]*EnumOption),
		StructsByName:   make(map[string]*Struct),
	}
}
