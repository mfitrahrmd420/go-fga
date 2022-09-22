package main

import "fmt"

func main() {
	// ConditionalStatement()
	// Loop()
	// ArrayAndSlice()
	// HashMap()
	// Alias()
	// StringInDepth()
	MiniQuiz()
}

func ConditionalStatement() {
	// => membuat program kita bisa menjalankan suatu logic
	// bercabang, dengan logika matematika (true false)

	// kalau hujan aku wfh
	// kalau cerah aku wfo

	// if else
	cuacaHujan := "hujan"
	cuacaCerah := "cerah"
	cuacaHariIni := "cerah"

	// hanya mengeksekusi di dalam if ketika bernilai true
	// cuaca hari ini cerah == hujan? false
	if cuacaHariIni == cuacaHujan {
		fmt.Println("kerja dari rumah")
	}
	if cuacaHariIni == cuacaCerah {
		fmt.Println("kerja dari kantor")
	}

	// ketika angka1
	// >0 dan <6 => ok
	// >= 6 dan <=10 => not ok
	// >= 10 => maybe ok
	angka1 := 10
	if angka1 > 0 && angka1 < 6 {
		fmt.Println("ok")
	} else if angka1 >= 6 && angka1 <= 10 {
		fmt.Println("not ok")
	} else if angka1 >= 10 {
		fmt.Println("maybe ok")
	}

	// switch
	switch cuacaHariIni {
	case cuacaHujan:
		fmt.Println("ok")
	case cuacaCerah:
		fmt.Println("very ok")
	}

	angka1 = -1
	switch {
	case angka1 > 0 && angka1 < 6:
		fmt.Println("ok")
	case angka1 >= 6 && angka1 <= 10:
		fmt.Println("not ok")
		fallthrough
	case angka1 >= 10:
		fmt.Println("maybe ok")
	default:
		fmt.Println("worst")
	}

	// best practice
	var lokasiKerja string // assign variable kosong
	if cuacaHariIni == cuacaHujan {
		lokasiKerja = "wfh"
	} else if cuacaHariIni == cuacaCerah {
		lokasiKerja = "wfo"
	}

	// jawaban 1
	if cuacaHariIni == cuacaHujan {
		lokasiKerja = "wfh"
	} else {
		lokasiKerja = "wfo"
	}

	// jawaban 2
	lokasi := "wfo" // assign variable dengan default value
	if cuacaHariIni == cuacaHujan {
		lokasi = "wfh"
	}
	fmt.Println(lokasiKerja, lokasi)

	if cuaca := "hujan"; cuaca == cuacaHujan {
		fmt.Println(cuaca)
	}
	// fmt.Println(cuaca)
}

func Loop() {
	// ++ statement
	num := 0
	num++
	fmt.Println(num)

	// looping di program => akan digunakan ketika:
	// 1. mengulang suatu proses hingga hitungan tertentu
	// 2. kita harus mengakses semua isi slice / map

	// for loop
	// ngeprint angka secara sekuensial dari 1 - 100
	for i := 0; i < 100; i++ {
		// fmt.Println(i)
	}

	k := 0
	for k < 100 {
		// fmt.Println(k)
		k++
	}

	// break
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
			// akan mengeluarkan program dari looping
		}
	}

	// continue statement
	// jika i == 6, berarti dia ga valid, tanpa ngeprint valid lagi
	for i := 0; i < 10; i++ {
		if i == 6 {
			// tidak valid
			fmt.Println("invalid")
			continue
			// continue akan melanjutkan hitungan loop
			// tanpa mengeksekusi line dibawahnya
		}
		fmt.Println("valid1")
		fmt.Println("valid2")
		fmt.Println("valid3")
		fmt.Println("valid4")
	}

	// inifite for loop
	j := 0
	for {
		if j == 1000 {
			break
		}
		j = j + 10
	}
	fmt.Println("infinit loop executed", j)
}

