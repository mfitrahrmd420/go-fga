package main

import "fmt"

func StudentPerkalian(x1, x2 int) {
	fmt.Println(x1 * x2)
}
func StudentPertambahan(x1, x2 int) {
	fmt.Println(x1 + x2)
}

// input -> output
// () dalam function -> untuk menampung input
// output dideklarasikan setelah tanda kurung

// input dan output tidak wajib
func Perkalian(in int) (out int) {
	out = in * 10
	return out
}

func Perkalian2(in int) int {
	return in * 10
}

func DoProcess() {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, Perkalian(i))
	}
	fmt.Println(arr)

	for i, v := range arr {
		arr[i] = Perkalian2(v)
	}
	fmt.Println(arr)
}

// kita ingin membuat function
// yang mengeluarkan output
// function juga

func GenerateOtherFunction() func(int) int {
	return func(i int) int {
		return i * i
	}
}

// callback function
// -> menjadikan type function sebagai input parameter
// dari suatu function
// input ini, bisa kita panggil di dalam func tersebut
// kita juga bisa menambahkan proses lainnya dalam
// pemanggilan func input ini
func CallbackFunc(in int, f func(int)) {
	f(in)
}

func LoopingProcess(executionAmount int, fin func()) {
	// function ini akan menjalankan
	// function fin sebanyak in kali

	// fin adalah function yang akan dieksekusi

	for i := 0; i < executionAmount; i++ {
		fmt.Println("eksekusi ke", i)
		fin()
	}
}

// mengaliaskan suatu function
type FunctionAlias func(int) int

func CallAliasFunction(in int, alias FunctionAlias) int {
	return alias(in)
}

// selain main function
// init function -> special function
// init function akan dijalankan pertama kali
// sebelum main function

// init function
// biasa digunakan jika kita perlu
// inisiasi nilai global variable
// atau menjalankan mencari configurasi
func init() {
	for i := 0; i < 10; i++ {
		fmt.Println("print from init function in main", i)
	}
}

func Misc() {
	// function default yang wajib ada di program Go
	// DoProcess()

	// scoping object in go
	// user.StoreUser(1, "calman")
	// fmt.Println(user.GetUser(1))

	// closure function
	generatedFunc := GenerateOtherFunction()
	fmt.Println(generatedFunc(9))
	// 0x104f27720
	// 0x105083720
	generatedFunc2 := func(int) int {
		return 1
	}
	generatedFunc2(1)

	// callback function
	CallbackFunc(1, func(i int) {
		fmt.Println(i)
	})
	LoopingProcess(5, func() {
		// function ini yang akan
		// dieksekusi sebanyak 5x
		a := 1 + 1
		b := a + 2
		fmt.Println(b)
	})

	// aliasing function
	var alias FunctionAlias
	alias = func(in int) int {
		return in * in
	}
	fmt.Println(CallAliasFunction(9, alias))

	// study case
	// aku ada suatu nama student
	// kalau student ini bernama calman
	// dia akan menjalankan proses perkalian
	// kalau student ini bernama tara
	// dia akan menjalankan proses pertambahan
	name := "calman"
	x1, x2 := 4, 5

	// if name == "calman" {
	// 	fmt.Println(x1 * x2)
	// } else {
	// 	fmt.Println(x1 + x2)
	// }

	if name == "calman" {
		StudentPerkalian(x1, x2)
	} else {
		StudentPertambahan(x1, x2)
	}
	// other method
	name = "tara"
	exec := StudentPerkalian
	if name == "tara" {
		exec = StudentPertambahan
	}
	exec(x1, x2)
}
