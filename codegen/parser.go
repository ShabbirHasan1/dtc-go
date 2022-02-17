package codegen

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TypeKind byte

const (
	TypeUnknown     TypeKind = 0
	TypeInt8        TypeKind = 1
	TypeUInt8       TypeKind = 2
	TypeInt16       TypeKind = 3
	TypeUInt16      TypeKind = 4
	TypeInt32       TypeKind = 5
	TypeUInt32      TypeKind = 6
	TypeInt64       TypeKind = 7
	TypeUInt64      TypeKind = 8
	TypeFloat32     TypeKind = 9
	TypeFloat64     TypeKind = 10
	TypeStringFixed TypeKind = 11
	TypeStringVLS   TypeKind = 12
	TypeAlias       TypeKind = 19
	TypeEnum        TypeKind = 20
	TypeUnion       TypeKind = 21
	TypeStruct      TypeKind = 30
)

type Type struct {
	Namespace   *Namespace
	File        *File
	Kind        TypeKind
	Offset      int
	Align       int
	Length      int
	LengthConst *Const
	Enum        *Enum
	Message     *Struct
	Alias       *Alias
	Union       *Union
}

type Namespaces struct {
	Files      map[string]*File
	Namespaces map[string]*Namespace
}

type Namespace struct {
	Namespaces      *Namespaces
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

type File struct {
	Path             string
	Namespaces       *Namespaces
	currentNamespace *Namespace
	Alias            []*Alias
	AliasByName      map[string]*Alias
	Constants        []*Const
	ConstantsByName  map[string]*Const
	Enums            []*Enum
	EnumsByName      map[string]*Enum
	EnumOptions      map[string]*EnumOption
	Structs          []*Struct
	StructsByName    map[string]*Struct
}

type Alias struct {
	Namespace *Namespace
	File      *File
	Name      string
	Type      Type
}

type Const struct {
	Namespace *Namespace
	File      *File
	Name      string
	Type      Type
	Length    int
	Value     *Value
	Comment   string
}

type Struct struct {
	Namespace    *Namespace
	File         *File
	Name         string
	Align        int
	Size         int
	VLS          bool
	Fields       []*Field
	FieldsByName map[string]*Field
}

type Field struct {
	Name string
	Type Type
	Init *Value
}

type ValueType int

const (
	ValueTypeInt        ValueType = 0
	ValueTypeFloat      ValueType = 1
	ValueTypeString     ValueType = 2
	ValueTypeConst      ValueType = 3
	ValueTypeEnumOption ValueType = 4
	ValueTypeSizeof     ValueType = 5
)

type Value struct {
	File       *File
	Namespace  *Namespace
	Type       ValueType
	Int        int64
	Float      float64
	String     string
	Const      *Const
	EnumOption *EnumOption
	Sizeof     string
}

type Enum struct {
	File          *File
	Namespace     *Namespace
	Name          string
	Type          TypeKind
	Options       []*EnumOption
	OptionsByName map[string]*EnumOption
}

type EnumOption struct {
	File      *File
	Namespace *Namespace
	Enum      *Enum
	Name      string
	Value     int64
	Comment   string
}

type Union struct {
	Fields []*Field
}

func (s *Struct) Layout() {
	maxAlign := 0
	// Ensure length and alignment is set
	for _, field := range s.Fields {
		setLengthAndAlign(&field.Type)
		if maxAlign < field.Type.Align {
			maxAlign = field.Type.Align
		}
		if field.Type.Kind == TypeStringVLS {
			s.VLS = true
		}
		if s.Align < field.Type.Align {
			field.Type.Align = s.Align
			if field.Type.Union != nil {
				for _, f := range field.Type.Union.Fields {
					f.Type.Align = s.Align
				}
				if field.Type.Kind == TypeStringVLS {
					s.VLS = true
				}
			}
		}
	}

	offset := 0

	for i := 0; i < len(s.Fields); i++ {
		field := s.Fields[i]
		offset = align(offset, field.Type.Align)
		field.Type.Offset = offset

		if field.Type.Union != nil {
			for _, f := range field.Type.Union.Fields {
				f.Type.Offset = offset
			}
		}

		offset += field.Type.Length
	}
	if maxAlign < s.Align {
		s.Size = align(offset, maxAlign)
	} else {
		s.Size = align(offset, s.Align)
	}
}

func align(offset, align int) int {
	if offset == 0 || align == 0 {
		return 0
	}
	extras := offset % align
	if extras == 0 {
		return offset
	}
	return ((offset / align) + 1) * align
}

func parseError(lineIndex int, line, message string) error {
	return fmt.Errorf("line: %d, '%s' -> %s", lineIndex+1, line, message)
}

func NewNamespaces() *Namespaces {
	return &Namespaces{
		Files:      make(map[string]*File),
		Namespaces: make(map[string]*Namespace),
	}
}

func (namespaces *Namespaces) AddHeaders(paths ...string) error {
	for _, path := range paths {
		if _, err := namespaces.AddHeader(path); err != nil {
			return err
		}
	}
	return nil
}

func (namespaces *Namespaces) AddHeader(path string) (*File, error) {
	if namespaces.Files[path] != nil {
		return namespaces.Files[path], nil
	}
	contents_, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	contents := string(contents_)
	var (
		file = &File{
			Path:             path,
			Namespaces:       namespaces,
			currentNamespace: namespaces.GetNamespace(""), // Global namespace
			AliasByName:      make(map[string]*Alias),
			ConstantsByName:  make(map[string]*Const),
			EnumsByName:      make(map[string]*Enum),
			StructsByName:    make(map[string]*Struct),
			EnumOptions:      make(map[string]*EnumOption),
		}
		packStack    []int
		blockName    string
		blockExtends Type
		blockType    TypeKind
		block        []string
		state        = 0 // 0 = root, 1 == expect '{', 2 == expect '};'
		curlyCount   = 0
	)

	lines := strings.Split(contents, "\n")
	for lineIndex, line := range lines {
		line = strings.TrimSpace(line)

		switch state {
		case 1:
			index := strings.Index(line, "{")
			if index == -1 {
				continue
			}
			line = strings.TrimSpace(line[index+1:])
			if len(line) > 0 {
				block = append(block, line)
			}
			state = 2
			curlyCount = 0
			continue
		case 2:
			for _, c := range line {
				switch c {
				case '{':
					curlyCount++
				case '}':
					curlyCount--
				}
			}

			if curlyCount == -1 {
				//if line == "};" && curlyCount == 0 {
				switch blockType {
				case TypeStruct:
					pack := 8
					if len(packStack) > 0 {
						pack = packStack[len(packStack)-1]
					}
					_, err = file.parseStruct(pack, blockName, block)
					if err != nil {
						return nil, parseError(lineIndex, line, err.Error())
					}
				case TypeEnum:
					_, err = file.parseEnum(blockName, blockExtends.Kind, block)
					if err != nil {
						return nil, parseError(lineIndex, line, err.Error())
					}
				}
				state = 0
				continue
			} else {
				block = append(block, line)
				continue
			}
		}

		switch {
		case strings.HasPrefix(line, "namespace"):
			line = strings.TrimSpace(line[len("namespace"):])
			index := strings.LastIndex(line, "{")
			namespaceName := ""
			if index > -1 {
				namespaceName = strings.TrimSpace(line[0:index])
			} else {
				index = strings.Index(line, "/")
				if index > -1 {
					namespaceName = strings.TrimSpace(line[0:index])
				} else {
					namespaceName = strings.TrimSpace(line)
				}
			}
			file.currentNamespace = namespaces.GetNamespace(namespaceName)

		case strings.HasPrefix(line, "const"):
			line = strings.TrimSpace(line[len("const"):])
			_, err = file.parseConst(line)
			if err != nil {
				return nil, parseError(lineIndex, line, err.Error())
			}

		case strings.HasPrefix(line, "enum"):
			line = strings.TrimSpace(line[len("enum"):])

			blockType = TypeEnum
			block = nil

			index := strings.Index(line, ":")
			if index == -1 {
				return nil, parseError(lineIndex, line, "expected ':' and data type")
			}
			blockName = strings.TrimSpace(line[0:index])
			line = strings.TrimSpace(line[index+1:])

			curlyCount = 0
			index = strings.Index(line, "{")
			if index > -1 {
				blockExtends = file.dataTypeOf(strings.TrimSpace(line[0:index]))
				state = 2
			} else {
				blockExtends = file.dataTypeOf(strings.TrimSpace(line))
				state = 1
			}

		case strings.HasPrefix(line, "#pragma"):
			line = strings.TrimSpace(line[len("#pragma"):])
			if !strings.HasPrefix(line, "pack(") {
				continue
			}
			line = line[len("pack("):]
			line = strings.ReplaceAll(line, ")", "")
			line = strings.TrimSpace(line)
			parts := strings.Split(line, ",")
			switch strings.TrimSpace(parts[0]) {
			case "push":
				align, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
				if err != nil {
					return nil, parseError(lineIndex, line, "invalid #pragma pack(push, #ERROR_VALUE#) "+strings.TrimSpace(parts[1]))
				}
				packStack = append(packStack, int(align))
			case "pop":
				if len(packStack) > 0 {
					packStack = packStack[0 : len(packStack)-1]
				}
			}

		case strings.HasPrefix(line, "struct"):
			line = strings.TrimSpace(line[len("struct"):])
			blockType = TypeStruct
			block = nil
			blockName = ""
			blockExtends = Type{}
			curlyCount = 0

			index := strings.Index(line, "{")
			if index > -1 {
				blockName = strings.TrimSpace(line[0:index])
				state = 2
			} else {
				blockName = line
				state = 1
			}
			index = strings.Index(blockName, "//")
			if index > -1 {
				blockName = strings.TrimSpace(blockName[0:index])
			}
			index = strings.Index(blockName, "/*")
			if index > -1 {
				blockName = strings.TrimSpace(blockName[0:index])
			}

		case strings.HasPrefix(line, "typedef"):
			line = strings.TrimSpace(line[len("typedef"):])
			index := strings.LastIndex(line, " ")
			if index == -1 {
				return nil, parseError(lineIndex, line, "invalid typedef")
			}
			base := strings.TrimSpace(line[0:index])
			name := strings.TrimSpace(line[index+1:])
			if strings.HasSuffix(name, ";") {
				name = strings.TrimSpace(name[0 : len(name)-1])
			}
			alias := &Alias{
				File:      file,
				Namespace: file.currentNamespace,
				Name:      name,
				Type:      file.dataTypeOf(base),
			}
			file.Alias = append(file.Alias, alias)
			file.AliasByName[alias.Name] = alias
			file.currentNamespace.Alias = append(file.currentNamespace.Alias, alias)
			file.currentNamespace.AliasByName[alias.Name] = alias
		}
	}

	for _, s := range file.Structs {
		s.Layout()
	}

	namespaces.Files[path] = file
	return file, nil
}

func (namespaces *Namespaces) GetNamespace(name string) *Namespace {
	namespace := namespaces.Namespaces[name]
	if namespace == nil {
		namespace = &Namespace{
			Namespaces:      namespaces,
			Name:            name,
			Alias:           nil,
			AliasByName:     make(map[string]*Alias),
			Constants:       nil,
			ConstantsByName: make(map[string]*Const),
			Enums:           nil,
			EnumsByName:     make(map[string]*Enum),
			EnumOptions:     make(map[string]*EnumOption),
			Structs:         nil,
			StructsByName:   make(map[string]*Struct),
		}
		namespaces.Namespaces[name] = namespace
	}
	return namespace
}

//func Parse(imports []*File, contents string) (*File, error) {
//	var (
//		file = &File{
//			AliasByName:     make(map[string]*Alias),
//			ConstantsByName: make(map[string]*Const),
//			EnumsByName:     make(map[string]*Enum),
//			StructsByName:   make(map[string]*Struct),
//			EnumOptions:     make(map[string]*EnumOption),
//		}
//		err          error
//		packStack    []int
//		blockName    string
//		blockExtends Type
//		blockType    TypeKind
//		block        []string
//		state        = 0 // 0 = root, 1 == expect '{', 2 == expect '};'
//		curlyCount   = 0
//	)
//
//	if len(imports) > 0 {
//		for _, imp := range imports {
//			if imp == nil {
//				continue
//			}
//			file.Imports[imp.Namespace] = imp
//		}
//	}
//	lines := strings.Split(contents, "\n")
//	for lineIndex, line := range lines {
//		line = strings.TrimSpace(line)
//
//		switch state {
//		case 1:
//			index := strings.Index(line, "{")
//			if index == -1 {
//				continue
//			}
//			line = strings.TrimSpace(line[index+1:])
//			if len(line) > 0 {
//				block = append(block, line)
//			}
//			state = 2
//			curlyCount = 0
//			continue
//		case 2:
//			for _, c := range line {
//				switch c {
//				case '{':
//					curlyCount++
//				case '}':
//					curlyCount--
//				}
//			}
//
//			if curlyCount == -1 {
//				//if line == "};" && curlyCount == 0 {
//				switch blockType {
//				case TypeStruct:
//					pack := 8
//					if len(packStack) > 0 {
//						pack = packStack[len(packStack)-1]
//					}
//					_, err = file.parseStruct(pack, blockName, block)
//					if err != nil {
//						return nil, parseError(lineIndex, line, err.Error())
//					}
//				case TypeEnum:
//					_, err = file.parseEnum(blockName, blockExtends.Kind, block)
//					if err != nil {
//						return nil, parseError(lineIndex, line, err.Error())
//					}
//				}
//				state = 0
//				continue
//			} else {
//				block = append(block, line)
//				continue
//			}
//		}
//
//		switch {
//		case strings.HasPrefix(line, "namespace"):
//			line = strings.TrimSpace(line[len("namespace"):])
//			index := strings.LastIndex(line, "{")
//			if index > -1 {
//				file.Namespace = strings.TrimSpace(line[0:index])
//			} else {
//				index = strings.Index(line, "/")
//				if index > -1 {
//					file.Namespace = strings.TrimSpace(line[0:index])
//				} else {
//					file.Namespace = strings.TrimSpace(line)
//				}
//			}
//
//		case strings.HasPrefix(line, "const"):
//			line = strings.TrimSpace(line[len("const"):])
//			_, err = file.parseConst(line)
//			if err != nil {
//				return nil, parseError(lineIndex, line, err.Error())
//			}
//
//		case strings.HasPrefix(line, "enum"):
//			line = strings.TrimSpace(line[len("enum"):])
//
//			blockType = TypeEnum
//			block = nil
//
//			index := strings.Index(line, ":")
//			if index == -1 {
//				return nil, parseError(lineIndex, line, "expected ':' and data type")
//			}
//			blockName = strings.TrimSpace(line[0:index])
//			line = strings.TrimSpace(line[index+1:])
//
//			curlyCount = 0
//			index = strings.Index(line, "{")
//			if index > -1 {
//				blockExtends = file.dataTypeOf(strings.TrimSpace(line[0:index]))
//				state = 2
//			} else {
//				blockExtends = file.dataTypeOf(strings.TrimSpace(line))
//				state = 1
//			}
//
//		case strings.HasPrefix(line, "#pragma"):
//			line = strings.TrimSpace(line[len("#pragma"):])
//			if !strings.HasPrefix(line, "pack(") {
//				continue
//			}
//			line = line[len("pack("):]
//			line = strings.ReplaceAll(line, ")", "")
//			line = strings.TrimSpace(line)
//			parts := strings.Split(line, ",")
//			switch strings.TrimSpace(parts[0]) {
//			case "push":
//				align, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
//				if err != nil {
//					return nil, parseError(lineIndex, line, "invalid #pragma pack(push, #ERROR_VALUE#) "+strings.TrimSpace(parts[1]))
//				}
//				packStack = append(packStack, int(align))
//			case "pop":
//				if len(packStack) > 0 {
//					packStack = packStack[0 : len(packStack)-1]
//				}
//			}
//
//		case strings.HasPrefix(line, "struct"):
//			line = strings.TrimSpace(line[len("struct"):])
//			blockType = TypeStruct
//			block = nil
//			blockName = ""
//			blockExtends = Type{}
//			curlyCount = 0
//
//			index := strings.Index(line, "{")
//			if index > -1 {
//				blockName = strings.TrimSpace(line[0:index])
//				state = 2
//			} else {
//				blockName = line
//				state = 1
//			}
//			index = strings.Index(blockName, "//")
//			if index > -1 {
//				blockName = strings.TrimSpace(blockName[0:index])
//			}
//			index = strings.Index(blockName, "/*")
//			if index > -1 {
//				blockName = strings.TrimSpace(blockName[0:index])
//			}
//
//		case strings.HasPrefix(line, "typedef"):
//			line = strings.TrimSpace(line[len("typedef"):])
//			index := strings.LastIndex(line, " ")
//			if index == -1 {
//				return nil, parseError(lineIndex, line, "invalid typedef")
//			}
//			base := strings.TrimSpace(line[0:index])
//			name := strings.TrimSpace(line[index+1:])
//			if strings.HasSuffix(name, ";") {
//				name = strings.TrimSpace(name[0 : len(name)-1])
//			}
//			alias := &Alias{Name: name}
//			alias.Type = file.dataTypeOf(base)
//
//			file.Alias = append(file.Alias, alias)
//			file.AliasByName[alias.Name] = alias
//		}
//	}
//
//	for _, s := range file.Structs {
//		s.Layout()
//	}
//
//	return file, nil
//}

func (f *File) parseConst(line string) (*Const, error) {
	line = strings.TrimSpace(line)
	var (
		constant = &Const{
			File:      f,
			Namespace: f.currentNamespace,
		}
		index = strings.Index(line, "//")
		err   error
	)
	if index > -1 {
		constant.Comment = strings.TrimSpace(line[index:])
		line = strings.TrimSpace(line[0:index])
	}
	index = strings.Index(line, "=")
	if index == -1 {
		return nil, errors.New("constant missing '=' expression")
	}
	constant.Value, err = f.parseValue(strings.TrimSpace(line[index+1:]))
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line[0:index])

	line = strings.ReplaceAll(line, "\t", " ")
	index = strings.LastIndex(line, " ")
	constant.Name = strings.TrimSpace(line[index+1:])
	constant.Type = f.dataTypeOf(strings.TrimSpace(line[0:index]))

	if f.ConstantsByName[constant.Name] != nil {
		return nil, errors.New("duplicate constant declared: " + line)
	}
	f.Constants = append(f.Constants, constant)
	f.ConstantsByName[constant.Name] = constant
	f.currentNamespace.Constants = append(f.currentNamespace.Constants, constant)
	f.currentNamespace.ConstantsByName[constant.Name] = constant
	return constant, nil
}

