package piper

import "reflect"

type indexableSource struct {
	indexable reflect.Value
	offset    int
}

func (s *indexableSource) Next() (values []reflect.Value, ok bool) {
	ok = s.offset < s.indexable.Len()

	if ok {
		values = []reflect.Value{s.indexable.Index(s.offset)}
		s.offset++
	}

	return
}

func FromIndexable(indexable interface{}) Pipe {
	return Pipe{Source: &indexableSource{indexable: reflect.ValueOf(indexable)}}
}
