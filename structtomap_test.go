package structtomap

import (
	"reflect"
	"testing"
)

type TestStruct struct {
	Name1 int
	Name2 string
	Name3 string
}

func TestStructToMap(t *testing.T) {
	testStruct := TestStruct{1, "test", "more of a test"}

	_, err := StructToMap("notastruct")
	if err == nil {
		t.Error("no error given when passed a string")
	}

	mapResult, err := StructToMap(testStruct)
	if err != nil {
		t.Error("error returned on good struct")
	}

	if mapResult["Name1"] != 1 || mapResult["Name2"] != "test" {
		t.Error("resulting map does not contain the correct data", mapResult)
	}
}

func TestGetStructType(t *testing.T) {
	testStruct1 := TestStruct{1, "test", "more of a test"}
	sType1 := getStructType(testStruct1)
	if sType1.Kind() != reflect.Struct {
		t.Error("struct not returned when passed struct")
	}

	testStruct2 := &TestStruct{1, "test", "more of a test"}
	sType2 := getStructType(testStruct2)
	if sType2.Kind() != reflect.Struct {
		t.Error("struct not returned when passed pointer")
	}
}
