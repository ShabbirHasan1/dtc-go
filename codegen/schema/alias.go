package schema

type Alias struct {
	Namespace NamespaceKind
	Name      string
	Type      Type
	Doc       *TypeDoc
}
