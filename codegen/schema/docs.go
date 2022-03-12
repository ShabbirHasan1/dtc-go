package schema

import (
	"errors"
	"os"
	"strings"
)

type Docs struct {
	Messages       []*MessageDoc          `json:"-"`
	MessagesByName map[string]*MessageDoc `json:"messages"`
	Enums          []*TypeDoc             `json:"-"`
	TypesByName    map[string]*TypeDoc    `json:"types"`
}

func (d *Docs) MessageNamed(name string) *MessageDoc {
	if d == nil || d.MessagesByName == nil {
		return nil
	}
	for _, m := range d.MessagesByName {
		if m.Name == name {
			return m
		}
	}
	return nil
}

func (d *Docs) TypeNamed(name string) *TypeDoc {
	if d == nil || d.TypesByName == nil {
		return nil
	}
	for _, t := range d.TypesByName {
		if t.Name == name {
			return t
		}
	}
	return nil
}

func (d *Docs) NormalizeNames() {
	m := make(map[string]string)
	for _, msg := range d.MessagesByName {
		m[convertMessageNameToSnakeUpper(msg.Name)] = msg.Name
	}

	for _, msg := range d.MessagesByName {
		for i, _ := range msg.Description {
			for from, to := range m {
				msg.Description[i] = strings.ReplaceAll(msg.Description[i], from, to)
			}

			for _, f := range msg.Fields {
				for i, _ := range f.Description {
					for from, to := range m {
						f.Description[i] = strings.ReplaceAll(f.Description[i], from, to)
					}
				}
			}
		}
	}

	for _, msg := range d.TypesByName {
		for i, _ := range msg.Description {
			for from, to := range m {
				msg.Description[i] = strings.ReplaceAll(msg.Description[i], from, to)
			}

			for _, f := range msg.Options {
				for i, _ := range f.Description {
					for from, to := range m {
						f.Description[i] = strings.ReplaceAll(f.Description[i], from, to)
					}
				}
			}
		}
	}
}

func (d *Docs) Merge(other *Docs) *Docs {
	if other == nil {
		return d
	}
	if d == nil {
		return other
	}
	for _, message := range other.Messages {
		if d.MessagesByName[message.Name] == nil {
			d.Messages = append(d.Messages, message)
			d.MessagesByName[message.Name] = message
		}
	}
	for _, enum := range other.Enums {
		if d.TypesByName[enum.Name] == nil {
			d.Enums = append(d.Enums, enum)
			d.TypesByName[enum.Name] = enum
		}
	}
	return d
}

type TypeDoc struct {
	Name        string           `json:"name"`
	Short       string           `json:"short,omitempty"`
	Description []string         `json:"description,omitempty"`
	Options     []*EnumOptionDoc `json:"options,omitempty"`
}

func (td *TypeDoc) OptionNamed(name string) *EnumOptionDoc {
	if td == nil || len(td.Options) == 0 {
		return nil
	}
	for _, option := range td.Options {
		if option.Name == name {
			return option
		}
	}
	return nil
}

type EnumOptionDoc struct {
	Name        string   `json:"name"`
	Value       string   `json:"value,omitempty"`
	Description []string `json:"description,omitempty"`
}

type MessageDoc struct {
	Name         string      `json:"name"`
	Description  []string    `json:"description,omitempty"`
	Fields       []*FieldDoc `json:"fields"`
	ServerClient bool        `json:"serverToClient"`
	ClientServer bool        `json:"clientToServer"`
}

func (md *MessageDoc) FieldNamed(name string) *FieldDoc {
	if md == nil || len(md.Fields) == 0 {
		return nil
	}

	for _, field := range md.Fields {
		if field.Name == name {
			return field
		}
		if strings.HasPrefix(field.Name, "m_") && field.Name[2:] == name {
			return field
		}
	}
	return nil
}

type FieldDoc struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Description []string `json:"description,omitempty"`
}

