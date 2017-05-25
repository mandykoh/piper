package piper

import "reflect"

type Source func() (values []reflect.Value, restOrEnd Source)