func setLengthAndAlign(t *Type) {
	kind := t.Kind
	if kind == TypeUnion {
		for _, field := range t.Union.Fields {
			setLengthAndAlign(&field.Type)
			if field.Type.Length > t.Length {
				t.Length = field.Type.Length
			}
			if field.Type.Align > t.Align {
				t.Align = field.Type.Align
			}
		}
		return
	}
	if kind == TypeEnum {
		kind = t.Enum.Type
	} else if kind == TypeAlias {
		kind = t.Alias.Type.Kind
	}
	switch kind {
	case TypeUnknown:
	case TypeInt8, TypeUInt8:
		t.Align = 1
		t.Length = 1
	case TypeInt16, TypeUInt16:
		t.Align = 2
		t.Length = 2
	case TypeInt32, TypeUInt32:
		t.Align = 4
		t.Length = 4
	case TypeInt64, TypeUInt64:
		t.Align = 8
		t.Length = 8
	case TypeFloat32:
		t.Align = 4
		t.Length = 4
	case TypeFloat64:
		t.Align = 8
		t.Length = 8
	case TypeStringFixed:
		t.Align = 1
	case TypeStringVLS:
		t.Align = 2
		t.Length = 4
	case TypeAlias:
		// impossible
	case TypeEnum:
		// impossible
	case TypeStruct:
		if t.Message != nil {
			t.Align = t.Message.Size
		}
	}
}

