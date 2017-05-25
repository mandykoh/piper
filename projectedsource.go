package piper

import "reflect"

type projectedSource struct {
	source      Source
	projections []reflect.Value
}

func (s projectedSource) Next() ([]reflect.Value, bool) {

	values, ok := s.source.Next()
	if !ok {
		return nil, false
	}

	convertToRuntimeTypes(values)

	var results []reflect.Value
	for _, p := range s.projections {
		results = append(results, p.Call(values)...)
	}

	return results, true
}
