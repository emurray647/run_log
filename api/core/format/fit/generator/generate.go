// +build ignore
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"
	"unicode"
)

func main() {

	dir, _ := os.Getwd()

	types, messages := read(dir + "/profiles/Profile.csv")

	// for _, message := range messages {
	// 	fmt.Println(message)
	// }

	// for _, message := range messages {
	// 	if message.MessageNum == 20 {
	// 		for _, field := range message.Fields {
	// 			fmt.Println(field.FieldType)
	// 		}
	// 	}
	// }

	dir = dir + "/generated/"
	fmt.Println(dir)

	os.Mkdir(dir, os.ModePerm)

	file, err := os.Create(dir + "messages.go")
	if err != nil {
		panic(err.Error())
	}

	packageTemplate.Execute(file, struct {
		Messages []Message
	}{
		Messages: messages,
	})

	file, err = os.Create(dir + "types.go")
	if err != nil {
		panic(err.Error())
	}
	typesTemplate.Execute(file, struct {
		Types []Type
	}{
		Types: types,
	})

	file, err = os.Create(dir + "units.go")
	if err != nil {
		panic(err.Error())
	}
	unitsTemplate.Execute(file, struct {
		Messages []Message
	}{
		Messages: messages,
	})

}

type TypeValue struct {
	Name  string
	Value uint
}

type Type struct {
	Name     string
	BaseType string
	Values   []TypeValue
}

type Unit struct {
	Scale  uint
	Offset int
	Unit   string
}

type FieldType struct {
	Name     string
	BaseType string
}

type Field struct {
	FieldNumber uint
	Name        string
	FieldType   FieldType //string // int or float or ...
	FieldUnit   Unit
}

type Message struct {
	Name       string
	MessageNum uint
	Fields     []Field
}

func stringToInt(s string) (int, error) {
	// fmt.Printf("The string is : %s\n", s)
	if len(s) > 2 && s[0] == '0' && (s[1] == 'X' || s[1] == 'x') {
		value, err := strconv.ParseInt(s[2:], 16, 32)
		return int(value), err
	}
	value, err := strconv.Atoi(s)
	return value, err

}

func read(path string) ([]Type, []Message) {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}

	reader := csv.NewReader(file)

	var types []Type
	var messages []Message

	messageMode := false
	var currentType *Type
	var currentMessage *Message

	for i := 0; ; i++ {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		// type vs message
		if line[0] == "Message Name" {
			messageMode = true
			continue
		} else if line[0] == "Type Name" {
			messageMode = false
			continue
		}

		// fmt.Println(line)

		if !messageMode {
			if len(line) < 4 {
				continue
			}
			if line[0] == "" && line[1] == "" && line[2] == "" && line[3] == "" {
				continue
			}
			// fmt.Println(line)
			// fmt.Println(len(line))

			if line[0] != "" {
				if currentType != nil {
					types = append(types, *currentType)
				}
				currentType = &Type{
					Name:     line[0],
					BaseType: line[1],
				}
			} else {
				value, err := stringToInt(line[3])
				if err != nil {
					panic(err.Error())
				}
				currentType.Values = append(currentType.Values, TypeValue{
					Name:  line[2],
					Value: uint(value),
				})
			}

		} else {
			// message mode
			if line[0] == "" && line[1] == "" && line[2] == "" {
				continue
			}

			if line[0] != "" {
				if currentMessage != nil {
					messages = append(messages, *currentMessage)
				}
				currentMessage = &Message{Name: line[0]}
			} else if line[1] != "" && line[2] != "" && line[3] != "" {
				value, _ := stringToInt(line[1])
				fieldType := FieldType{
					Name:     line[3],
					BaseType: "",
				}
				field := Field{
					FieldNumber: uint(value),
					Name:        line[2],
					FieldType:   fieldType,
				}
				if line[6] != "" {
					value, _ := stringToInt(line[6])
					field.FieldUnit.Scale = uint(value)
				} else {
					field.FieldUnit.Scale = 1
				}

				if line[7] != "" {
					value, _ := stringToInt(line[7])
					field.FieldUnit.Offset = value
				} else {
					field.FieldUnit.Offset = 0
				}
				field.FieldUnit.Unit = line[8]

				currentMessage.Fields = append(currentMessage.Fields, field)
			}

		}
	}

	var mesgNumType *Type
	for _, t := range types {
		if t.Name == "mesg_num" {
			mesgNumType = &t
			break
		}
	}

	for i := range messages {

		for _, typeValue := range mesgNumType.Values {
			if typeValue.Name == messages[i].Name {
				messages[i].MessageNum = typeValue.Value
			}
		}

		for j := range messages[i].Fields {
			messages[i].Fields[j].FieldType.BaseType = getBaseType(messages[i].Fields[j].FieldType.Name, types)
			if messages[i].Fields[j].Name == "type" {
				messages[i].Fields[j].Name = "type_"
			}
		}

	}

	// for _, t := range types {
	// 	fmt.Println(t)
	// }
	// for _, m := range messages {
	// 	if m.MessageNum == 20 {
	// 		// fmt.Println(m)
	// 		for j := range m.Fields {
	// 			fmt.Println(m.Fields[j])
	// 		}
	// 	}
	// 	// fmt.Printf("%d: %s\n", m.messageNum, m.name)
	// }

	return types, messages
}