func (f *File) dataTypeOf(value string) Type {
	value = strings.TrimSpace(value)
	switch value {
	case "char":
		return Type{Kind: TypeStringFixed}
	case "int8_t":
		return Type{Kind: TypeInt8, Align: 1, Length: 1}
	case "uint8_t", "unsigned char":
		return Type{Kind: TypeUInt8, Align: 1, Length: 1}
	case "int16_t", "short":
		return Type{Kind: TypeInt16, Align: 2, Length: 2}
	case "uint16_t", "unsigned short":
		return Type{Kind: TypeUInt16, Align: 2, Length: 2}
	case "int32_t", "int":
		return Type{Kind: TypeInt32, Align: 4, Length: 4}
	case "uint32_t", "unsigned int", "unsigned long":
		return Type{Kind: TypeUInt32, Align: 4, Length: 4}
	case "int64_t", "long long":
		return Type{Kind: TypeInt64, Align: 8, Length: 8}
	case "uint64_t", "unsigned long long":
		return Type{Kind: TypeUInt64, Align: 8, Length: 8}
	case "float":
		return Type{Kind: TypeFloat32, Align: 4, Length: 4}
	case "double":
		return Type{Kind: TypeFloat64, Align: 8, Length: 8}
	case "DTC_VLS::vls_t", "vls_t":
		return Type{Kind: TypeStringVLS, Align: 2, Length: 4}
	}

	namespace := f.currentNamespace
	index := strings.LastIndex(value, "::")
	if index > -1 {
		namespaceName := strings.TrimSpace(value[0:index])
		namespace = f.Namespaces.GetNamespace(namespaceName)
		value = strings.TrimSpace(value[index+2:])
	}

	enum := namespace.EnumsByName[value]
	if enum != nil {
		t := Type{
			File:      enum.File,
			Namespace: enum.Namespace,
			Kind:      TypeEnum,
			Enum:      enum,
		}
		setLengthAndAlign(&t)
		return t
	}

	message := namespace.StructsByName[value]
	if message != nil {
		t := Type{
			File:      message.File,
			Namespace: message.Namespace,
			Kind:      TypeStruct,
			Message:   message,
		}
		setLengthAndAlign(&t)
		return t
	}

	alias := namespace.AliasByName[value]
	if alias != nil {
		t := Type{
			File:      alias.File,
			Namespace: alias.Namespace,
			Kind:      TypeAlias,
			Alias:     alias,
		}
		setLengthAndAlign(&t)
		return t
	}

	return Type{Kind: TypeUnknown}
}

