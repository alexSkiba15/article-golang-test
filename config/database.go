package config

import (
	"fmt"
	"log"

	"github.com/alexSkiba15/article-golang-test/adapters/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	sPostgres.DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

func (r *SPostgres) Close() {
	db, err := r.DB.DB()
	if err != nil {
		fmt.Printf("error in closing db %s", err)
	}
	err = db.Close()
	if err != nil {
		fmt.Printf("error in closing db %s", err)
	}
}

func (r *SPostgres) Transaction(fc func(*SPostgres) error) (err error) {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		r.DB = tx
		return fc(r)
	})
}
