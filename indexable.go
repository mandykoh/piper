package piper

import "reflect"

type indexable struct {
	items  reflect.Value
	offset int
}

func (i *indexable) Source() ([]reflect.Value, Source) {
	var values []reflect.Value

	if i.offset < i.items.Len() {
		values = []reflect.Value{i.items.Index(i.offset)}
		i.offset++
		return values, i.Source
	}

	return nil, nil
}

func FromIndexable(items interface{}) Pipe {
	i := &indexable{items: reflect.ValueOf(items)}
	return Pipe{Source: i.Source}
}
