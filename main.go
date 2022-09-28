package main

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/Calmantara/go-fga/user"
)

type Motor interface {
	Berjalan()
	PengisianDaya()
	Maintenance()
}

type Tesla struct {
	maintenanceLocation string
	batteryType         string
	batteryLevel        float64
}

// bagaimana caranya agar tesla
// dapat tergolongkan sebagai implementator dari motor?

// caranya adalah dengan implemetasi
// function yang dimiliki motor, oleh pointer dari tesla
func (t *Tesla) Berjalan() {
	// tesla akan bisa berjalan jika batery level > 30%
	if t.batteryLevel > 30.0 {
		fmt.Println("bisa berjalan")
		return
	}
	fmt.Println("harus mengisi daya")
}
func (t *Tesla) PengisianDaya() {}
func (t *Tesla) Maintenance()   {}

type Honda struct {
	maintenanceLocation string
	fuelType            string
	fuelLevel           float64
}

func (h *Honda) Berjalan() {
	// honda akan bisa berjalan jika batery level > 10%
	if h.fuelLevel > 10.0 {
		fmt.Println("bisa berjalan")
		return
	}
	fmt.Println("harus mengisi daya")
}
func (h *Honda) PengisianDaya() {}
func (h *Honda) Maintenance()   {}

func main() {
	// function yang secara default di panggil di awal penjalanan program

	// defer statement
	// statement yang membantu kita
	// untuk ensure function akan dijalankan
	// di AKHIR process
	// meskipun deklarasinya di awal process
	defer fmt.Println("process done")

	// #4 Concurrency
	// cara implementasinya -> go
	// go AsyncFunction()
	// main function tidak tau apakah
	// async function sudah dijalan dan selesai
	// atau belum

	// apakah main function bisa tau
	// function ini sudah selesai atau belum
	// bisa:
	// 	- wait group: akan menunggu hingga process selesai
	//  - channel: akan menunggu hingga mendapatkan value dari goroutine

	// goroutine:
	// 		process yang terjadi di dalam process besar lainnya
	// process besar ini (Thread)
	// go routine (Thread yang ringan atau kecil)

	// usersName := make(map[int]string, 0)

	// go func() {
	// 	for key, val := range usersName {
	// 		fmt.Println(key, val)
	// 	}
	// }()
	// for i := 0; i < 10; i++ {
	// 	usersName[i] = fmt.Sprintf("user%v", i)
	// }

	// dalam function di bawah ini
	// main function hanya melempar value ke goroutine
	// go routine mana yang execute duluan
	// main function tidak peduli dengan itu
	// for i := 0; i < 20; i++ {
	// 	go PrintName(fmt.Sprintf("name%v", i))
	// }

	// time.Sleep(time.Second * 3) // delay 5 detik

	// #5 Wait Group -> termasuk ke dalam sync package
	// variabel yang memaksa main function
	// baru boleh menjalankan process selanjutnya
	// setelah semua goroutine selesai dijalankan
	// var wg sync.WaitGroup
	// WaitGroupFunc(&wg)
	// 1. setiap goroutine yang ingin di sync,
	// wg harus add 1 di dalam variablenya
	// 2. setiap selesai menjalankan goroutine
	// wg harus memanggil done
	// 3. wg harus memanggil wait untuk menunggu semua
	// goroutine selesai
	// wg.Wait() // blocking: wait until all goroutine done

	// WaitGroupFuncInside()

	// common problem in concurrency
	fmt.Println(DataRace())
	// DataRace:
	// memory hanya boleh diakses (read / write)
	// sekali dalam satu waktu
	// Data Race terjadi ketika ->
	// memory diakses oleh program / goroutine berbeda
	// dan melakukan operasi read / write

}

func DataRace() string {
	t := "Hi"
	go func() {
		t = "Hello"
	}()
	return t
}

func WaitGroupFunc(wg *sync.WaitGroup) {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(input int) {
			defer wg.Done()
			fmt.Println("user", input)
		}(i)
	}
}

func WaitGroupFuncInside() {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(input int) {
			defer wg.Done()
			fmt.Println("user", input)
		}(i)
	}
	wg.Wait()
}

func PrintName(s string) {
	fmt.Println(s)
}

