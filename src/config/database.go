package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest-project/src/adapters"
)

func DBConnection(pgConfig PostgresDB) *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.User,
		pgConfig.DBName,
		pgConfig.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&adapters.Article{})
	if err != nil {
		panic(err)
	}
	return db
}
