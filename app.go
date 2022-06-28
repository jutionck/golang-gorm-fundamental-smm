package main

import (
	"enigmacamp.com/golang-gorm/config"
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/repository"
	generateid "github.com/jutionck/generate-id"
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
	repo := repository.NewCustomerRepository(db)

	//Insert
	customer := model.Customer{
		Id:      generateid.GenerateId(),
		Name:    "Rifqi Puasa",
		Address: "Depok",
		Phone:   "28299292",
		Email:   "rifqi.puasa@gmail.com",
		Balance: 10000,
	}
	err := repo.Create(&customer)
	if err != nil {
		log.Println(err.Error())
	}

	// Update
	//customerExisting := model.Customer{
	//	Id: "64ef0857-a08e-4a62-8eda-4e24c6aef326",
	//}
	//err := repo.Update(&customerExisting, model.Customer{
	//	Address: "",
	//	Balance: 150000,
	//})
	//if err != nil {
	//	log.Println(err.Error())
	//}

}
