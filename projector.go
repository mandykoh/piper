package piper

import "reflect"

type projector struct {
	source      Source
	projections []reflect.Value
}

func (p projector) Source() ([]reflect.Value, Source) {
	var values []reflect.Value

	values, p.source = p.source()
	if p.source == nil {
		return nil, nil
	}

	convertToRuntimeTypes(values)

	var results []reflect.Value
	for _, projection := range p.projections {
		results = append(results, projection.Call(values)...)
	}

	return results, p.Source
}
