package ethereal

import (
	"flag"
)

func CliRun() {
	database := flag.String("database", "migrate", "action database")
	flag.Parse()

	switch *database {
	case "migrate":
		app.Db.AutoMigrate(&Todo{})
	case "rollback":
		app.Db.DropTable(&Todo{})
	}
}
