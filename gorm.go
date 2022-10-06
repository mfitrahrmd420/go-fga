package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func gormFunc() {
	// gorm -> ORM yang digunakan untuk menuliskan SQL tetapi di wrap menjadi suatu function

	// connect to database
	dsn := "host=localhost user=postgres password=postgresAdmin dbname=users port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// GORM -> exec, functions
	// ketika kita ingin membuat table
	// db.Exec(``)
	// db.AutoMigrate(&Class{}) // function akan membuat table class di pg kita

	// dalam menggunakan GORM
	// setidaknya ada 2 cara untuk menuju ke table yang kita inginkan
	// .Model()
	// .Table()

	// Get
	var users []User
	db.Model(&User{}).
		Select("id, first_name").
		Where("id in ?", []int{1, 2, 3}).
		Find(&users)
	for _, val := range users {
		fmt.Println(val)
	}

	// insert
	// user := User{FirstName: "john", LastName: "smith"}
	// db.Model(&User{}).
	// 	Create(&user)

	// batch insert
	// insertedUser := []User{
	// 	{FirstName: "okta", LastName: "lia"},
	// 	{FirstName: "sam", LastName: "heri"},
	// }
	// db.Model(&User{}).
	// Create(&insertedUser)

	// update
	tx := db.Model(&User{}).Where("id = 1").Update("first_name", "putra")
	if tx.Error != nil {
		log.Fatal("error update user with id 1")
	}

	// delete
	// tx = db.Model(&User{}).Where("id = 1").Delete(&User{})
	// if tx.Error != nil {
	// 	log.Fatal("error update user with id 1")
	// }

	deletedUser := User{ID: 2}
	tx = db.Model(&User{}).Delete(&deletedUser)
	if tx.Error != nil {
		log.Fatal("error update user with id 1")
	}

	// khasusnya Preload (eager loading)
	// hook
}
