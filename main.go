package schema

import (
	"reflect"
	"strings"
)

type FieldType string

const (
	FieldString FieldType = "string"
	FieldInt    FieldType = "int"
	FieldBool   FieldType = "bool"
)

type Field struct {
	Key     string    `json:"key"`
	Type    FieldType `json:"type"`
	Default *string   `json:"default"`
	Enum    *[]string `json:"enum"`
}

func Marshal(obj interface{}) ([]byte, error) {
	panic("not implemented")
}

func Unmarshal(data []byte, obj interface{}) error {
	panic("not implemented")
}

func Describe(obj interface{}) ([]*Field, error) {
	schemaValue := reflect.ValueOf(obj)

	fields := make([]*Field, 0, schemaValue.NumField())

	for i := 0; i < schemaValue.NumField(); i++ {
		field := schemaValue.Field(i)
		fieldType := schemaValue.Type().Field(i)

		newField := &Field{
			Key:  fieldType.Tag.Get("key"),
			Type: FieldType(field.Kind().String()),
		}

		if defaultVal, ok := fieldType.Tag.Lookup("default"); ok {
			newField.Default = &defaultVal
		}

		if enumVal, ok := fieldType.Tag.Lookup("enum"); ok {
			enumVals := strings.Split(enumVal, ",")
			newField.Enum = &enumVals
		}

		fields = append(fields, newField)
	}

	return fields, nil
}
