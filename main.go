package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Calmantara/go-fga/user" // blank import
)

func main() {
	argLen := len(os.Args[1:])
	fmt.Println("Argument length: ", argLen)

	// convert string menjadi int
	num, err := strconv.Atoi("1")
	fmt.Println(num, err)

	// assignment struct
	// properti struct dengan huruf kecil
	// juga tidak akan bisa diakses oleh package lain
	var user1 user.User

	// cara 1
	user1.ID = 1
	user1.Name = "calman"
	user1.POB = "Indonesia"
	user1.DOB = "1960-08-08"
	fmt.Println(user1)

	user1 = user.User{
		ID:   10,
		Name: "tara",
		DOB:  "1945-08-17",
		POB:  "Indonesia",
	}
	user1.CallName()

	user2 := user.User{
		ID:   10,
		Name: "tara2",
		DOB:  "1945-08-17",
		POB:  "Indonesia",
	}
	user2.CallName()

	// teacher dan user
	teacher1 := user.Teacher{ID: 99, Name: "calman"}
	// teacher1.CallName() -> tidak valid, teacher tidak ada func CallName

	// 1. kita ga bisa langsung assign di dalam structnya
	// 2. kita bisa langsung assign property di dalamnya
	student1 := user.Student{user.User{ID: 1}}
	student1.Name = "TARA"
	student1.CallName()
	// student1.init()
	fmt.Println(teacher1, student1)

	// apakah struct bisa memiliki
	// fungsinya dia sendiri?
	user3 := user2
	fmt.Println(user3.GetAddress())     // kosong
	user3.SetAddress("Jakarta Selatan") // address updated
	fmt.Println(user3.GetAddress())     // keluar Jakarta Selatan

	// POINTER
	// kalau variable biasa, kita tambahin & didepannya
	// artinya kita akan mengakses memory address dari variabel tsb

	// kalau variable pointer kita kasih * didepannya
	// artinya kita akan mengakses value dari pointer tsb
	num = 9

	var num1 *int
	num1 = &num // num1 akan mengakses memory yang sama dengan num
	fmt.Println(num1, *num1)

	num = 10
	*num1 = 100

	// pointer of struct
	ptrUser := &user.User{
		ID:   100,
		Name: "Tara",
	}
	fmt.Println(ptrUser, *ptrUser)
	ptrUser.SetAddress("hello")
	fmt.Println(ptrUser.GetAddress())

	usr := user.User{
		ID:   99,
		Name: "Calman",
	}
	usr.SetAddress("BSD")
	fmt.Println(usr.GetAddress())

	// function input parameter pointer
	cNum := 100
	ChangeNum(&cNum, 1000)
	fmt.Println(cNum)

	var ptrCNum *int
	ptrCNum = &cNum
	ChangeNum(ptrCNum, 10000)

	// input param as pointer
	ChangeName(&usr, "Tara")
	fmt.Println(usr)

	// 1. init data student pada init function
	// 2. program akan mengeluarkan pesan error, ketika id tidak ditemukan
	// "student dengan id xx tidak ada pada database"
	// 3. id harus positive integer
}

func ChangeNum(source *int, destination int) {
	*source = destination
}

func ChangeName(src *user.User, name string) {
	// (*src).Name = name
	src.Name = name
}