func ParseMessageDocs(paths ...string) (*Docs, error) {
	docs := &Docs{
		MessagesByName: make(map[string]*MessageDoc),
		TypesByName:    make(map[string]*TypeDoc),
	}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		lines := strings.Split(string(data), "\n")
	OUTER:
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			line = strings.TrimSpace(line)
			if len(line) == 0 {
				continue
			}

			index := strings.IndexByte(line, '[')
			if index == -1 {
				continue
			}
			if index == 0 {
				return nil, errors.New("expected message header line")
			}
			line = line[index+1:]
			index = strings.Index(line, " ")
			message := &MessageDoc{
				Name: strings.TrimSpace(line[0:index]),
			}

			if len(message.Name) == 0 {
				return nil, errors.New("message name is empty: " + line)
			}
			if message.Name == "Server" {
				println("Hey")
			}

			index = strings.IndexByte(line, ']')
			if index == -1 {
				return nil, errors.New("expected ']' in message header line")
			}
			line = strings.TrimSpace(line[index+1:])
			message.ServerClient = strings.Contains(line, "Server >> Client")
			message.ClientServer = strings.Contains(line, "Client >> Server")

			i++
			line = lines[i]
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "[") {
				i++
				line = lines[i]
				line = strings.TrimSpace(line)
			}
			for len(line) == 0 {
				i++
				line = lines[i]
				line = strings.TrimSpace(line)
			}
			for !strings.HasPrefix(line, "Field Name") {
				message.Description = append(message.Description, line)

				i++
				line = strings.TrimSpace(lines[i])
			}

			// Remove last empty line
			if len(message.Description) > 0 && len(message.Description[len(message.Description)-1]) == 0 {
				message.Description = message.Description[0 : len(message.Description)-1]
			}

			docs.Messages = append(docs.Messages, message)
			docs.MessagesByName[message.Name] = message

			i++
			line = strings.TrimSpace(lines[i])
			for len(line) == 0 {
				i++
				line = strings.TrimSpace(lines[i])
			}

			for ; i < len(lines); i++ {
				line = strings.TrimSpace(lines[i])
				index = strings.IndexByte(line, '[')
				// Is next message?
				if index > 0 {
					if !strings.Contains(strings.TrimSpace(line[0:index]), " ") {
						i--
						if len(message.Fields) > 0 {
							f := message.Fields[len(message.Fields)-1]
							if len(f.Description) > 0 && len(f.Description[len(f.Description)-1]) == 0 {
								f.Description = f.Description[0 : len(f.Description)-1]
							}
						}
						continue OUTER
					}
					index = -1
				}
				if index == -1 {
					if len(message.Fields) == 0 {
						continue
					}
					message.Fields[len(message.Fields)-1].Description = append(message.Fields[len(message.Fields)-1].Description, line)
					continue
				}

				if len(message.Fields) > 0 {
					f := message.Fields[len(message.Fields)-1]
					if len(f.Description) > 0 && len(f.Description[len(f.Description)-1]) == 0 {
						f.Description = f.Description[0 : len(f.Description)-1]
					}
				}

				line = line[1:]
				index = strings.IndexByte(line, ']')
				if index == -1 {
					return nil, errors.New("expected ']' when parsing field type and name")
				}
				message.Fields = append(message.Fields, &FieldDoc{
					Name: strings.TrimSpace(line[index+1:]),
					Type: strings.TrimSpace(line[0:index]),
				})
			}

			if len(message.Fields) > 0 {
				f := message.Fields[len(message.Fields)-1]
				if len(f.Description) > 0 && len(f.Description[len(f.Description)-1]) == 0 {
					f.Description = f.Description[0 : len(f.Description)-1]
				}
			}
		}
	}

	return docs, nil
}

func ParseTypeDocs(path string) (*Docs, error) {
	docs := &Docs{
		TypesByName: make(map[string]*TypeDoc),
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
OUTER:
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		index := strings.IndexByte(line, '[')
		if index == -1 {
			continue
		}
		line = line[index+1:]
		index = strings.Index(line, "]")
		enum := &TypeDoc{
			Name: strings.TrimSpace(line[0:index]),
		}

		if len(enum.Name) == 0 {
			return nil, errors.New("enum name is empty: " + line)
		}

		line = strings.TrimSpace(line[index+1:])
		enum.Short = line

		i++
		line = lines[i]
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "[") {
			i++
			line = lines[i]
			line = strings.TrimSpace(line)
		}
		for len(line) == 0 {
			i++
			line = lines[i]
			line = strings.TrimSpace(line)
		}
		for !strings.Contains(line, "=") && !strings.Contains(line, "[") {
			enum.Description = append(enum.Description, line)

			i++
			line = strings.TrimSpace(lines[i])
		}

		// Remove last empty line
		if len(enum.Description) > 0 && len(enum.Description[len(enum.Description)-1]) == 0 {
			enum.Description = enum.Description[0 : len(enum.Description)-1]
		}

		docs.Enums = append(docs.Enums, enum)
		docs.TypesByName[enum.Name] = enum

		if !strings.Contains(line, "=") || strings.IndexByte(line, '[') == 0 {
			i--
			continue
		}

		for ; i < len(lines); i++ {
			line = strings.TrimSpace(lines[i])
			parts := strings.Split(line, "=")
			index = strings.IndexByte(line, '[')
			// Is next message?
			if index > -1 {
				i--
				continue OUTER
			}

			if len(parts) == 1 {
				continue
			}

			option := &EnumOptionDoc{
				Name:  strings.TrimSpace(parts[0]),
				Value: strings.TrimSpace(parts[1]),
			}

		LOOP:
			for i, c := range option.Value {
				switch c {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				default:
					if i == 0 {
						continue
					}
					description := strings.TrimSpace(option.Value[i:])
					option.Value = option.Value[0:i]

					if strings.HasPrefix(description, ".") {
						description = strings.TrimSpace(description[1:])
					}
					if len(description) > 0 {
						option.Description = append(option.Description, description)
					}
					break LOOP
				}
			}

			enum.Options = append(enum.Options, option)
		}
	}
	return docs, nil
}

func toSnakeCase(name string) string {
	out := make([]byte, 0, len(name)*2)

	// NewOCOOrder
	// TradeOrder
	for i := 0; i < len(name); i++ {
		var (
			isUpper     = string(name[i]) != strings.ToLower(string(name[i]))
			nextIsLower = false
			prevIsLower = false
		)
		if i > 0 {
			prevIsLower = string(name[i-1]) == strings.ToLower(string(name[i-1]))
		}
		if i < len(name)-1 {
			nextIsLower = strings.ToLower(string(name[i+1])) == string(name[i+1])
		}
		if isUpper {
			if prevIsLower || nextIsLower {
				if len(out) > 0 {
					out = append(out, '_')
				}
			}
		}

		out = append(out, strings.ToLower(string(name[i]))...)
	}

	return string(out)
}

func ToSnakeCaseUpper(name string) string {
	return strings.ToUpper(toSnakeCase(name))
}

func convertMessageNameToSnakeUpper(name string) string {
	if strings.HasPrefix(name, "s_") {
		name = name[2:]
	}
	return ToSnakeCaseUpper(name)
}