func getBaseType(name string, types []Type) string {
	for _, t := range types {
		if name == t.Name {
			return t.BaseType
		}
	}
	return name
}

//////////////////////////////////////
// used for templates
/////////////////////////////////////

func typeToType(s string) string {
	correctionMap := make(map[string]string)
	correctionMap["sint8"] = "int8"
	correctionMap["sint16"] = "int16"
	correctionMap["sint32"] = "int32"
	correctionMap["sint64"] = "int64"
	correctionMap["uint8z"] = "uint8"
	correctionMap["uint16z"] = "uint16"
	correctionMap["uint32z"] = "uint32"
	correctionMap["uint64z"] = "uint64"

	if val, ok := correctionMap[s]; ok {
		return val
	}

	return s
}

func getTypeSize(s string) uint {
	switch s {
	case "enum", "sint8", "uint8", "uint8z", "string", "byte":
		return 1
	case "sint16", "uint16", "uint16z":
		return 2
	case "sint32", "uint32", "uint32z", "float32":
		return 4
	default:
		return 8
	}
}

func getUnderlyingGoType(s string) string {
	switch s {
	case "enum", "byte", "uint8", "uint8z":
		return "uint8"
	case "sint8":
		return "int8"
	case "uint16", "uint16z":
		return "uint16"
	case "sint16":
		return "int16"
	case "uint32", "uint32z":
		return "uint32"
	case "sint32":
		return "int32"
	case "float32":
		return "float32"
	case "uint64", "uint64z":
		return "uint64"
	case "sint64":
		return "int64"
	case "float64":
		return "float64"
	case "string":
		return "string"
	default:
		return ""

	}
}

func createDataMessageName(name string) string {
	var builder strings.Builder

	nextCap := true

	for _, c := range name {
		if c == '_' {
			nextCap = true
			continue
		}

		if nextCap {
			nextCap = false
			builder.WriteRune(unicode.ToUpper(c))
		} else {
			builder.WriteRune(c)
		}

	}

	builder.WriteString("DataMessage")

	return builder.String()
}

