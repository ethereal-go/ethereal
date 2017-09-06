package ethereal

import (
	"github.com/spf13/viper"
	"strings"
	"os"
	"context"
	"reflect"
)

// Here functions helpers
// ----------------------------------

/**
 / Get configuration value
 */
func config(name string, byDefault ...interface{}) interface{} {
	var temp string
	if temp = os.Getenv(name); temp == "" {
		if temp = viper.GetString(strings.ToLower(name)); temp == ""{
			viper.SetDefault(name, byDefault)
		}
	}
	return temp
}

/**
 / Add value in Context structure
 */
func ctxStruct(app *Application, value interface{})  {
	app.Context = context.WithValue(App.Context, getType(value), value)
}

func ctx(app *Application, key interface{}, value interface{})  {
	app.Context = context.WithValue(App.Context, key, value)
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