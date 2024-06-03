package database

import (

	"biblioapp/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

	

    DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Loan{})
	

	cols, _ := DB.Migrator().ColumnTypes("loans")
	

for _, col := range cols {
    fmt.Println("Column name:", col.Name())
    fmt.Println("Column type:", col.DatabaseTypeName())
}
	
}
