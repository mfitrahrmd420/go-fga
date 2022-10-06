package main

import (
	_ "github.com/lib/pq"
)

// sql untuk membuat table
var schema = `
CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    last_name TEXT,
    first_name TEXT
);`

// define struct untuk user
type User struct {
	ID        uint64 `json:"id" db:"id" gorm:"id"` // tag suatu property struct
	FirstName string `json:"first_name" db:"first_name" gorm:"first_name"`
	LastName  string `json:"last_name" db:"last_name" gorm:"last_name"`
}

type Class struct {
	ID     uint64 `json:"id" db:"id" gorm:"id"` // tag suatu property struct
	Name   string `json:"name" db:"name" gorm:"name"`
	UserID uint64 `json:"user_id" db:"user_id" gorm:"user_id"`
}

// tag ini akan dibaca oleh sqlx sebagai properti yang valid untuk digunakan

func main() {
	// Clean Architecture
	// bersifat tidak mutlak: setiap company memiliki standard masing2

	// aku bagi berdasarkan kesamaan job desc
	// usecase -> tempat menyimpan business logic
	// repository -> codingan yang berhubungan sama database
	// server -> untuk menghandle server yang dijalankan (gin)
	// domain -> untuk mengelompokkan di suatu project, ada model apa aja
}
