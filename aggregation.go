package piper

import "reflect"

type aggregation struct {
	source     WrappedSource
	done       bool
	seed       reflect.Value
	aggregator reflect.Value
}

func (a aggregation) Source() ([]reflect.Value, WrappedSource) {
	aggregateResult := []reflect.Value{a.seed}

	if a.done {
		return nil, nil
	}

	for {
		var values []reflect.Value

		values, a.source = a.source()
		if a.source == nil {
			a.done = true
			return aggregateResult, a.Source
		}

		convertToRuntimeTypes(values)

		values = append(aggregateResult, values...)
		aggregateResult = a.aggregator.Call(values)
	}
}
