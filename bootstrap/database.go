package bootstrap

import (
	"errors"
	"fmt"

	"github.com/vinoMamba/gohub/app/models/user"
	"github.com/vinoMamba/gohub/pkg/config"
	"github.com/vinoMamba/gohub/pkg/database"
	"github.com/vinoMamba/gohub/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	database.Connect(dbConfig, logger.NewGormLogger())
	database.DB.AutoMigrate(&user.User{})
}
