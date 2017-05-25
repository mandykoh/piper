package piper

import "reflect"

type single struct {
	values []interface{}
	done   bool
}

func (s *single) Source() ([]reflect.Value, Source) {
	if s.done {
		return nil, nil
	}

	values := make([]reflect.Value, len(s.values))
	for i, v := range s.values {
		values[i] = reflect.ValueOf(v)
	}

	s.done = true

	return values, s.Source
}

func FromSingle(itemValues ...interface{}) Pipe {
	s := &single{values: itemValues}
	return From(s.Source)
}
