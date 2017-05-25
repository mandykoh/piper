package piper

import "reflect"

type Pipe struct {
	Source Source
}

func (p Pipe) Flatten() Pipe {
	return Pipe{Source: &flatteningSource{source: p.Source}}
}

func (p Pipe) Select(projections ...interface{}) Pipe {
	projectionFuncs := make([]reflect.Value, len(projections))
	for i, projection := range projections {
		projectionFuncs[i] = reflect.ValueOf(projection)
	}

	return Pipe{Source: projectedSource{source: p.Source, projections: projectionFuncs}}
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
