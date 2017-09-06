package ethereal

import (
	"reflect"
)

/**
 / Function add new value in context which there are in
 / App struct.
 / Depending on count parameters
 / One: key name struct
 / Two:key choose
 */
func AddContext(value ...interface{})  {
	switch reflect.ValueOf(value[1]).Kind() {
	case reflect.Struct:
		ctxStruct(value[0].(*Application), value[1])
	default:
		ctx(value[0].(*Application), value[1], value[2])
	}
}