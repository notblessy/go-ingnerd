package main

import (
	"github.com/notblessy/go-ingnerd/src/config"
	"github.com/notblessy/go-ingnerd/src/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.DbConnection()
)

func main() {
	defer config.KillDbConnection(db)

	routes.Routes()
}
