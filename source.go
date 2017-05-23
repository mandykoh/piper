package piper

import "reflect"

type Source interface {
	Next() (values []reflect.Value, ok bool)
}
