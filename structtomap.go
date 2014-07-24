//Package structtomap is a small library for converting a struct to a map for easy iteration with for/range.
package structtomap

import (
	"errors"
	"reflect"
)

//Convert converts a struct to a map for easy iteration with for range.
//`struc` can be a pointer or a concrete struct.
// error will be nil if everything worked.
func Convert(struc interface{}) (map[string]interface{}, error) {

	returnMap := make(map[string]interface{})

	sType := getStructType(struc)

	if sType.Kind() != reflect.Struct {
		return returnMap, errors.New("variable given is not a struct or a pointer to a struct")
	}

	for i := 0; i < sType.NumField(); i++ {
		structFieldName := sType.Field(i).Name
		structVal := reflect.ValueOf(struc)
		returnMap[structFieldName] = structVal.FieldByName(structFieldName).Interface()
	}

	return returnMap, nil
}

func getStructType(struc interface{}) reflect.Type {
	sType := reflect.TypeOf(struc)
	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
	}

	return sType
}
