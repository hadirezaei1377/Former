package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var tmpl = template.Must(template.ParseFiles("form.html"))

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	Age         int
	Gender      string
	NationalID  string
	Description string
	Mobile      string
	IsActive    bool
}

func main() {
	http.HandleFunc("/", showForm)
	http.HandleFunc("/submit", submitForm)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func showForm(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS people (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		first_name TEXT,
		last_name TEXT,
		age INTEGER,
		gender TEXT,
		national_id TEXT,
		description TEXT,
		mobile TEXT,
		is_active BOOLEAN
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	person := Person{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
	}

	age, err := parseAge(r.FormValue("age"))
	if err != nil {
		log.Fatal(err)
	}
	person.Age = age

	person.Gender = r.FormValue("gender")
	person.NationalID = r.FormValue("national_id")
	person.Description = r.FormValue("description")
	person.Mobile = r.FormValue("mobile")
	person.IsActive = r.FormValue("is_active") == "on"

	insertQuery := `
	INSERT INTO people (first_name, last_name, age, gender, national_id, description, mobile, is_active)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err = db.Exec(insertQuery, person.FirstName, person.LastName, person.Age, person.Gender,
		person.NationalID, person.Description, person.Mobile, person.IsActive)
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, "Form submitted successfully!")
}

func parseAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}

	if age < 0 {
		return 0, fmt.Errorf("age cant be negative!")
	}

	return age, nil
}
