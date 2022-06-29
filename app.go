package main

import (
	"enigmacamp.com/golang-gorm/config"
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/repository"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

	// Insert Customer
	//password, _ := HashPassword("password")
	//customer01 := model.Customer{
	//	Id:      generateid.GenerateId(),
	//	Name:    "Bulan Sutisna",
	//	Address: "Bali",
	//	Balance: 20000,
	//	UserCredential: model.UserCredential{
	//		UserName: "bulanbulanan",
	//		Password: password,
	//	},
	//	Email: "bulan.s@gmail.com",
	//	Phone: "20202020",
	//}
	//repo.Create(&customer01)

	// Update Existing
	//customer02 := model.Customer{
	//	Id: "f349932f-5aa8-45d2-b6b3-b78bc4b1e213",
	//}
	//customer02, err := repo.FindById(customer02.Id)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//fmt.Println("FindById: ", customer02)
	//userCredential01 := model.UserCredential{
	//	UserName: "bulansehatsehatya",
	//	Password: "password",
	//}
	//customer02.UserCredential = userCredential01
	//err = repo.UpdateBy(&customer02)
	//if err != nil {
	//	log.Println(err.Error())
	//}

	// Update with Preload
	customer02, err := repo.FindFirstWithPreload(map[string]interface{}{"id": "f349932f-5aa8-45d2-b6b3-b78bc4b1e213"}, "UserCredential")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Before: ", customer02)
	c := customer02.(model.Customer)
	c.UserCredential.Password = "rahasianegara"
	err = repo.UpdateBy(&c)
	if err != nil {
		log.Println(err.Error())
	}
	customer02, err = repo.FindFirstWithPreload(map[string]interface{}{"id": "f349932f-5aa8-45d2-b6b3-b78bc4b1e213"}, "UserCredential")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("After: ", customer02)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}
