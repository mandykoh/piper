package piper

import "reflect"

type many struct {
	items  reflect.Value
	offset int
}

func (m *many) Source() ([]reflect.Value, Source) {
	var values []reflect.Value

	if m.offset < m.items.Len() {
		values = []reflect.Value{m.items.Index(m.offset)}
		m.offset++
		return values, m.Source
	}

	return nil, nil
}

func FromMany(items interface{}) Pipe {
	m := &many{items: reflect.ValueOf(items)}
	return From(m.Source)
}
