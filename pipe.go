package piper

import "reflect"

type Pipe struct {
	Source Source
}

func (p Pipe) Flatten() Pipe {
	f := &flatten{source: p.Source}
	return Pipe{Source: f.Source}
}

func (p Pipe) Select(projections ...interface{}) Pipe {
	projectionFuncs := make([]reflect.Value, len(projections))
	for i, projection := range projections {
		projectionFuncs[i] = reflect.ValueOf(projection)
	}

	proj := projector{source: p.Source, projections: projectionFuncs}
	return Pipe{Source: proj.Source}
}

func (p Pipe) To(sink interface{}) {
	sinkFunc := reflect.ValueOf(sink)

	for {
		var values []reflect.Value

		values, p.Source = p.Source()
		if p.Source == nil {
			break
		}

		convertToRuntimeTypes(values)
		sinkFunc.Call(values)
	}
}

func (p Pipe) Where(predicate interface{}) Pipe {
	f := filter{source: p.Source, test: reflect.ValueOf(predicate)}
	return Pipe{Source: f.Source}
}
