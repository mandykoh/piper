package piper

import "reflect"

type projectedSource struct {
	source     Source
	projection reflect.Value
}

func (s projectedSource) Next() (values []reflect.Value, ok bool) {

	values, ok = s.source.Next()
	if !ok {
		return
	}

	for i := 0; i < len(values); i++ {
		values[i] = reflect.ValueOf(values[i].Interface())
	}

	values = s.projection.Call(values)
	return
}