func (f *File) parseStruct(pack int, name string, lines []string) (*Struct, error) {
	var (
		message = &Struct{
			File:         f,
			Namespace:    f.currentNamespace,
			Name:         name,
			Align:        pack,
			Fields:       nil,
			FieldsByName: make(map[string]*Field),
		}
		funcLines  []string
		funcDecl   string
		state      = 0 // 0 = begin, 1 = waiting for end '}'
		err        error
		curlyCount = 0
		comment    = ""
	)
	_ = comment

	for i, line := range lines {
		line = strings.TrimSpace(line)
		index := strings.Index(line, "//")
		if index > -1 {
			comment = strings.TrimSpace(line[index:])
			line = strings.TrimSpace(line[0:index])
		}
		_ = i

		if len(line) == 0 {
			continue
		}

		if state == 1 {
			for _, c := range line {
				switch c {
				case '{':
					curlyCount++
				case '}':
					curlyCount--
				}
			}
			if curlyCount == -1 {
				state = 0

				if strings.Contains(funcDecl, "Clear()") {
					// Process default field values
					for _, line = range funcLines {
						if !strings.Contains(line, "=") {
							continue
						}

						parts := strings.Split(line, "=")
						field := message.FieldsByName[strings.TrimSpace(parts[0])]
						if field == nil {
							continue
						}
						field.Init, err = f.parseValue(strings.TrimSpace(parts[1]))
						if err != nil {
							return nil, err
						}
					}
				} else if funcDecl == "union" {
					union := Type{
						Kind:  TypeUnion,
						Union: &Union{},
					}

					for _, line := range funcLines {
						// Should be field declaration
						field, err := f.parseField(line)
						if err != nil {
							return nil, err
						}

						union.Union.Fields = append(union.Union.Fields, field)

						if field.Type.Length > union.Length {
							union.Length = field.Type.Length
						}
						if field.Type.Align > union.Align {
							union.Align = field.Type.Align
						}

						message.FieldsByName[field.Name] = field
					}

					message.Fields = append(message.Fields, &Field{
						Name: "",
						Type: union,
					})
				}

				funcDecl = ""
			} else {
				funcLines = append(funcLines, line)
			}
			continue
		}
		//if strings.HasPrefix(line, name) {
		//	// constructor
		//	continue
		//}

		if !strings.HasSuffix(line, ";") {
			index = strings.Index(line, "{")
			if index > -1 {
				funcLines = nil
				begin := strings.TrimSpace(line[0:index])
				if len(begin) > 0 {
					funcDecl = begin
				}
				line = strings.TrimSpace(line[index+1:])
				if len(line) > 0 {
					funcLines = append(funcLines, line)
				}
				curlyCount = 0
				state = 1
			} else if len(line) > 0 {
				funcDecl = line
			}

			continue
		}

		line = strings.TrimSpace(line[0 : len(line)-1])
		if strings.HasSuffix(line, "const") {
			continue
		}

		if strings.HasSuffix(line, ")") {
			if !strings.Contains(line, "sizeof(") {
				continue
			}
		}

		// Should be field declaration
		field, err := f.parseField(line)
		if err != nil {
			return nil, err
		}

		if field != nil {
			message.Fields = append(message.Fields, field)
			message.FieldsByName[field.Name] = field
		}
	}

	f.Structs = append(f.Structs, message)
	f.StructsByName[message.Name] = message
	f.currentNamespace.Structs = append(f.currentNamespace.Structs, message)
	f.currentNamespace.StructsByName[message.Name] = message

	return message, nil
}

