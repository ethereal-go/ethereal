package migrations

import (
	"github.com/agoalofalife/ethereal"
)

func Run() {
	db := ethereal.Database()
	db.AutoMigrate(&ethereal.Todo{})
}
