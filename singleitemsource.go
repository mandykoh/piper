package piper

import "reflect"

type singleItemSource struct {
	values []interface{}
	done   bool
}

func (s *singleItemSource) Next() (values []reflect.Value, ok bool) {
	if s.done {
		return nil, false
	}

	values = make([]reflect.Value, len(s.values))
	for i, v := range s.values {
		values[i] = reflect.ValueOf(v)
	}

	s.done = true

	return values, true
}

func From(itemValues ...interface{}) Pipe {
	return Pipe{Source: &singleItemSource{values: itemValues}}
}
