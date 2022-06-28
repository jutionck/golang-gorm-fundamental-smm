package main

import (
	"enigmacamp.com/golang-gorm/config"
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/repository"
	"fmt"
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
	//	Name:    "Jamal Udin",
	//	Address: "Surabya",
	//	Phone:   "2929022",
	//	Email:   "jamal.udin@gmail.com",
	//	Balance: 10000,
	//}
	//err := repo.Create(&customer)
	//if err != nil {
	//	log.Println(err.Error())
	//}

	//customerExisting := model.Customer{
	//	Id: "0454ad0e-e6f2-4566-a6f5-6cafb8b02e26",
	//}
	//err := repo.Update(&customerExisting, map[string]interface{}{
	//	"address":   "",
	//	"balance":   15000,
	//	"is_status": 0,
	//})
	//if err != nil {
	//	log.Println(err.Error())
	//}

	// Delete
	//err := repo.Delete(&customerExisting)
	//if err != nil {
	//	log.Println(err.Error())
	//}

	// Find By Id
	//customerExisting, err := repo.FindById(customerExisting.Id)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Println(customerExisting)

	// FindByAllBy
	customers := []model.Customer{}
	customers, err := repo.FindAllBy(map[string]interface{}{
		"address": "Depok",
	})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("FindByAllBy: ", customers)

	// FindFirstBy
	customer := model.Customer{}
	customer, err = repo.FindFirstBy(map[string]interface{}{
		"address": "Depok",
	})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("FindFirstBy: ", customer)

	// FindBy
	customers01 := []model.Customer{}
	customers01, err = repo.FindBy("name LIKE ? AND is_status = ?", "%J%", 1)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("FindBy: ", customers01)
}