func (f *File) parseField(line string) (*Field, error) {
	line = strings.TrimSpace(line)
	index := strings.Index(line, "//")
	comment := ""
	if index > -1 {
		comment = strings.TrimSpace(line[index:])
		line = strings.TrimSpace(line[0:index])
	}
	_ = comment

	if strings.HasSuffix(line, ";") {
		line = strings.TrimSpace(line[0 : len(line)-1])
	}

	var (
		expression = ""
		field      = &Field{}
		err        error
	)
	index = strings.Index(line, "=")
	if index > -1 {
		expression = strings.TrimSpace(line[index+1:])
		line = strings.TrimSpace(line[0:index])
	}

	line = strings.ReplaceAll(line, "\t", " ")
	index = strings.LastIndex(line, " ")
	if index == -1 {
		return nil, errors.New("expected a space between data type and name")
	}

	field.Type = f.dataTypeOf(strings.TrimSpace(line[0:index]))
	field.Name = strings.TrimSpace(line[index+1:])

	index = strings.Index(field.Name, "[")
	if index > -1 {
		endIndex := strings.Index(field.Name, "]")
		if index >= endIndex {
			return nil, errors.New("invalid fixed string declaration")
		}
		lengthExpression := strings.TrimSpace(field.Name[index+1 : endIndex])
		field.Name = strings.TrimSpace(field.Name[0:index])

		length, err := strconv.ParseInt(lengthExpression, 10, 64)
		if err != nil {
			value := f.findConstValue(lengthExpression)
			if value == nil || value.Const == nil {
				return nil, errors.New("could not determine fixed string length")
			}
			field.Type.LengthConst = value.Const
			field.Type.Length = int(value.Const.Value.Int)
		} else {
			field.Type.Length = int(length)
		}
	}

	if field.Name == "" {
		return nil, nil
	}

	if len(expression) > 0 {
		field.Init, err = f.parseValue(expression)
		if err != nil {
			return nil, err
		}
	}

	return field, nil
}

