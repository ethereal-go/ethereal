package ethereal

import (
	"reflect"
	"github.com/ethereal-go/ethereal/root/app"
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
		CtxStruct(value[0].(*app.Application), value[1])
	default:
		Ctx(value[0].(*app.Application), value[1], value[2])
	}
}