package ethereal

import (
	"github.com/graphql-go/graphql"
	"fmt"
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
	switch reflect.ValueOf(value[0]).Kind() {
	case reflect.Struct:
		ctxStruct(&App, value[0])
	default:
		ctx(&App, value[0], value[1])
	}
}

func contextBootstrapping()  {
	AddContext(JwtTokenRule{
		exclude: []*graphql.Object{
		  usersType,
		}})
}