package main

import (
	"enigmacamp.com/golang-gorm/config"
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/repository"
	"github.com/jutionck/generate-id"
	"log"
)

func main() {
	cfg := config.NewConfig()
	db := cfg.DbConn()
	defer func(cfg *config.Config) {
		err := cfg.DbClose()
		if err != nil {
			log.Println(err.Error())
		}
	}(&cfg)
	//err := db.AutoMigrate(&model.Customer{})
	//if err != nil {
	//	return
	//}

	repo := repository.NewCustomerRepository(db)

	// Insert
	customer := model.Customer{
		Id:      generateid.GenerateId(),
		Name:    "Bulan Bintang",
		Phone:   "829202002",
		Email:   "bulan.bintang@gmail.com",
		Balance: 10000,
	}
	err := repo.Create(&customer)
	if err != nil {
		log.Println(err.Error())
	}
}
