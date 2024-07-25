package migrate

import (
	"pgd-server.com/config"
	"pgd-server.com/src/entities"
)

func MigrateDB() {
	config.DB.AutoMigrate(
		&entities.User{},
		&entities.Customer{},
	)
}
