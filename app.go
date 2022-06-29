package main

import (
	"enigmacamp.com/golang-gorm/config"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	cfg := config.NewConfig()
	cfg.DbConn()
	defer func(cfg *config.Config) {
		err := cfg.DbClose()
		if err != nil {
			log.Println(err.Error())
		}
	}(&cfg)

	//repo := repository.NewCustomerRepository(db)
	//pass, _ := HashPassword("password")
	//customer01 := model.Customer{
	//	Id:             generateid.GenerateId(),
	//	Name:           "Fadli Zona Nyaman",
	//	Phone:          "14045",
	//	Email:          "orderaja@gmail.com",
	//	Balance:        5000,
	//	UserCredential: model.UserCredential{UserName: "fadlizone", Password: pass},
	//	Address: []model.Address{
	//		{
	//			StreetName: "Jl Nin Aja",
	//			City:       "Bogor",
	//			PostalCode: "123",
	//		},
	//		{
	//			StreetName: "Jl Terus",
	//			City:       "Bandung",
	//			PostalCode: "456",
	//		},
	//	},
	//}
	//err := repo.Create(&customer01)
	//isError(err)
	//customer02, err := repo.FindFirstWithPreload(
	//	map[string]interface{}{"id": "e439f380-20e9-41b8-aa95-0db32456e22f"},
	//	"UserCredential",
	//)
	//isError(err)
	//log.Println(customer02.ToString())
}

func isError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
