package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(mariadb)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"password",
		"household-accounts",
	)
	// Get a database handle.
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected!")
	return db, nil
}
