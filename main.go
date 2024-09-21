package schema

type FieldType string

const (
	FieldString FieldType = "string"
	FieldInt    FieldType = "int"
	FieldBool   FieldType = "bool"
)

type Field struct {
	Key     string
	Type    FieldType
	Default *string
	Enum    *[]string
}
