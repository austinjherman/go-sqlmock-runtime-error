package main

import (
	"log"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type base struct {
	ID        uuid.UUID      `json:"-" gorm:"primaryKey;type:uuid;not null"`
  CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
  UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
  DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type user struct {
  base
  Name     string `json:"-"`
  Email    string `json:"-" gorm:"unique_index:user_email_index"`
  Password string `json:"-" gorm:"size:72"`
}

func newMockDatabase() (*gorm.DB, sqlmock.Sqlmock) {

	// get db and mock
	sqlDB, mock, err := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp),
	)
	if err != nil {
		log.Fatalf("[sqlmock new] %s", err)
	}
	defer sqlDB.Close()

	// create dialector
	dialector := mysql.New(mysql.Config{
		Conn: sqlDB,
		DriverName: "mysql",
		SkipInitializeWithVersion: true,
	})

	// open the database
	db, err := gorm.Open(dialector, &gorm.Config{ PrepareStmt: true })
	if err != nil {
		log.Fatalf("[gorm open] %s", err)
	}

	return db, mock
}

func initDB(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&user{})
}