package piper

import "reflect"

type many struct {
	items  reflect.Value
	offset int
}

func (m *many) Source() ([]reflect.Value, WrappedSource) {
	var values []reflect.Value

	if m.offset < m.items.Len() {
		values = []reflect.Value{m.items.Index(m.offset)}
		m.offset++
		return values, m.Source
	}

	return nil, nil
}
