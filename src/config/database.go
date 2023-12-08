package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rest-project/src/adapters/models"
)

type SPostgres struct {
	*gorm.DB
}

func DBConnection(pgConfig PostgresDB) *SPostgres {
	sPostgres := new(SPostgres)
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		pgConfig.Host,
		pgConfig.Port,
		pgConfig.User,
		pgConfig.DBName,
		pgConfig.Password,
	)

	sPostgres.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Println("postgres connection opened")

	if err = sPostgres.setup(); err != nil {
		log.Fatalln("error in setup postgres ", err)
	}

	return sPostgres
}

func (r *SPostgres) setup() error {
	return r.DB.AutoMigrate(
		new(models.Article),
	)
}

func (r *SPostgres) Close() error {
	db, err := r.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (r *SPostgres) Transaction(fc func(*SPostgres) error) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		r.DB = tx
		return fc(r)
	})
}
