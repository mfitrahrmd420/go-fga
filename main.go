package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	// penggunaan dari channel
	// pembuatan worker pool

	// [1,2,3,4,5,6,7,8,9]

	chanInt := make(chan int)
	go func() {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := 0; i < len(arr); i++ {
			chanInt <- arr[i]
		}
		close(chanInt)
	}()

	for i := 0; i < 3; i++ {
		go func() {
			for val := range chanInt {
				fmt.Println(val * 10)
			}
		}()
	}

	time.Sleep(10 * time.Second)

}

func channeling() {
	// defer recoveryFunction()

	// _ = errors.New("error occur in program")
	// panic(err)

	// panic vs exit
	// panic bisa di tangkap oleh recover function
	// terhadap error yang terjadi
	// exit akan langsung mengeluarkan program
	// os.Exit(1)

	// Channel
	// pointer yang digunakan untuk
	// komunikasi antar goroutine
	// secara aman -> bisa menghindari data race

	chanInt := make(chan int, 3)

	chanInt2 := make(chan int)
	chanInt3 := make(chan int)
	// chanInt <- 10 // artinya kita memberikan value ke channel tsb
	// <- chanInt // artinya kita mengambil value dari channel tsb

	// deadlock -> tidak bisa memproses process selanjutnya
	go func() {
		for i := 0; i < 3; i++ {
			chanInt <- i // berhenti tapi di background
		}
		// close -> mengindikasikan bahwa
		// channel sudah tidak dapat dimasukin nilai lagi
		// close(chanInt)

		// "send on closed channel"
		// chanInt <- 0
	}()

	go func() {
		// name: go function 1
		for val := range chanInt3 {
			chanInt2 <- val
		}
		close(chanInt2)
	}()

	go func() {
		// name: go function 2
		for i := 0; i < 10; i++ {
			chanInt3 <- i
			time.Sleep(time.Duration(int(time.Millisecond) * i))
		}
		close(chanInt3)
	}()

	// go function 2
	// 	-> bertugas untuk memasukkan value ke chan3
	//  ->

	// go function 1
	// -> me-listen channel3
	// -> dia juga memasukkan nilai ke channel2

	// main function
	// me-listion channel2

	// ketika channel di assign
	// proses assignment akan diproses lebih lanjut setelah
	// channel di ambil juga nilainya

	// jika hal seperti ini terjadi, akan terjadi deadlock

	// for val := range chanInt {
	// 	// proses ini akan mendengarkan channel
	// 	// sehingga ketika channel di assign suatu value,
	// 	// dia akanlangsung menangkap value tcb
	// 	fmt.Println(val)
	// }

	// untuk menghindari deadlock saat assignment
	// 1. bisa menggunakan close
	// 2. kita hanya menerima sebanyak channel capacity

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(<-chanInt)
	// }

	// for {
	// 	select {
	// 	case int1 := <-chanInt2:
	// 		fmt.Println("got from chan2", int1)
	// 	case int2 := <-chanInt3:
	// 		fmt.Println("got from chan3", int2)
	// 	}
	// }

	for val := range chanInt2 {
		fmt.Println(val)
	}

	// Dengan menggunakan channel, kita bisa tau
	// kapan sebuah go routine selesai menjalankan program / tugasnya

	// apakah goroutine + channel
	// bisa mempercepat program kita?
	// 	tidak selalu -> semakin banyak goroutine, dia akan memakan resource
}

func panicExplain() {
	defer recoveryFunction()
	// deferAndExit()

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Printf("end of loop%v\n", i)
	// }
	// fmt.Println("outside function deferAndExit")
	var err error
	email := "calman@gmail.com"
	if err = isEmailExist(email); err != nil {
		fmt.Printf("email not valid:%v\n", err)
		panic(err)
		return
	}
	//"runtime error: invalid memory address or nil pointer dereference" PANIC
	fmt.Printf("email is exist. you are ready to go! got error:%v", err.Error())

	// Panic
	// untuk mengeluarkan program dengan suatu indikasi
	// misalnya, ketika program gagal connect to database
	// program tidak seharusnya berjalan
	// ketika kita mencoba memanggil suatu function dari interface
	// ketika interface itu masih nil

	// panic => dia akan langsung mengeluarkan program
	// err => data type yang akan menampung informasi kesalahan pada function / program kita

	// Contoh Error:
	// salah masukin password -> panic?
	// cuma error "password anda salah" <= error
}

func recoveryFunction() {
	// recover akan menangkap error yang terjadi saat panic
	// atau saat program selesai dijalankan

	// kalau di bahasa pemrograman lain
	// seperti catch function

	if err := recover(); err != nil {
		fmt.Printf("program exit caused by error:%v\n", err)
		return
	}
	fmt.Printf("program exit normallay\n")
}

func isEmailExist(email string) (err error) {
	emails := map[string]bool{
		"calman@gmail.com": true,
		"tara@gmail.com":   true,
		"abdi@gmail.com":   true,
		"gulam@gmail.com":  true,
	}
	if !emails[email] { // ketika tidak exist
		// err = errors.New("email is not exist in our system")

		// tergantung linter => convention yang digunakan saat ngoding
		// error tidak boleh berakhir dengan tanda seru / line baru
		// error tidak boleh mengandung huruf kapital
		err = fmt.Errorf("%v IS NOT EXIST in our system", email)
		return err
	}
	// error bis di check apakah dia nil
	// atau memiliki nilai
	return nil
}

func errorHandling() {
	// error, panic, recover

	// error => situasi yang tidak diinginkan
	// baik itu dari data yang tidak valid,
	//   - password salah
	//   - user tidak ditemukan
	// kondisi yang tidak normal
	//   - database tidak bisa connect
	//   - menghubungkan dengan server lain, tapi tidak bisa connect
	// kegunaan => mengindikasikan bahwa program kita atau data kita
	// tidak baik baik saja
	var err error // interface dengan function Error() => mengubah error menjadi string

	err = errors.New("this is custom error")
	fmt.Println(err.Error() == "this is custom error")
	// error biasanya digunakan untuk return / output dari suatu function
}

func deferAndExit() {
	// defer recoverFunction()
	var pwdEnv string
	defer func() {
		fmt.Printf("current dir is:%v\n", pwdEnv)
	}()

	name := "calman1"
	if name == "calman" {
		fmt.Println("name is calman")
		return // dia akan return di sini
		// dia keluar di sini, dan memanggil defer
	}
	fmt.Println("PAKSA KELUAR")
	// os package adalah package yang
	// mengakses system komputer kita secara langsung

	for i := 0; i < 10; i++ {
		defer fmt.Printf("end of loop in function%v\n", i)
	}

	// 0 mengindikasikan success
	// != 0 mengindikasikan error
	// exit => akan memaksa program untuk keluar
	// exit ini berguna nantinya untuk:
	// 	 1. mengetahui program kita mati karena apa
	//   2. kita bisa memanfaatkan grace exit di program go
	// 			- ini akan dibahas pada web application
	// os.Exit(1) // di line ini, dia akan langsung mematikan program go kita
	pwdEnv = os.Getenv("PWD")
	fmt.Println(pwdEnv)

	hostName, _ := os.Hostname()
	fmt.Println(hostName)

	fmt.Println("this is not calman")
}
