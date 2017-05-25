package piper

import "reflect"

func From(s Source) Pipe {
	return Pipe{Source: s}
}

func convertToRuntimeType(value reflect.Value) reflect.Value {

	if value.Kind() == reflect.Interface {
		trueValue := reflect.ValueOf(value.Interface())
		if trueValue.IsValid() {
			return trueValue
		}
	}

	return value
}

func convertToRuntimeTypes(values []reflect.Value) {
	for i := 0; i < len(values); i++ {
		values[i] = convertToRuntimeType(values[i])
	}
}
