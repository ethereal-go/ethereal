package ethereal

import (
	"flag"
)

func CliRun() {
	database := flag.String("database", "migrate", "action database")
	flag.Parse()

	switch *database {
	case "migrate":
		app.Db.AutoMigrate(tables()...)
	case "rollback":
		app.Db.DropTable(tables()...)
	case "refresh":
		app.Db.DropTable(tables()...)
		app.Db.AutoMigrate(tables()...)
	}
}

func tables() []interface{} {
	return []interface{}{&User{}, &Role{}}
}
