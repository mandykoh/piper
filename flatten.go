package piper

import "reflect"

type flatten struct {
	source             WrappedSource
	currentValues      []reflect.Value
	indexables         []reflect.Value
	indexablePositions []int
	indexes            []int
}

func (f *flatten) Source() ([]reflect.Value, WrappedSource) {
	for {
		values, ok := f.nextValuesFromIndexables()
		if ok {
			return values, f.Source
		}

		var newValues []reflect.Value
		newValues, f.source = f.source()
		if f.source == nil {
			return nil, nil
		}

		f.resetIndexables(newValues)
	}
}

func (f *flatten) incrementIndexables() {

	if len(f.indexables) > 0 {
		for i := len(f.indexables) - 1; ; {
			f.indexes[i]++
			if f.indexes[i] >= f.indexables[i].Len() {
				f.indexes[i] = 0
				i--
				if i < 0 {
					f.currentValues = nil
					break
				}
			} else {
				break
			}
		}
	} else {
		f.currentValues = nil
	}
}

func (f *flatten) nextValuesFromIndexables() ([]reflect.Value, bool) {

	if f.currentValues == nil {
		return nil, false
	}

	values := f.currentValues

	for i := 0; i < len(f.indexables); i++ {
		value := f.indexables[i].Index(f.indexes[i])
		values[f.indexablePositions[i]] = value
	}

	f.incrementIndexables()

	return values, true
}

func (f *flatten) resetIndexables(values []reflect.Value) {
	f.currentValues = values

	f.indexables = nil
	f.indexablePositions = nil

	for i := 0; i < len(values); i++ {
		values[i] = convertToRuntimeType(values[i])

		switch values[i].Kind() {
		case reflect.Array, reflect.Slice:
			f.indexables = append(f.indexables, values[i])
			f.indexablePositions = append(f.indexablePositions, i)
		}
	}

	f.indexes = make([]int, len(f.indexables))
}
