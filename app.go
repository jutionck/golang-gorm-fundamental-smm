package main

import (
	"enigmacamp.com/golang-gorm/config"
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/repository"
	"enigmacamp.com/golang-gorm/utils"
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
	//productRepo := repository.NewProductRepository(db)
	customerRepo := repository.NewCustomerRepository(db)

	// Case 1:
	// Membuat customer baru sekaligus product baru
	pwd, _ := utils.HashPassword("password")
	customer01 := model.Customer{
		Id:   generateid.GenerateId(),
		Name: "Bulan Bintang",
		Address: []model.Address{
			{
				StreetName: "Jl Jalan Aja",
				City:       "Ragunan",
				PostalCode: "12345",
			},
		},
		Phone:   "102030",
		Email:   "bulan.bintang@gmail.com",
		Balance: 10000,
		UserCredential: model.UserCredential{
			UserName: "bulanbintang",
			Password: pwd,
		},
		Products: []*model.Product{
			{
				ProductName: "Caca Marica",
			},
			{
				ProductName: "Beng Beng",
			},
		},
	}
	err := customerRepo.Create(&customer01)
	IsError(err)

}

func IsError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
