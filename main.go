package main

import (
	"fmt"
	"reflect"

	"github.com/Calmantara/go-fga/user"
)

var class string

func main() {
	/***PACKAGE***/
	// this is read function
	Read()
	// call user
	user.CallUser("calman")

	/***Data Type and Variable***/

	/* imutable:
	- int -> uint, uint8, uint16, uint32 . . . int64
	- float -> 32 dan 64
	- string
	- bool
	*/
	class = "kelas 1"

	fmt.Print("\n IMMUTABLE \n")
	var name string
	name = "calman"
	fmt.Println(name)

	var name2 string = "calman"
	fmt.Println(name2)

	name3 := "calman2"
	fmt.Println(name3)

	var unsignedInt uint
	unsignedInt = 1000
	fmt.Println(unsignedInt)

	// multi declaration
	a1, a2, a3, a4 := "string1", "string2", "string3", "string4"
	fmt.Println(a1, a2, a3, a4)

	d1, d2 := "string", true
	fmt.Println(d1, d2)

	/* mutable:
	- slice / array
	- map
	*/
	fmt.Print("\n MUTABLE \n")
	var arr1 []int
	arr1 = []int{1, 2, 3, 4}
	fmt.Println(arr1)

	arr2 := []int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	map1 := map[string]int{"nama": 1}
	fmt.Println(map1["nama"])

	// constant
	const c1 string = "CALMAN"
	// c1 = "CALMAN LAGI" tidak akan pernah terjadi

	// Aritamtik
	// + - / * %
	fmt.Print("\n ARITMATIK \n")
	v1 := 1 + 1
	v2 := (1 * 1) + 6/10
	fmt.Println(v1, v2)

	// relational
	// ==, !=, >, <, >=, <=
	// string hubungan erat dengan ASCII
	fmt.Print("\n RELATIONAL \n")
	fmt.Println(name == name2)
	fmt.Println(name == name3)
	fmt.Println(name != name3)
	fmt.Println(name > name3)
	fmt.Println(name < name3)
	fmt.Println(name >= name3)
	fmt.Println(name <= name3)

	// logical
	// AND => &&
	// OR => ||
	// NEGASI (NOT) => !
	fmt.Print("\n LOGICAL \n")
	fmt.Println(true && false) // false
	fmt.Println(true && true)  // true
	fmt.Println(false || true) // true
	fmt.Println(!false)        // true

	// name = calman
	// name2 = calman
	fmt.Println((name == "calman") && (name2 != "aldo"))
	// cara baca:
	// name == calman -> true
	// name2 != aldo -> true
	// true AND true -> true

	// bitwise operator (nice to have)
	// >>, <<, |, &
	fmt.Print("\n BITWISE \n")
	variable1 := 5              // 00000101
	fmt.Println(variable1 >> 1) // 00000010
	fmt.Println(variable1 << 2) // 00010100

	// memasukkan data type ke string
	// menambahkan anotasi berdasarkan data type di dalam string
	fmt.Print("\n String Format \n")
	fmt.Printf("hello:%s age:%d message:%v", name, 10, arr1)

	// pengecekan data type
	fmt.Print("\n Pengecekan Data Type \n")
	fmt.Println(reflect.TypeOf(name))

	// global variable
	Function()
}

func Function() {
	fmt.Println(class)
}
