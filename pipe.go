package piper

import "reflect"

type Pipe struct {
	Source Source
}

func (p Pipe) Flatten() Pipe {
	return Pipe{Source: &flatteningSource{source: p.Source}}
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

		for i := 0; i < len(values); i++ {
			values[i] = reflect.ValueOf(values[i].Interface())
		}
		sinkFunc.Call(values)
	}
}

func (p Pipe) Where(filter interface{}) Pipe {
	return Pipe{Source: filteredSource{source: p.Source, filter: reflect.ValueOf(filter)}}
}
