package piper

import "reflect"

type Pipe struct {
	Source WrappedSource
}

func (p Pipe) Aggregate(seed, aggregator interface{}) Pipe {
	a := &aggregation{source: p.Source, seed: reflect.ValueOf(seed), aggregator: reflect.ValueOf(aggregator)}
	return Pipe{Source: a.Source}
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

func From(sourceFunc interface{}) Pipe {
	return Pipe{Source: wrapRawSource(sourceFunc)}
}

func FromMany(items interface{}) Pipe {
	m := &many{items: reflect.ValueOf(items)}
	return fromSource(m.Source)
}

func FromSingle(itemValues ...interface{}) Pipe {
	s := &single{values: itemValues}
	return fromSource(s.Source)
}

func fromSource(s WrappedSource) Pipe {
	return Pipe{Source: s}
}

func wrapRawSource(rs interface{}) (newSource WrappedSource) {
	sourceFunc := reflect.ValueOf(rs)

	newSource = func() ([]reflect.Value, WrappedSource) {
		values := sourceFunc.Call([]reflect.Value{})
		hasValue := false

		if len(values) > 0 {
			sourceFunc = values[len(values)-1]
			if sourceFunc.Kind() == reflect.Func && !sourceFunc.IsNil() {
				values = values[:len(values)-1]
				hasValue = true
			}
		}

		if !hasValue {
			return nil, nil
		}

		return values, newSource
	}

	return
}
