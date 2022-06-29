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
	productRepo := repository.NewProductRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	//
	//product01 := []model.Product{
	//	{
	//		ProductName: "Kacang Asin",
	//	},
	//	{
	//		ProductName: "Keripik Mangga",
	//	},
	//	{
	//		ProductName: "Keripik Pisang",
	//	},
	//}
	//err := productRepo.Create(&product01[2])
	//IsError(err)

	//passwordHash, _ := utils.HashPassword("password")
	//customer01 := model.Customer{
	//	Id:   generateid.GenerateId(),
	//	Name: "Bulan Bintang",
	//	Address: []model.Address{
	//		{
	//			StreetName: "Jl Jalan Aja",
	//			City:       "Ragunan",
	//			PostalCode: "12345",
	//		},
	//	},
	//	Phone:   "102030",
	//	Email:   "bulan.bintang@gmail.com",
	//	Balance: 10000,
	//	UserCredential: model.UserCredential{
	//		UserName: "bulanbintang",
	//		Password: passwordHash,
	//	},
	//}
	//err := customerRepo.Create(&customer01)
	//IsError(err)

	// Save Many To Many
	product01, err := productRepo.FindById(2)
	product02, _ := productRepo.FindById(3)
	product03, _ := productRepo.FindById(4)
	customer01, err := customerRepo.FindById("8dad3fac-4c6e-4e9a-9053-7c93f9806cd1")
	IsError(err)
	fmt.Println(product01.ToString())

	err = customerRepo.OpenProductForExistingCustomer(&model.Customer{
		Id: customer01.Id,
		Products: []model.Product{
			{
				ID: product01.ID,
			},
			{
				ID: product02.ID,
			},
			{
				ID: product03.ID,
			},
		},
	})
	IsError(err)

}

func IsError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
