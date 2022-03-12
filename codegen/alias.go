package codegen

type Alias struct {
	Namespace *Namespace
	Name      string
	Type      Type
	Doc       *TypeDoc
}
