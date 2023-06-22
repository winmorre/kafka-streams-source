package models

import (
	"errors"
	"fmt"
	"reflect"
)

type Trade struct {
	Symbol    string   `json:"s"`
	Price     float64  `json:"p"`
	Timestamp float64  `json:"t"`
	Volume    float64  `json:"v"`
	C         []string `json:"c,omitempty"`
}

func SetField(obj interface{}, fieldName string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(fieldName)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", fieldName)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("provided value type didn't match obj field type")
	}
	structFieldValue.Set(val)
	return nil
}

func (t *Trade) ToStruct(value map[string]interface{}) error {
	for k, v := range value {
		switch k {
		case "s":
			err := SetField(t, "Symbol", v)
			if err != nil {
				return err
			}
		case "t":
			err := SetField(t, "Timestamp", v)
			if err != nil {
				return err
			}
		case "p":
			err := SetField(t, "Price", v)
			if err != nil {
				return err
			}
		case "v":
			err := SetField(t, "Volume", v)
			if err != nil {
				return err
			}
		}
	}
	println(t.Price)
	return nil
}
