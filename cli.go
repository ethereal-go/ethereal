package ethereal

import (
	"flag"
)

func CliRun() {
	database := flag.String("database", "migrate", "action database")
	flag.Parse()

	switch *database {
	case "migrate":
		app.Db.AutoMigrate(&User{}, &Role{})
	case "rollback":
		app.Db.DropTable(&User{}, &Role{})
	case "refresh":
		app.Db.DropTable(&User{}, &Role{})
		app.Db.AutoMigrate(&User{}, &Role{})

	}
}
