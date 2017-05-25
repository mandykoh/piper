package piper

import "reflect"

type WrappedSource func() (values []reflect.Value, restOrEnd WrappedSource)
