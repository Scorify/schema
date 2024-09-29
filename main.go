package schema

import (
	"encoding/json"
	"fmt"
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
	schemaValues := reflect.ValueOf(obj)

	fields := make(map[string]interface{}, schemaValues.NumField())

	for i := 0; i < schemaValues.NumField(); i++ {
		field := schemaValues.Field(i)
		fieldType := schemaValues.Type().Field(i)

		fields[fieldType.Tag.Get("key")] = field.Interface()
	}

	return json.Marshal(fields)
}

func Unmarshal(data []byte, obj interface{}) error {
	schemaValues := reflect.ValueOf(obj).Elem()

	var fields map[string]interface{}
	err := json.Unmarshal(data, &fields)
	if err != nil {
		return err
	}

	for i := 0; i < schemaValues.NumField(); i++ {
		field := schemaValues.Field(i)
		fieldType := schemaValues.Type().Field(i)

		val, ok := fields[fieldType.Tag.Get("key")]
		if ok {
			switch field.Kind() {
			case reflect.String:
				field.SetString(val.(string))
			case reflect.Int:
				field.SetInt(int64(val.(float64)))
			case reflect.Bool:
				field.SetBool(val.(bool))
			default:
				return fmt.Errorf("unsupported type: %s", field.Kind().String())
			}
		} else {
			return fmt.Errorf("missing key: %s", fieldType.Tag.Get("key"))
		}
	}

	return nil
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
