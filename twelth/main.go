package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


var db *sql.DB

func connectDB(){
	connStr := "host=localhost port=5432 user=postgres password =1234 dbname = testdb sslmode = disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil{
		log.Fatal("Error while opening DB", err)
	}
	err = db.Ping()
	if err != nil{
		log.Fatal("DDB isnt responding", err)
	}
	fmt.Println("Connected to SQL success")
}

func createTable(){
	query := `CREATE TABLE IF NOT EXIST students(id SERIAL PRIMARY KEY, name VARCHAR(100) NOT NULL, age INT NOT NULL, student_group VARCHAR(100) NOT NULL);`

	_, err := db.Exec(query)
	if err != nil{
		log.Fatal("Error create table", err)
	}
	fmt.Println("Students table is created")
}
type Student struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Group string `json:"group"`
}
func createStudent(w http.ResponseWriter, r *http.Request){
	var student Student

	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil{
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if student.Name == ""{
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if student.Age <= 0 {
		http.Error(w, "Age must be greater than 0", http.StatusBadRequest)
		return
	}
	if student.Group == ""{
		http.Error(w, "Group is required", http.StatusBadRequest)
		return
	}
	query := `INSERT INTO students(name,age,student_group) VALUES ($1,$2,$3) RETURNING id`
	err = db.QueryRow(query, student.Name,student.Age,student.Group).Scan(&student.ID)
	if err != nil{
		http.Error(w, "Failed to insert student", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)

}

func getStudents(w http.ResponseWriter, r *http.Request){
	rows,err := db.Query("SELECT * FROM students")
	if err != nil{
		http.Error(w, "Failed to get students", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []Student

	for rows.Next(){
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Group)
		if err != nil{
			http.Error(w, "Error to scan student", http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}



func main(){

  fmt.Println("Сервер запущен на http://localhost:8080")
  http.ListenAndServe(":8080", nil)
}