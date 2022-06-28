package config

import (
	"enigmacamp.com/golang-gorm/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Config struct {
	db *gorm.DB
}

func (c *Config) DbClose() error {
	db, _ := c.db.DB()
	err := db.Close()
	if err != nil {
		panic(err)
	}
	return err
}

func (c *Config) initDb() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	env := os.Getenv("ENV")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	enigmaDb, err := db.DB()
	err = enigmaDb.Ping()
	if err != nil {
		panic(err)
	}
	if env == "dev" {
		c.db = db.Debug()
	} else if env == "migration" {
		c.db = db.Debug()
		err := c.db.AutoMigrate(&model.Customer{})

		if err != nil {
			return
		}
	} else {
		c.db = db
	}
}

func (c *Config) DbConn() *gorm.DB {
	return c.db
}

func NewConfig() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}
