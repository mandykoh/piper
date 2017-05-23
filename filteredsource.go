package piper

import "reflect"

type filteredSource struct {
	source Source
	filter reflect.Value
}

func (s filteredSource) Next() (values []reflect.Value, ok bool) {

	for {
		values, ok = s.source.Next()
		if !ok {
			return
		}

		filterResult := s.filter.Call(values)
		if filterResult[0].Bool() {
			return
		}
	}
}