func (f *File) findConstValue(str string) *Value {
	index := strings.Index(str, "::")
	namespace := f.currentNamespace
	if index > -1 {
		namespace = f.Namespaces.GetNamespace(strings.TrimSpace(str[0:index]))
		str = strings.TrimSpace(str[index+1:])
	}
	constant := namespace.ConstantsByName[str]
	if constant != nil {
		return &Value{
			File:      constant.File,
			Namespace: namespace,
			Type:      ValueTypeConst,
			Int:       constant.Value.Int,
			Float:     constant.Value.Float,
			Const:     constant,
		}
	}

	option := namespace.EnumOptions[str]
	if option != nil {
		return &Value{
			File:       option.File,
			Namespace:  option.Namespace,
			Int:        option.Value,
			Type:       ValueTypeEnumOption,
			EnumOption: option,
		}
	}

	return nil
}

func (f *File) parseValue(str string) (*Value, error) {
	str = strings.TrimSpace(str)
	if len(str) > 1 && strings.HasSuffix(str, ";") {
		str = strings.TrimSpace(str[0 : len(str)-1])
	}

	index := strings.Index(str, "::")
	if index > -1 {
		return f.findConstValue(str), nil
	}

	if str == "{}" {
		return &Value{
			File:   f,
			Type:   ValueTypeString,
			String: "",
		}, nil
	}

	index = strings.Index(str, "sizeof(")
	if index > -1 {
		str = strings.TrimSpace(str[index+len("sizeof("):])
		index = strings.Index(str, ")")
		if index > -1 {
			str = strings.TrimSpace(str[0:index])
		}
		return &Value{
			File:   f,
			Type:   ValueTypeSizeof,
			Sizeof: str,
		}, nil
	}

	if strings.Index(str, "\"") > -1 {
		return &Value{
			File:   f,
			Type:   ValueTypeString,
			String: strings.ReplaceAll(str, "\"", ""),
		}, nil
	}
	if strings.Index(str, ".") > -1 {
		number, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return &Value{
				File:  f,
				Type:  ValueTypeFloat,
				Float: number,
			}, nil
		}
	}
	number, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return &Value{
			File: f,
			Type: ValueTypeInt,
			Int:  number,
		}, nil
	}

	return f.findConstValue(str), nil
}

