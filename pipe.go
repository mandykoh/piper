package piper

import "reflect"

type Pipe struct {
	Source Source
}

func (p Pipe) Select(projection interface{}) Pipe {
	return Pipe{Source: projectedSource{source: p.Source, projection: reflect.ValueOf(projection)}}
}

func (p Pipe) To(sink interface{}) {
	sinkFunc := reflect.ValueOf(sink)

	for {
		values, ok := p.Source.Next()
		if !ok {
			break
		}

		sinkFunc.Call(values)
	}
}

func (p Pipe) Where(filter interface{}) Pipe {
	return Pipe{Source: filteredSource{source: p.Source, filter: reflect.ValueOf(filter)}}
}
