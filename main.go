package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Default handler")
}

type Config struct {
	Database ConfigDatabase
}

type ConfigDatabase struct {
	DataSource string
}

type Person struct {
	Id   int
	Name string
}

var database ConfigDatabase

func main() {

	err := envconfig.Process("database", &database)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "DataSource: %v\n"
	_, err = fmt.Printf(format, database.DataSource)
	if err != nil {
		log.Fatal(err.Error())
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations",
	}
	migrate.SetTable("migrations")
	migdb, err := sql.Open("mysql", database.DataSource)
	if err != nil {
		log.Fatal(err.Error())
	}

	n, err := migrate.Exec(migdb, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Applied %d migrations!\n", n)

	http.HandleFunc("/", Index)
	http.HandleFunc("/people", HandlePeople)

	s := &http.Server{
		Addr: ":8080",
		// Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

	// router := httprouter.New()
	// router.GET("/", Index)
	// router.GET("/hello/:name", Hello)

	// log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func HandlePeople(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ListPeople(w, r)
	} else if r.Method == http.MethodPost {
		PostPeople(w, r)
	} else if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	} else {
		http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
	}
}

func ListPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("ListPeople")
	db, err := gorm.Open(mysql.Open(database.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var people []Person

	db.Model(&Person{}).Find(&people)
	res, err := json.Marshal(people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(res))
}

func PostPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("PostPeople")
	db, err := gorm.Open(mysql.Open(database.DataSource), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	person := Person{
		Name: "aaa",
	}
	result := db.Create(&person)
	log.Println(result.Name)

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

}
