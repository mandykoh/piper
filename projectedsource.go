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

	values = s.projection.Call(values)
	return
}
