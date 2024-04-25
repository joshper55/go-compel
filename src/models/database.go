package models

import (
	"compel/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func ConnectToDB() {
	host := config.CNF.Postgres.Host
	port := config.CNF.Postgres.Port
	user := config.CNF.Postgres.User
	password := config.CNF.Postgres.Password
	database := config.CNF.Postgres.Database
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable Timezone=%v", host, user, password, database, port, "UTC")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

}
func Migrate() {
	DB.AutoMigrate(
		&CatInstruments{},
	)
}
