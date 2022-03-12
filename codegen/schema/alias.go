package schema

type Alias struct {
	Namespace *Namespace
	Name      string
	Type      Type
	Doc       *TypeDoc
}
