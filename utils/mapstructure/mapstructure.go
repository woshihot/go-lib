package mapstructure

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
)

/*
	m map -> s struct
	m arr -> s []struct
	m string -> s string
*/
func MapToStructure(m interface{}, sPointer interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           sPointer,
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	if err = decoder.Decode(m); err != nil {
		return err
	}
	return nil
}

func StructureToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
