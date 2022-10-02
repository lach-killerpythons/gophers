// pull random lines from a SQL database

package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
	//"reflect"
)

// SELECT "animals.txt" FROM animalsDB WHERE id=515
//SELECT COUNT(*) from animalsDB

func countRows(myTable string, myDB string) {
	db, _ := sql.Open("sqlite3", myDB)

	var count int

	myCall := fmt.Sprintf("SELECT COUNT(*) FROM %s", myTable)

	err := db.QueryRow(myCall).Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
		count = 0
	default:
		fmt.Println(myDB, ":", myTable, " Number of rows =", count)
	}
}

func countRows2(myTable string, myDB string) int {
	db, _ := sql.Open("sqlite3", myDB)

	var count int

	myCall := fmt.Sprintf("SELECT COUNT(*) FROM %s", myTable)

	err := db.QueryRow(myCall).Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
		count = 0
	default:
		//fmt.Println(myDB, ":", myTable, " Number of rows =", count)

	}
	return count
}

func testRandom(myVar string) { // () {
	//myChoice := countRows2(myTable, myDB)
	i := 12
	//myString := "SELECT "animals.txt" %q" // FROM animalsDB WHERE id=%i' //  +
	//myString := "SELECT " + `"` + "animals.txt" + `"`
	//myString2 := `SELECT "animals.txt"`
	myStringX := fmt.Sprintf(`SELECT "%s.txt"`, myVar)
	myQuery := fmt.Sprintf("%s FROM animalsDB WHERE id=%d", myStringX, i)
	fmt.Println(myQuery)

}

func pickRandom_t(myTable string, myDB string, myVar string) {
	rowMax := countRows2(myTable, myDB) // finds number of rows
	rand.Seed(time.Now().UnixNano())    // new seed every time
	i := rand.Intn(rowMax)              // pick random choice
	myStringX := fmt.Sprintf(`SELECT *`)
	myQuery := fmt.Sprintf("%s FROM %s WHERE id=?", myStringX, myTable) // join the query string

	db, _ := sql.Open("sqlite3", myDB)
	myRow := db.QueryRow(myQuery, i) // Returns single row
	//var myStr string
	var id int //pointer 1
	//var animals string                 // pointer 2
	error := myRow.Scan(&id, &myVar) // Scan only returns type error
	if error == nil {
		fmt.Println("random choice: ", myVar, "id=", id)
	} else {
		log.Fatal(error)
	}

}

func pickRandom(myTable string, myDB string, myVar string) string {
	rowMax := countRows2(myTable, myDB) // finds number of rows
	rand.Seed(time.Now().UnixNano())    // new seed every time
	i := rand.Intn(rowMax)              // pick random choice
	myStringX := fmt.Sprintf(`SELECT *`)
	myQuery := fmt.Sprintf("%s FROM %s WHERE id=?", myStringX, myTable) // join the query string

	db, _ := sql.Open("sqlite3", myDB)
	myRow := db.QueryRow(myQuery, i) // Returns single row
	//var myStr string
	var id int //pointer 1
	//var animals string                 // pointer 2
	error := myRow.Scan(&id, &myVar) // Scan only returns type error
	if error != nil {
		log.Fatal(error)
	}

	return myVar

}

const animalsDB = "animalsDB"
const professionsDB = "professionsDB"
const nameDB = "nameDB"

func main() {

	start := time.Now()
	//a1 := countRows2("animalsDB", "animalsDB")
	///r1 := rand.Intn(a1)
	//fmt.Println("random one ", r1)
	//countRows("nameDB", "nameDB") // professionsDB
	//countRows("professionsDB", "professionsDB")
	//testRandom("animals.txt")

	// myTable, myDB , myColumn
	myAnimal := pickRandom(animalsDB, animalsDB, "animals")
	myJob := pickRandom(professionsDB, professionsDB, "professions.txt") //messy column names
	myFirstName := pickRandom(nameDB, nameDB, "names.txt")
	mySecondName := pickRandom(nameDB, nameDB, "names.txt")
	fmt.Println("hi my name is", myFirstName, mySecondName, "I work as a ", myJob, "my favorite animal is a", myAnimal)
	elapsed := time.Since(start)
	log.Printf("time to execute is %s", elapsed)

}
