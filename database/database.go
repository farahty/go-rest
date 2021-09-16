package database

import (
	"fmt"

	"github.com/nimerfarahty/go-rest/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Conn *gorm.DB
)

func Connect() error {
	var err error
	dsn := buildDSN()
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	println("database connection initialed successfully.")
	return nil
}

func buildDSN() string {

	conf := config.Keys.Database

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Gaza",
		conf.Host,
		conf.User,
		conf.Password,
		conf.Database,
		conf.Port)
}
