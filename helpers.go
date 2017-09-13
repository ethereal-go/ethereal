package ethereal

import (
	"context"
	"github.com/ethereal-go/ethereal/root/app"
	"reflect"
)

// Here functions helpers
// ----------------------------------

///**
/// Get configuration value
//*/
//
//func GetCnf(name string, byDefault ...interface{}) interface{} {
//	var temp string
//	if temp = os.Getenv(name); temp == "" {
//		if temp = viper.GetString(strings.ToLower(name)); temp == "" {
//			viper.SetDefault(name, byDefault)
//		}
//	}
//	return temp
//
//}

/**
/ Add value in Context structure
*/
func CtxStruct(app *app.Application, value interface{}) context.Context {
	app.Context = context.WithValue(app.Context, getType(value), value)
	return app.Context
}

func Ctx(app *app.Application, key interface{}, value interface{}) context.Context {
	app.Context = context.WithValue(App.Context, key, value)
	return app.Context
}

/**
/ Get type
*/
func getType(unknown interface{}) string {
	if t := reflect.TypeOf(unknown); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

//func BasePathClient() string {
//	workPath, err := os.Getwd()
//	if err != nil {
//		panic(err)
//	}
//	return workPath
//}