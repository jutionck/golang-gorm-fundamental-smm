package main

import (
	"enigmacamp.com/golang-gorm/config"
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/repository"
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
	//customer := model.Customer{
	//	Id:      generateid.GenerateId(),
	//	Name:    "Rifqi Puasa",
	//	Address: "Depok",
	//	Phone:   "28299292",
	//	Email:   "rifqi.puasa@gmail.com",
	//	Balance: 10000,
	//}
	//err := repo.Create(&customer)
	//if err != nil {
	//	log.Println(err.Error())
	//}

	// Update
	customerExisting := model.Customer{
		Id: "51999e87-9634-4a37-935f-99bf6851adf9",
	}
	err := repo.Update(&customerExisting, map[string]interface{}{
		"address":   "",
		"balance":   15000,
		"is_status": 0,
	})
	if err != nil {
		log.Println(err.Error())
	}

}