func createDataFieldName(name string) string {
	var builder strings.Builder
	nextCap := true
	for _, c := range name {
		if c == '_' {
			nextCap = true
			continue
		}
		if nextCap {
			nextCap = false
			builder.WriteRune(unicode.ToUpper(c))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func unitStringToEnumString(unit string) string {
	switch unit {
	case "year", "years":
		return "UNIT_YEAR"
	case "hr", "hour":
		return "UNIT_HOUR"
	case "min", "minutes":
		return "UNIT_MINUTE"
	case "s":
		return "UNIT_SECOND"
	case "ms":
		return "UNIT_MILLISECOND"
	case "m":
		return "UNIT_METER"
	case "mm":
		return "UNIT_MILLIMETER"
	case "m/s":
		return "UNIT_METERPERSECOND"
	case "steps":
		return "UNIT_COUNT"
	case "bpm", "rpm":
		return "UNIT_COUNTRATE"
	case "kcal":
		return "UNIT_KILOCALORIE"
	case "degrees":
		return "UNIT_DEGREES"
	case "semicircles":
		return "UNIT_SEMICIRCLES"
	}
	return "UNIT_UNKNOWN"

}

var unitsTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"toUpper":    strings.ToUpper,
	"unitToEnum": unitStringToEnumString,
}).Parse(`// Code generated; DO NOT EDIT

// Code generated; DO NOT EDIT

package generated

import "fmt"
import "math"

type Unit_e uint8

const (
	UNIT_YEAR Unit_e = iota
	UNIT_HOUR
	UNIT_MINUTE
	UNIT_SECOND
	UNIT_MILLISECOND
	UNIT_METER
	UNIT_MILLIMETER
	UNIT_METERPERSECOND
	UNIT_COUNT
	UNIT_COUNTRATE
	UNIT_KILOCALORIE
	UNIT_DEGREES
	UNIT_SEMICIRCLES
	UNIT_UNKNOWN
)

type Unit struct {
	Unit   Unit_e
	Scale  uint
	Offset int
}

var conversionMultiplier map[Unit_e]map[Unit_e]float64

func init() {
	conversionMultiplier = make(map[Unit_e]map[Unit_e]float64)

	conversionMultiplier[UNIT_SEMICIRCLES] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_SEMICIRCLES][UNIT_DEGREES] = 180.0 / math.Pow(2, 31)

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_HOUR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_HOUR][UNIT_HOUR] = 1.0

	conversionMultiplier[UNIT_MINUTE] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_MINUTE][UNIT_MINUTE] = 1.0
	
	conversionMultiplier[UNIT_SECOND] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_SECOND][UNIT_SECOND] = 1.0

	conversionMultiplier[UNIT_METER] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_METER][UNIT_METER] = 1.0

	conversionMultiplier[UNIT_SECOND] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_SECOND][UNIT_SECOND] = 1.0
/* TODO: ...


	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0
	*/
}

func Convert(value interface{}, startUnit Unit, goalUnit Unit_e) float64 {
	var startValue float64
	switch i := value.(type) {
	case float64:
		startValue = i
	case float32:
		startValue = float64(i)
	case int64:
		startValue = float64(i)
	case int32:
		startValue = float64(i)
	case int16:
		startValue = float64(i)
	case int8:
		startValue = float64(i)
	case uint64:
		startValue = float64(i)
	case uint32:
		startValue = float64(i)
	case uint16:
		startValue = float64(i)
	case uint8:
		startValue = float64(i)
	default:
		// startValue = float64(i)
		panic(fmt.Errorf("error converting")) // TODO
	}

	newValue := float64(startValue)/float64(startUnit.Scale) - float64(startUnit.Offset)
	multiplier := conversionMultiplier[startUnit.Unit][goalUnit]
	newValue *= multiplier
	// newValue = (newValue + float64(goalUnit.Offset)) * float64(goalUnit.Scale)
	return newValue
}

func ConvertEnum(value interface{}, startUnit Unit_e, goalUnit Unit_e) float64 {
	return Convert(value, 
		Unit{Unit: startUnit, Scale: 1, Offset: 0},
		goalUnit,
	 )
}

var (
{{- range $message := .Messages }}
{{- range $field := $message.Fields }}
	{{- $fieldUnit := $field.FieldUnit }}
	{{ toUpper $message.Name }}_{{ toUpper $field.Name }}_DEFAULT_UNIT = Unit{
		Unit: {{ unitToEnum $fieldUnit.Unit }}, 
		Scale: {{ $fieldUnit.Scale }}, 
		Offset: {{ $fieldUnit.Offset }}, 
	}	
{{- end }}
{{- end }}
)

`))

var typesTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"typeCorrection": typeToType,
}).Parse(`// Code generated; DO NOT EDIT

package generated

{{ range $type := .Types }}
{{- if (eq $type.BaseType "enum") }}
type {{ $type.Name }} uint8
const (
    {{- range $type_value := $type.Values}}
    {{ $type.Name }}_{{ $type_value.Name }} = {{ $type_value.Value }}
    {{- end }}
)
{{- else }}
type {{ $type.Name }} {{  typeCorrection $type.BaseType}}
{{- end }}
{{ if (eq $type.Name "mesg_num") }}
var validIDs = []uint{
{{- range $type_value := $type.Values }}
	{{ $type_value.Value }},
{{- end }}
}
func IsValidMessageID(messageID uint) bool {
	for _, id := range validIDs {
		if messageID == id {
			return true
		}
	}
	return false
}
{{ end }}
{{ end }}

`))

