package utils

import (
	"encoding/json"
	"log"
	"reflect"
)

func StructToJSON(s any) []byte {
	sType := reflect.TypeOf(s)
	if sType.Kind() != reflect.Struct {
		log.Print("argument is not a struct. Type: ")
		log.Fatal(sType.Kind())
	}

	raw, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("Error when marshaling a strct.\nStruct: %s\nError: %s", s, err)
	}

	return raw
}

func MapToJSON(s any) []byte {
	sType := reflect.TypeOf(s)
	if sType.Kind() != reflect.Map {
		log.Print("argument is not a struct. Type: ")
		log.Fatal(sType.Kind())
	}

	raw, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("Error when marshaling a strct.\nStruct: %s\nError: %s", s, err)
	}

	return raw
}