func AsyncFunction() {

	// ini akan berjalan di belakang layar
	// dibelakang process main function
	fmt.Println("inside go routine")

}

func InterfaceImpl() {
	// #1 Interface
	// -> abstrak yang merepresentasikan suatu kumpulan methods
	// yang dapat diimplementasikan dengan menggunakan struct
	var motor1 Motor
	// motor1.Berjalan() error occur here
	motor1 = &Tesla{maintenanceLocation: "US", batteryType: "Lithium", batteryLevel: 100.0}
	fmt.Println(motor1)

	motor1.Berjalan()

	var motor2 Motor
	motor2 = &Honda{maintenanceLocation: "ID", fuelType: "Pertamax", fuelLevel: 5.0}
	motor2.Berjalan()
	fmt.Println(motor1 == motor2)

	var motor3 Motor
	motor3, _ = motor3.(*Honda)
	fmt.Println("motor3 type:", reflect.TypeOf(motor3))
	// if ok {
	// 	fmt.Println("motor3 become honda type")
	// }
	motor3 = &Tesla{maintenanceLocation: "SG", batteryType: "Lithium", batteryLevel: 10.0}
	fmt.Println("motor3 type:", reflect.TypeOf(motor3))

	// back to honda
	motor3 = &Honda{maintenanceLocation: "AUS", fuelType: "Pertalite", fuelLevel: 100.0}
	fmt.Println("motor3 type:", reflect.TypeOf(motor3))
}

func EmptyInterfaceImpl() {
	// #2 Empty Interface
	// variable bebas yang dapat di assign ke siapapun
	var variable1 interface{}
	variable1 = 1 // int
	fmt.Println(variable1)
	variable1 = 10.0 // float64
	fmt.Println(variable1)
	variable1 = "Hello World" // string
	fmt.Println(variable1)
	variable1 = Honda{} // struct honda
	fmt.Println(variable1)
	// variable1 = motor1
	// fmt.Println(variable1)

	// const const1 interface{} = 1
	// const const2 Honda = Honda{}
	// const tidak akan bisa di assign dengan data type:
	//  - interface
	//  - struct

	/*
		cannot use 1 (constant of type int)
		as Motor value in assignment:
		int does not implement Motor (missing method Berjalan)
	*/
	// motor3 = 1 error occur here

	// empty interface:
	// 	tidak memiliki method
	//  sehingga type apapun bisa di assign ke interface tsb

	// abstraction interface:
	//  memiliki method/ function
	//  yang menjadi syarat minimum agar
	//  suatu type bisa diassign ke interface tsb

	// bukan suatu best practice ketika
	// kita menggunakan variable type menjadi
	// empty interface semua. karena GO bukan diciptakan untuk itu

	// assertion / type casting
	var intf1 any
	intf1 = 10
	num := intf1.(int)
	fmt.Println(num)

	// Error Assertion
	// var intf2 any
	// intf2 = 10
	// num2 := intf2.(float64)
	// "interface conversion: interface {} is int, not float64"
	// fmt.Println(num2)
}

func ReflectImpl() {
	teacher1 := user.NewTeacher("Calman", 1)
	student1 := user.NewStudent("Tara", 1)
	fmt.Println(teacher1, student1)

	// #3 Reflection
	// package built in golang
	// yang digunakan untuk manipulasi
	// data type di golang

	// menentukan type data
	int1 := 10
	var int2 any
	int2 = 11
	fmt.Println(reflect.TypeOf(int1), reflect.TypeOf(int2))
	if reflect.TypeOf(int2).Kind() != reflect.Int {
		fmt.Println("bukan int nih bos")
	}

	// deep equal
	motor1 := Honda{maintenanceLocation: "ID", fuelType: "Pertamax", fuelLevel: 10.0}
	motor2 := Honda{maintenanceLocation: "ID", fuelType: "Pertamax", fuelLevel: 10.0}
	fmt.Println(motor1 == motor2, reflect.DeepEqual(motor1, motor2))

	// ?type=sekolah
	// ?type=Sekolah
	str1 := "sekolah"
	str2 := "Sekolah"
	fmt.Println(str1 == str2,
		reflect.DeepEqual(str1, str2),
		strings.EqualFold(str1, str2))

	// deep equal -> function return boolean
	// boolean -> data type

	// numCopied := reflect.Copy(motor1, motor2)
}
