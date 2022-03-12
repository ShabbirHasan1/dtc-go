package schema

type Const struct {
	Namespace *Namespace
	Name      string
	Type      Type
	Length    int
	Value     *Value
	Comment   string
	Comments  []string
}
