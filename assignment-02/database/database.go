package database

import (
	"all-assignment-scalable-web-service-with-go/assignment-02/models"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "postgres"
	PASS     = "Indra19"
	DB_NAME  = "assignment2_sesi8"
	APP_PORT = ":8888"
)

func ConnectDB() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode= disable",
		HOST, PORT, USER, PASS, DB_NAME)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	log.Default().Println("connected to database")

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	// best practice, limit database connections or called connections pool
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxIdleTime(10 * time.Second)
	db.DB().SetConnMaxLifetime(10 * time.Second)

	db.AutoMigrate(models.Item{})
	db.AutoMigrate(models.Order{})
	return db
}
