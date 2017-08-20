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
		app.Db.Model(&User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	case "rollback":
		app.Db.DropTable(tables()...)
	case "refresh":
		app.Db.DropTable(tables()...)
		app.Db.AutoMigrate(tables()...)
		app.Db.Model(&User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	}
}

func tables() []interface{} {
	return []interface{}{&User{}, &Role{}}
}
