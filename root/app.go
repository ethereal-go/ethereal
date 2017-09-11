package root

import (
	"github.com/jinzhu/gorm"
	"github.com/graphql-go/graphql"
	"context"
	"github.com/qor/i18n"
	"github.com/ethereal-go/ethereal/root/middleware"
	"github.com/ethereal-go/ethereal/root/config"
)

// Base structure
type Application struct {
	Db              *gorm.DB
	I18n            *i18n.I18n
	Middleware      *middleware.Middleware
	GraphQlMutation graphql.Fields
	GraphQlQuery    graphql.Fields
	Context         context.Context
	Config          *config.Config
}