func (f *File) parseEnum(name string, kind TypeKind, lines []string) (*Enum, error) {
	var (
		enum = &Enum{
			File:          f,
			Namespace:     f.currentNamespace,
			Name:          name,
			Type:          kind,
			OptionsByName: make(map[string]*EnumOption),
		}
		options = strings.Split(strings.Join(lines, ""), ",")
	)
	for _, option := range options {
		option = strings.TrimSpace(option)
		var (
			err   error
			opt   = &EnumOption{File: f, Namespace: f.currentNamespace, Enum: enum}
			index = strings.Index(option, "//")
		)
		if index == -1 {
			index = strings.Index(option, "/*")
		}
		if index > -1 {
			opt.Comment = strings.TrimSpace(option[index:])
			option = strings.TrimSpace(option[0:index])
		}

		index = strings.Index(option, "=")

		if index == -1 {
			opt.Name = strings.TrimSpace(option)
			// Find previous value
			if len(enum.Options) == 0 {
				opt.Value = 0
			} else {
				opt.Value = enum.Options[len(enum.Options)-1].Value + 1
			}
		} else {
			opt.Value, err = strconv.ParseInt(strings.TrimSpace(option[index+1:]), 10, 64)
			if err != nil {
				return nil, errors.New("invalid enum option declaration: " + option + " : " + err.Error())
			}
			opt.Name = strings.TrimSpace(option[0:index])
		}

		if enum.OptionsByName[opt.Name] != nil {
			return nil, errors.New("duplicate enum option declaration: " + option)
		}
		if f.currentNamespace.EnumOptions[opt.Name] != nil {
			return nil, errors.New("file: duplicate enum option name: " + option)
		}

		enum.Options = append(enum.Options, opt)
		enum.OptionsByName[opt.Name] = opt
		f.EnumOptions[opt.Name] = opt
		f.currentNamespace.EnumOptions[opt.Name] = opt
	}

	f.Enums = append(f.Enums, enum)
	f.EnumsByName[enum.Name] = enum
	f.currentNamespace.Enums = append(f.currentNamespace.Enums, enum)
	f.currentNamespace.EnumsByName[enum.Name] = enum

	return enum, nil
}
