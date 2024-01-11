package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Student struct {
	Name  string
	Class int
	Mark  int
}

func main() {
	connStr := "user=testing dbname=testing password=testing sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Server Start Listening http://localhost:8080")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/insert.html", insertShowHandler)
	http.HandleFunc("/read.html", readShowHandler)
	http.HandleFunc("/update.html", updateShowHandler)
	http.HandleFunc("/delete.html", deleteShowHandler)

	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/update", updateHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, nil)

}

func insertShowHandler(w http.ResponseWriter, r *http.Request) {

	tmp, err := template.ParseFiles("src/insert.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, nil)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=testing dbname=testing password=testing sslmode=disable" // Connect the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	name := r.FormValue("name")
	class := r.FormValue("class")
	mark := r.FormValue("mark")

	query := "INSERT INTO student (name, class, mark) VALUES ($1, $2, $3)"

	_, err1 := db.Exec(query, name, class, mark)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func readShowHandler(w http.ResponseWriter, r *http.Request) {

	connStr := "user=testing dbname=testing password=testing sslmode=disable" // Connect the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err1 := db.Query("SELECT * FROM student") // Execuate tthe Query
	if err1 != nil {
		panic(err)
	}
	defer result.Close()

	var students []Student
	for result.Next() { // Load the data in data structure
		var student Student
		result.Scan(&student.Name, &student.Class, &student.Mark) // Connect the database
		students = append(students, student)
	}

	tmp, err2 := template.ParseFiles("src/read.html")
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, students) // Pass the data
}
func updateShowHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("src/update.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, nil)
}
func deleteShowHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("src/delete.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, nil)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=testing dbname=testing password=testing sslmode=disable" // Connect the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	name := r.FormValue("name")

	query := "DELETE FROM student WHERE name = $1"
	_, err1 := db.Exec(query, name)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=testing dbname=testing password=testing sslmode=disable" // Connect the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	name := r.FormValue("name")

	query := "SELECT class, mark FROM student WHERE name = $1"
	row := db.QueryRow(query, name)

	var class, mark int
	err1 := row.Scan(&class, &mark)
	if err1 != nil {
		http.Error(w, "Name not found", http.StatusNotFound)
		return
	}

	temp, err2 := template.ParseFiles("src/update.html")
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Name  string
		Class int
		Mark  int
	}{
		Name:  name,
		Class: class,
		Mark:  mark,
	}

	temp.Execute(w, data)
}
func updateHandler(w http.ResponseWriter, r *http.Request) {
	connStr := "user=testing dbname=testing password=testing sslmode=disable" // Connect the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	name := r.FormValue("name")
	class := r.FormValue("class")
	mark := r.FormValue("mark")

	query := "UPDATE student SET class = $2, mark = $3 WHERE name = $1"
	_, err1 := db.Exec(query, name, class, mark)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
