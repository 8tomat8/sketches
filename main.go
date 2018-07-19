package main

import (
	"encoding/json"
	"os"
	"reflect"
)

type T struct {
	y [3]int16
}

type str struct {
	A int
	B string
	C struct {
		F float64
	}
}

func main() {
	enc := json.NewEncoder(os.Stdout)

	enc.Encode(f(([]int)(nil)))
	enc.Encode(f((*int)(nil)))
	enc.Encode(f(str{
		1,
		"asd",
		struct {
			F float64
		}{F: 123123.123123},
	}))
	var s *str
	s = nil
	enc.Encode(f(s))

	var m map[string]int
	enc.Encode(f(m))
	enc.Encode(f("asdasd"))

}

func f(d interface{}) interface{} {
	switch d.(type) {
	case bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex64, complex128, string:
		return d
	}
	rVal := reflect.ValueOf(d)

	if rVal.Kind() == reflect.Struct {
		return d
	}

	if rVal.IsNil() {
		switch rVal.Kind() {
		case reflect.Slice:
			return []struct{}{}
		default:
			return struct{}{}
		}
	}
	return d
}
