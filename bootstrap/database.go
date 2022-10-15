package bootstrap

import (
	"errors"
	"fmt"

	"github.com/vinoMamba/gohub/app/models/user"
	"github.com/vinoMamba/gohub/pkg/config"
	"github.com/vinoMamba/gohub/pkg/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("db.type") {
	case "postgres":
		dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
			config.Get("db.postgres.host"),
			config.Get("db.postgres.username"),
			config.Get("db.postgres.password"),
			config.Get("db.postgres.name"),
			config.Get("db.postgres.port"),
		)
		dbConfig = postgres.Open(dsn)
	default:
		panic(errors.New("database connection not supported"))
	}

	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	database.DB.AutoMigrate(&user.User{})
}