func ArrayAndSlice() {
	// assign
	arrInt := []int{1, 2, 3, 4}

	// append -> memasukkan element kedalam suatu array
	arrInt = append(arrInt, 1)

	// get data
	// 0:1, 1:2, 2:3, 3:4, 4:1
	fmt.Println(arrInt[2])
	fmt.Println(arrInt[:3])

	// jika ada array dengan jumlah elm 5,
	// apa yang terjadi kalau aku akses [6]
	// fmt.Println(arrInt[6]) // ini kana panic

	// loop over
	// print element yang ada di array

	// pendekatan pertama
	// len -> untuk mendapatkan jumlah elm dalam array
	fmt.Println(len(arrInt))
	for i := 0; i < len(arrInt); i++ {
		fmt.Println(arrInt[i])
	}

	// pendekatan kedua
	for i, v := range arrInt {
		// i -> mengindikasikan index keberapa
		// v -> mengindikasikan value di index tersebut
		fmt.Println(i, v)
	}

	// quiz
	// [1,2,3,4]
	// [5,6,7,8]
	// [1,2,3,4,5,6,7,8]
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{5, 6, 7, 8}

	// jawaban 1
	// ... di go -> rentetan value
	// param append -> elems ...Type
	arr3 := append(arr1, arr2...)
	fmt.Println(arr3)

	// jawaban 2
	for _, v := range arr2 {
		arr1 = append(arr1, v)
	}

	// misal
	arrInt2 := []int{1, 2, 3, 4}
	arrInt3 := arrInt2[0:3]
	arrInt3[1] = 10
	fmt.Println(arrInt2, arrInt3)

	arrMD1 := [][][]int{{{1, 2, 3, 4}, {1, 2, 3, 4}}}
	fmt.Println(arrMD1)

	// real case pengguanaan
	type Struct1 struct {
		ArrString []string
		ArrInt    []int
	}
	arrMD2 := []Struct1{
		{ArrString: []string{"abc", "def"}, ArrInt: []int{1, 2, 3, 4, 5}},
		{ArrString: []string{"abc", "def"}, ArrInt: []int{1, 2, 3, 4, 5}},
	}
	fmt.Println(arrMD2)
}

func HashMap() {
	// assign
	map1 := map[string]int{"key1": 1, "key2": 2}
	map2 := make(map[string]int, 0)

	// nil -> map belum mengalokasikan memory apa apa
	var map3 map[string]int

	fmt.Println(map1, map2, map3)
	// get data
	fmt.Println(map1["key1"])
	val, ok := map1["key3"]
	val, ok = map1["key2"]
	// doing validation
	if ok {
		fmt.Println(val)
	}
	fmt.Println(val, ok)
	// valid
	val, ok = map2["key1"]
	val, ok = map3["key1"]

	// assign value ke map
	fmt.Println(map3 == nil)
	map2["key1"] = 1

	map3 = map2
	map3["key1"] = 1

	// delete
	delete(map3, "key1")

	// kenapa map2 juga jadi 0
	// -> mengakses memory yang sama

	// loop over
	for key, val := range map1 {
		fmt.Println(key, val)
	}

	// interface{}
	// any
	// map4 := make(map[string]interface{})
	// map4 := make(map[string]any)
}

func Alias() {
	// menghindari ambigu specific type

	// assign alias
	// immutable: int, string, dll
	// mutable: map, arr, dll

	type CalmanVariable string
	var1 := []CalmanVariable{"cal", "man"}
	fmt.Println(var1)

	age := 1
	// CalmanCustom(age) // tidak bisa digunakan, karena beda type alias

	// casting alias
	//  -> mengubah type dari variable
	CalmanCustom(Calman(age))
}

// example
type Calman int

func CalmanCustom(age Calman) {
	fmt.Println(age)
}

func StringInDepth() {
	// -> kumpulan dari rune / ascii / byte
	str := "abcdefg" // c -> rune sendiri, a -> rune sendiri dst
	// ascii / rune / byte
	for _, val := range str {
		// val -> angka ascii / rune dari suatu character
		fmt.Println(val)
		fmt.Println(string(val))
	}
	// typecast string -> array of bytes
	by := []byte(str)
	for _, v := range by {
		fmt.Println(v)
	}
}

func MiniQuiz() {
	// quiz 1
	// 1. ada array of alias Int -> alias dari int
	// 2. aku mau menghitung berapa jumlah Int yang sama dalam 1 array
	// 3. aku ingin ngeprint value Int beserta jumlahnya

	// [1,2,3,4,1,1,1,2,3,4,5,6,7]Int
	// hasil -> 1:4, 2:2, 3:2, 4:2, 5:1, 6:1, 7:1
	type Int int
	arrInt := []Int{1, 2, 3, 4, 1, 1, 1, 2, 3, 4, 5, 6, 7}
	mapInt := map[Int]int{}

	for _, val := range arrInt {
		mapInt[val]++
	}

	for key, val := range mapInt {
		fmt.Println(key, val)
	}

	// quiz 2
	// 1. aku ada array of string
	// 2. aku akan membuat array string baru
	// 	- jika jumlah element dalam array tsb >= 3

	// [cal, cal, cal, man, man, tar, tar, ra, ra, ra]
	// [cal, ra]
	str := []string{"cal", "cal", "cal", "man", "man", "tar", "tar", "ra", "ra", "ra"}
	mapSum := map[string]int{}
	mapDuplicate := map[string]bool{}

	var newArr []string
	for _, v := range str {
		mapSum[v]++
		if (mapSum[v] >= 3) && !mapDuplicate[v] {
			newArr = append(newArr, v)
			mapDuplicate[v] = true
		}
	}
	fmt.Println(newArr)
}
