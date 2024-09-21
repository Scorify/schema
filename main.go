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

func Marshal(obj interface{}) ([]byte, error) {
	panic("not implemented")
}

func Unmarshal(data []byte, obj interface{}) error {
	panic("not implemented")
}

func Describe(obj interface{}) ([]*Field, error) {
	panic("not implemented")
}
