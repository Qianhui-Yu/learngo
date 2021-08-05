package reflect

import "reflect"

var A int

func init() {
	A++
}

func SetErrNil(err *error) {
	*err = nil
}

func Indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}