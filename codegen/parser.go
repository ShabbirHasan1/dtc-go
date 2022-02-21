package codegen

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
		blockType    Kind
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
				case KindStruct:
					pack := 8
					if len(packStack) > 0 {
						pack = packStack[len(packStack)-1]
					}
					_, err = file.parseStruct(pack, blockName, block)
					if err != nil {
						return nil, parseError(lineIndex, line, err.Error())
					}
				case KindEnum:
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

			blockType = KindEnum
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
				blockExtends = file.typeOf(strings.TrimSpace(line[0:index]))
				state = 2
			} else {
				blockExtends = file.typeOf(strings.TrimSpace(line))
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
			blockType = KindStruct
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
				Type:      file.typeOf(base),
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
	constant.Type = f.typeOf(strings.TrimSpace(line[0:index]))

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
	if kind == KindUnion {
		for _, field := range t.Union.Fields {
			setLengthAndAlign(&field.Type)
			if field.Type.Size > t.Size {
				t.Size = field.Type.Size
			}
			if field.Type.Align > t.Align {
				t.Align = field.Type.Align
			}
		}
		return
	}
	if kind == KindEnum {
		kind = t.Enum.Type
	} else if kind == KindAlias {
		kind = t.Alias.Type.Kind
	}
	switch kind {
	case KindUnknown:
	case KindInt8, KindUint8:
		t.Align = 1
		t.Size = 1
	case KindInt16, KindUint16:
		t.Align = 2
		t.Size = 2
	case KindInt32, KindUint32:
		t.Align = 4
		t.Size = 4
	case KindInt64, KindUint64:
		t.Align = 8
		t.Size = 8
	case KindFloat32:
		t.Align = 4
		t.Size = 4
	case KindFloat64:
		t.Align = 8
		t.Size = 8
	case KindStringFixed:
		t.Align = 1
	case KindStringVLS:
		t.Align = 2
		t.Size = 4
	case KindAlias:
		// impossible
	case KindEnum:
		// impossible
	case KindStruct:
		if t.Message != nil {
			t.Align = t.Message.Size
		}
	}
}

func (f *File) typeOf(value string) Type {
	value = strings.TrimSpace(value)
	switch value {
	case "char":
		return Type{Kind: KindStringFixed}
	case "int8_t":
		return Type{Kind: KindInt8, Align: 1, Size: 1}
	case "uint8_t", "unsigned char":
		return Type{Kind: KindUint8, Align: 1, Size: 1}
	case "int16_t", "short":
		return Type{Kind: KindInt16, Align: 2, Size: 2}
	case "uint16_t", "unsigned short":
		return Type{Kind: KindUint16, Align: 2, Size: 2}
	case "int32_t", "int":
		return Type{Kind: KindInt32, Align: 4, Size: 4}
	case "uint32_t", "unsigned int", "unsigned long":
		return Type{Kind: KindUint32, Align: 4, Size: 4}
	case "int64_t", "long long":
		return Type{Kind: KindInt64, Align: 8, Size: 8}
	case "uint64_t", "unsigned long long":
		return Type{Kind: KindUint64, Align: 8, Size: 8}
	case "float":
		return Type{Kind: KindFloat32, Align: 4, Size: 4}
	case "double":
		return Type{Kind: KindFloat64, Align: 8, Size: 8}
	case "DTC_VLS::vls_t", "vls_t":
		return Type{Kind: KindStringVLS, Align: 2, Size: 4}
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
			Kind:      KindEnum,
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
			Kind:      KindStruct,
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
			Kind:      KindAlias,
			Alias:     alias,
		}
		setLengthAndAlign(&t)
		return t
	}

	return Type{Kind: KindUnknown}
}

func (f *File) parseStruct(pack int, name string, lines []string) (*Struct, error) {
	var (
		message = &Struct{
			File:         f,
			Namespace:    f.currentNamespace,
			Name:         name,
			MaxAlign:     pack,
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
						Kind:  KindUnion,
						Union: &Union{},
					}

					for _, line := range funcLines {
						// Should be field declaration
						field, err := f.parseField(line)
						if err != nil {
							return nil, err
						}

						union.Union.Fields = append(union.Union.Fields, field)

						if field.Type.Size > union.Size {
							union.Size = field.Type.Size
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
		if strings.HasSuffix(line, "const") || strings.HasSuffix(line, "noexcept") {
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

	field.Type = f.typeOf(strings.TrimSpace(line[0:index]))
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
			field.Type.SizeConst = value.Const
			field.Type.Size = int(value.Const.Value.Int)
		} else {
			field.Type.Size = int(length)
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

func (f *File) parseEnum(name string, kind Kind, lines []string) (*Enum, error) {
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
