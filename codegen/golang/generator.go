package golang

import "github.com/moontrade/dtc-go/codegen"

type Generator struct {
	namespaces *codegen.Namespaces
	packages   map[string]*Package
}

type Config struct {
	Dir                   string
	RootPackage           string
	DTCPackage            string
	DTCVLSPackage         string
	DTCNonStandardPackage string
}

func DefaultConfig(dir string) *Config {
	return &Config{
		Dir:                   dir,
		RootPackage:           "github.com/moontrade/dtc-go/model",
		DTCPackage:            "fixed",
		DTCVLSPackage:         "vls",
		DTCNonStandardPackage: "ns",
	}
}

func NewGenerator(config *Config, namespaces *codegen.Namespaces) *Generator {
	generator := &Generator{
		namespaces: namespaces,
		packages:   make(map[string]*Package),
	}

	for _, namespace := range namespaces.Namespaces {
		pkg := &Package{
			Name: namespace.Name,

			Namespace: nil,
		}
		_ = pkg
	}

	return generator
}
