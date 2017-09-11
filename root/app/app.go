package app

import (
	"context"
	"github.com/ethereal-go/ethereal/root/config"
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n"
)


// Base structure
type Application struct {
	Db              *gorm.DB
	I18n            *i18n.I18n
	//Middleware      *middleware.Middleware // whether this dependence upon server startup?
	GraphQlMutation graphql.Fields
	GraphQlQuery    graphql.Fields
	Context         context.Context
	Config          *config.Config
}