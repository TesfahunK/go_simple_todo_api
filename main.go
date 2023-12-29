package main

import (
	"gorm.io/gorm"
	"simple_api/config"
)

var (
	Database *gorm.DB
)

func main() {
	setupDb()
}

func setupDb() error {

	var err error

	Database, err = config.OpendbConnection()
	println(err != nil)

	if err != nil {

		return err
	}

	return nil
}
