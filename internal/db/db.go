package db

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {

	db, err := gorm.Open(sqlite.Open("../../events.db"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to open DB: %v", err))
	}

	if gin.Mode() == gin.ReleaseMode {
		db.Logger.LogMode(0)
	}

	DB = db
	rawDB := RawDB()

	rawDB.SetMaxIdleConns(20)
	rawDB.SetMaxOpenConns(100)

	err = Migrate()
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate DB: %v", err))

	}
}

// RawDB returns the raw SQL database instance.
func RawDB() *sql.DB {
	db, err := DB.DB()
	if err != nil {
		panic(err)
	}

	return db
}
