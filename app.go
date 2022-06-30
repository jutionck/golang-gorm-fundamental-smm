package main

import (
	"enigmacamp.com/golang-gorm/config"
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/repository"
	"enigmacamp.com/golang-gorm/utils"
	"fmt"
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
	productRepo := repository.NewProductRepository(db)
	customerRepo := repository.NewCustomerRepository(db)

	// Case 1:
	// Membuat customer baru sekaligus product baru
	pwd, _ := utils.HashPassword("password")
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
	//		Password: pwd,
	//	},
	//	Products: []*model.Product{
	//		{
	//			ProductName: "Caca Marica",
	//		},
	//		{
	//			ProductName: "Beng Beng",
	//		},
	//	},
	//}
	//err := customerRepo.Create(&customer01)
	//IsError(err)

	// Case 2:
	// Menambahkan product baru ke customer yang sudah terdaftar
	//cust, err := customerRepo.FindById("b1629987-21d6-4f4d-9afd-13889098d2ba")
	//IsError(err)
	//cust.Products = []*model.Product{
	//	{
	//		ProductName: "Cocolatos",
	//	},
	//}
	//err = customerRepo.UpdateBy(&cust)

	// Case 3:
	// Membuat customer baru dengan product yang sudah ada
	product01, err := productRepo.FindById(1)
	IsError(err)
	fmt.Println(product01.ToString())
	customer02 := model.Customer{
		Id:   generateid.GenerateId(),
		Name: "Rofik",
		Address: []model.Address{
			{
				StreetName: "Jalanin aja dulu gan",
				City:       "Makasar",
				PostalCode: "102020",
			},
		},
		Phone:   "20202",
		Email:   "rofik@gmail.com",
		Balance: 1000,
		UserCredential: model.UserCredential{
			UserName: "rofikrofik",
			Password: pwd,
		},
		Products: []*model.Product{&product01},
	}
	err = customerRepo.Create(&customer02)

}

func IsError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
