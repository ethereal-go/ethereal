package ethereal

import (
	"flag"
	"log"
)

func CliRun() {
	database := flag.String("database", "migrate", "action database")
	seed := flag.String("seed", "up", "action seeder")
	flag.Parse()

	switch *database {
	case "migrate":
		App.Db.AutoMigrate(tables()...)
	case "rollback":
		App.Db.DropTable(tables()...)
	case "refresh":
		App.Db.DropTable(tables()...)
		App.Db.AutoMigrate(tables()...)
	default:
		log.Println(`This value is not set.`)
	}

	switch *seed {
	case "up":
		role := Role{Name: "User", DisplayName: "User", Description: "Simple user"}
		user := User{Email: "", Name: "", Password: "", Role: role}
		App.Db.Save(&user)
	}
}

/**
/ Get list tables
*/
func tables() []interface{} {
	return []interface{}{&User{}, &Role{}}
}
