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

func (s *flatteningSource) incrementIndexables() {

	if len(s.indexables) > 0 {
		for i := len(s.indexables) - 1; ; {
			s.indexes[i]++
			if s.indexes[i] >= s.indexables[i].Len() {
				s.indexes[i] = 0
				i--
				if i < 0 {
					s.currentValues = nil
					break
				}
			} else {
				break
			}
		}
	} else {
		s.currentValues = nil
	}
}

func (s *flatteningSource) nextValuesFromIndexables() ([]reflect.Value, bool) {

	if s.currentValues == nil {
		return nil, false
	}

	values := s.currentValues

	for i := 0; i < len(s.indexables); i++ {
		value := s.indexables[i].Index(s.indexes[i])
		values[s.indexablePositions[i]] = value
	}

	s.incrementIndexables()

	return values, true
}

func (s *flatteningSource) resetIndexables(values []reflect.Value) {
	s.currentValues = values

	s.indexables = nil
	s.indexablePositions = nil

	for i := 0; i < len(values); i++ {
		values[i] = reflect.ValueOf(values[i].Interface())

		switch values[i].Kind() {
		case reflect.Array, reflect.Slice:
			s.indexables = append(s.indexables, values[i])
			s.indexablePositions = append(s.indexablePositions, i)
		}
	}

	s.indexes = make([]int, len(s.indexables))
}
