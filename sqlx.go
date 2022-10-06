package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func sqlxFunc() {
	// tx := db.MustBegin()
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	// tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	// tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	// tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
	// // Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	// tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	// tx.Commit()

	// // Query the database, storing results in a []Person (wrapped in []interface{})
	// people := []Person{}
	// db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	// jason, john := people[0], people[1]

	// fmt.Printf("%#v\n%#v", jason, john)
	// // Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}
	// // Person{FirstName:"John", LastName:"Doe", Email:"johndoeDNE@gmail.net"}

	// // You can also get a single result, a la QueryRow
	// jason = Person{}
	// err = db.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
	// fmt.Printf("%#v\n", jason)
	// // Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}

	// // if you have null fields and use SELECT *, you must use sql.Null* in your struct
	// places := []Place{}
	// err = db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// usa, singsing, honkers := places[0], places[1], places[2]

	// fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)
	// // Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// // Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}
	// // Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}

	// // Loop through rows using only one struct
	// place := Place{}
	// rows, err := db.Queryx("SELECT * FROM place")
	// for rows.Next() {
	// 	err := rows.StructScan(&place)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	fmt.Printf("%#v\n", place)
	// }
	// // Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// // Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}
	// // Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}

	// // Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// // which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	// _, err = db.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
	// 	map[string]interface{}{
	// 		"first": "Bin",
	// 		"last":  "Smuth",
	// 		"email": "bensmith@allblacks.nz",
	// 	})

	// // Selects Mr. Smith from the database
	// rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})

	// // Named queries can also use structs.  Their bind names follow the same rules
	// // as the name -> db mapping, so struct fields are lowercased and the `db` tag
	// // is taken into consideration.
	// rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)

	// // batch insert

	// // batch insert with structs
	// personStructs := []Person{
	// 	{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
	// 	{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
	// 	{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	// }

	// _, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
	//     VALUES (:first_name, :last_name, :email)`, personStructs)

	// // batch insert with maps
	// personMaps := []map[string]interface{}{
	// 	{"first_name": "Ardie", "last_name": "Savea", "email": "asavea@ab.co.nz"},
	// 	{"first_name": "Sonny Bill", "last_name": "Williams", "email": "sbw@ab.co.nz"},
	// 	{"first_name": "Ngani", "last_name": "Laumape", "email": "nlaumape@ab.co.nz"},
	// }

	// _, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
	//     VALUES (:first_name, :last_name, :email)`, personMaps)
}

func createTable(db *sqlx.DB) {
	// membuat table dari schema di go
	// mustexec -> menjalankan SQL script
	result := db.MustExec(schema)
	log.Println(result)
}

func sqlxQuery() {
	// configure connection untuk connect Golang dengan database
	db, err := sqlx.Connect("postgres", `
	host=localhost 
	port=5432 
	user=postgres 
	dbname=users 
	password=postgresAdmin 
	sslmode=disable`)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// check connection
	if err := db.Ping(); err != nil {
		log.Println("error in pinging the database")
		return
	}
	log.Println("success to connect to database")

	// CRUD

	// insert using normal query
	// res, err := db.Exec(`insert into users (first_name, last_name)
	// 			values('calman', 'tara');`)
	// if err != nil {
	// 	log.Fatal("error inserting value in table users")
	// 	return
	// }
	// row, _ := res.RowsAffected()
	// log.Println("success insert value, affected rows:", row)

	// insert using passing parameter
	// firstName := "john"
	// lastName := "wick"
	// res, err := db.Exec(`insert into users (first_name, last_name)
	//  		values($1, $2);`, firstName, lastName)
	// if err != nil {
	// 	log.Fatal("error inserting value in table users")
	// 	return
	// }
	// row, _ := res.RowsAffected()
	// log.Println("success insert value, affected rows:", row)

	// select
	users := []User{}
	db.Select(&users, "SELECT id, first_name, last_name FROM users WHERE id in ($1);", []int{1, 2, 3})
	println(users[0].ID, users[0].FirstName)

	// batch insert with structs
	// insertedUser := []User{
	// 	{FirstName: "Ardie", LastName: "Savea"},
	// 	{FirstName: "Sonny Bill", LastName: "Williams"},
	// 	{FirstName: "Ngani", LastName: "Laumape"},
	// }

	// _, err = db.NamedExec(`INSERT INTO users (first_name, last_name)
	//
	//	VALUES (:first_name, :last_name)`, insertedUser)
	//
	//	if err != nil {
	//		log.Println("error in batch inserting")
	//	}
}
