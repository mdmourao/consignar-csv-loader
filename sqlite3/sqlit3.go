package sqlite3

import (
	"github.com/mdmourao/consignar-csv-loader/models"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	db, err := gorm.Open(sqlite.Open("consignar.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DBConn = db
}

func Disconnect() {
	sqlDB, _ := DBConn.DB()
	sqlDB.Close()
}

func Migrate() error {
	return DBConn.AutoMigrate(&models.EntityDb{})
}

func CreateEntity(entity models.EntityDb) error {
	return DBConn.Create(&entity).Error
}
