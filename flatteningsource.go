package piper

import "reflect"

type flatteningSource struct {
	source             Source
	currentValues      []reflect.Value
	indexables         []reflect.Value
	indexablePositions []int
	indexes            []int
}

func (s *flatteningSource) Next() ([]reflect.Value, bool) {
	for {
		values, ok := s.nextValuesFromIndexables()
		if ok {
			return values, true
		}

		newValues, ok := s.source.Next()
		if !ok {
			return nil, false
		}

		s.resetIndexables(newValues)
	}
}

func (s *flatteningSource) nextValuesFromIndexables() (values []reflect.Value, ok bool) {

	if s.indexables == nil {
		return nil, false
	}

	for i := 0; i < len(s.indexables); i++ {
		value := s.indexables[i].Index(s.indexes[i])
		s.currentValues[s.indexablePositions[i]] = value
	}

	for i := len(s.indexables) - 1; ; {
		s.indexes[i]++
		if s.indexes[i] >= s.indexables[i].Len() {
			s.indexes[i] = 0
			i--
			if i < 0 {
				s.indexables = nil
				break
			}
		} else {
			break
		}
	}

	return s.currentValues, true
}

func (s *flatteningSource) resetIndexables(values []reflect.Value) {
	s.currentValues = values

	s.indexables = []reflect.Value{}
	s.indexablePositions = []int{}

	for i := 0; i < len(values); i++ {
		switch values[i].Kind() {
		case reflect.Array, reflect.Slice:
			s.indexables = append(s.indexables, values[i])
			s.indexablePositions = append(s.indexablePositions, i)
		}
	}

	s.indexes = make([]int, len(s.indexables))
}