func jsonTag(fieldName string) string {
	return "`json:\"" + fieldName + ",omitempty\"`"
}

var packageTemplate = template.Must(template.New("").Funcs(template.FuncMap{
	"typeToType":            typeToType,
	"getTypeSize":           getTypeSize,
	"getUnderlyingGoType":   getUnderlyingGoType,
	"createDataMessageName": createDataMessageName,
	"createDataFieldName":   createDataFieldName,
	"jsonTag":               jsonTag,
}).Parse(`// Code generated; DO NOT EDIT.

package generated

import (
	"encoding/binary"
    "fmt"
)

type DataMessageData interface {
	Read(uint, []byte, binary.ByteOrder) error
}

const (
{{- range $message := .Messages }}
	{{ (createDataMessageName $message.Name )}}ID = {{ $message.MessageNum }}
{{- end }}
)

var MessageIDToString = map[uint]string {
{{- range $message := .Messages }}
{{ $message.MessageNum }}: "{{ $message.Name }}",
{{- end }}
}

func CreateMessageData(messageNum uint) (DataMessageData, error) {
    switch messageNum {
    {{- range $message := .Messages }}
    case {{ $message.MessageNum }}:
		return &{{ (createDataMessageName $message.Name )}}{}, nil
    {{- end }}
    default:
        return nil, fmt.Errorf("unable to create message data with num %d", messageNum)
    }   
}

{{ range $message := .Messages }}
type {{ (createDataMessageName $message.Name) }} struct {
    {{- range $field := $message.Fields}}
    {{ createDataFieldName $field.Name }} *{{ typeToType $field.FieldType.Name }} {{ (jsonTag $field.Name) }}
    {{- end}}
	{{/*
	{{- range $field := $message.Fields }}
	{{ $field.FieldUnit }}
	{{- end }}
	*/}}
}
{{/*
func (p *{{ createDataMessageName $message.Name }}) MarshalJSON() ([]byte, error) {
	fmt.Println("Marshalling")
	return nil, nil
}
*/}}
{{ end }}

{{ range $message := .Messages }}
func new{{ createDataMessageName $message.Name }}() *{{ createDataMessageName $message.Name }} {
	return &{{ createDataMessageName $message.Name }}{}
}
{{ end }}

{{ range $message := .Messages}}
func (data *{{ (createDataMessageName $message.Name) }}) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    {{- range $field := $message.Fields }}
    case {{ $field.FieldNumber }}:
		{{- $typeSize := (getTypeSize $field.FieldType.BaseType) }}
		{{- $goType :=(getUnderlyingGoType $field.FieldType.BaseType) }}
		data.{{ createDataFieldName $field.Name }} = new({{ typeToType $field.FieldType.Name }})
        {{- if (eq $typeSize 1) }}
        *data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(buffer[0])
        {{- end }}
		{{- if (eq $goType "uint16")}}
		*data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(binary.LittleEndian.Uint16(buffer))
		{{- end }}
		{{- if (eq $goType "uint32")}}
		*data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(binary.LittleEndian.Uint32(buffer))
		{{- end }}
		{{- if (eq $goType "uint64")}}
		*data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(binary.LittleEndian.Uint64(buffer))
		{{- end }}
		{{- if (eq $goType "int16")}}
		*data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(binary.LittleEndian.Uint16(buffer))
		{{- end }}
		{{- if (eq $goType "int32")}}
		*data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(binary.LittleEndian.Uint32(buffer))
		{{- end }}
		{{- if (eq $goType "int64")}}
		*data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(binary.LittleEndian.Uint64(buffer))
		{{- end }}
		{{- if (eq $goType "float32")}}
		*data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(binary.LittleEndian.Float32(buffer))
		{{- end }}
		{{- if (eq $goType "float64")}}
		*data.{{ createDataFieldName $field.Name }} = {{ (typeToType $field.FieldType.Name) }}(binary.LittleEndian.Float64(buffer))
		{{- end }}
    {{- end }}
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "{{ $message.Name }}")
    }
	return nil
}
{{ end }}

`))
