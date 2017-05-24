package piper

import "reflect"

type flatteningSource struct {
	sourceStack []Source
}

func (s *flatteningSource) Next() (values []reflect.Value, ok bool) {

	for {
		currentSourceIndex := len(s.sourceStack) - 1
		if currentSourceIndex < 0 {
			return nil, false
		}

		currentSource := s.sourceStack[currentSourceIndex]
		values, ok = currentSource.Next()
		if ok {

			if len(values) == 1 {
				subSource, ok := values[0].Interface().(Source)
				if ok {
					s.sourceStack = append(s.sourceStack, subSource)
					continue
				}
			}

			return
		}

		s.sourceStack = s.sourceStack[:currentSourceIndex]
	}
}
