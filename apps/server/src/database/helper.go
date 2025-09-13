package database

import (
	"fmt"
	"strconv"

	"sodnix/apps/server/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT, config.TIME_ZONE,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("DB connection failed: " + err.Error())
	}

	if DB == nil {
		panic("DB is not initialized")
	}

	return DB
}

// PreloadWithFields preloads a relation but selects only the specified fields
func PreloadWithFields(relation string, fields []string) *gorm.DB {
	return DB.Preload(relation, func(tx *gorm.DB) *gorm.DB {
		return tx.Select(fields)
	})
}

func PreloadRelations(relations map[string][]string) *gorm.DB {
	for rel, fields := range relations {
		DB = PreloadWithFields(rel, fields)
	}
	return DB
}

func Migrate(models ...interface{}) error {
	if ok, _ := strconv.ParseBool(config.AUTO_MIGRATE); ok {
		return DB.AutoMigrate(models...)
	}
	return nil
}
