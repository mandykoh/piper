package piper

import "reflect"

type filter struct {
	source WrappedSource
	test   reflect.Value
}

func (f filter) Source() ([]reflect.Value, WrappedSource) {

	for {
		var values []reflect.Value

		values, f.source = f.source()
		if f.source == nil {
			return nil, nil
		}

		convertToRuntimeTypes(values)

		filterResult := f.test.Call(values)
		if filterResult[0].Bool() {
			return values, f.Source
		}
	}
}
