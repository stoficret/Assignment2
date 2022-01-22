package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//Global variable for microservice tutor
var students Student
var StudentID string
var db *sql.DB

type Student struct { // map this type to the record in the table
	StudentID string
	SName     string
	DOB       string
	Address   string
	PhoneNo   string
}

/* For Console
var ukey string
    db.QueryRow("Select LEFT(MD5(rand()),16)").Scan(&ukey)
    ukey = "T"+ukey
*/
////////////////////////////////////////////////////////////////////////////////////////////////////////
////																								////
////							Functions for MySQL Database										////
////																								////
////////////////////////////////////////////////////////////////////////////////////////////////////////
//Registering new student
func CreateNewStudent(db *sql.DB, s Student) {
	query := fmt.Sprintf("INSERT INTO Students VALUES ('%s', '%s', '%s')",
		s.StudentID, s.SName, s.DOB, s.Address, s.PhoneNo)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

//Updating existing student information
func UpdateStudent(db *sql.DB, s Student) {
	query := fmt.Sprintf("UPDATE Students SET SName='%s', DOB='%s', Address='%s', PhoneNo='%s' WHERE StudentID='%s'",
		s.SName, s.DOB, s.Address, s.PhoneNo, s.StudentID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Updating existing student information
func ViewTutor(db *sql.DB, s Student) {
	query := fmt.Sprintf("SELECT * FROM Students WHERE StudentID='%s'",
		s.StudentID, s.SName, s.DOB, s.Address, s.PhoneNo)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteStudent(db *sql.DB, s Student) {
	query := fmt.Sprintf("DELETE FROM Students WHERE StudentID='%s'",
		s.StudentID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}
func ListStudents(db *sql.DB, s Student) {
	query := fmt.Sprintf("SELECT * FROM Students",
		s.StudentID, s.SName, s.DOB, s.Address, s.PhoneNo)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

//Searches students based on personal information
func SearchStudents(db *sql.DB, s Student) {
	query := fmt.Sprintf("SELECT * FROM Students WHERE Name='%s'",
		s.StudentID, s.SName, s.DOB, s.Address, s.PhoneNo)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
////																								////
////									Functions for HTTP											////
////																								////
////////////////////////////////////////////////////////////////////////////////////////////////////////
func student(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-type") == "application/json" {
		// POST is for creating new student
		if r.Method == "POST" {
			// read the string sent to the service
			var NewStudent Student
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				// convert JSON to object
				json.Unmarshal(reqBody, &NewStudent)
				//Check if user fill up the required information for registering student's account
				if NewStudent.StudentID == "" || NewStudent.SName == "" || NewStudent.DOB == "" || NewStudent.Address == "" || NewStudent.PhoneNo == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply student " + "information " + "in JSON format"))
					return
				} else {
					CreateNewStudent(db, NewStudent) //Once everything is checked, student's account will be created
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Successfully created Student's account"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply student information " +
					"in JSON format"))
			}
		}
		//---PUT is for creating or updating existing student---
		if r.Method == "PUT" {
			queryParams := r.URL.Query() //used to resolve the conflict of calling API using the '%s'?StudentID='%s' method
			StudentID = queryParams["StudentID"][0]
			reqBody, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(reqBody, &students)
				//Check if user fill up the required information for updating Passenger's account information
				if students.StudentID == "" || students.SName == "" || students.DOB == "" || students.Address == "" || students.PhoneNo == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte("422 - Please supply student " + " information " + "in JSON format"))
				} else {
					students.StudentID = StudentID
					UpdateStudent(db, students)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte("202 - Successfully updated student's information"))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply " + "student information " + "in JSON format"))
			}
		}

	}
	if r.Method == "GET" {

		if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}

	}
	//---Deny any deletion of student's account or other student's information
	if r.Method == "DELETE" {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("403 - For audit purposes, student's account cannot be deleted."))
	}
}

func liststudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}
	}
}

func searchstudents(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if _, ok := students[params["StudentID"]]; ok {
			json.NewEncoder(w).Encode(
				students[params["StudentID"]])
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No student found"))
		}

	}
}

func main() {
	// instantiate students
	ridesharing_db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")

	db = ridesharing_db
	// handle error
	if err != nil {
		panic(err.Error())
	}
	//handle the API connection for students
	router := mux.NewRouter()
	router.HandleFunc("/students", student).Methods(
		"GET", "POST", "PUT", "DELETE")
	router.HandleFunc("/students/liststudents", liststudent).Methods(
		"GET")
	router.HandleFunc("/tutors/searchtutors", searchstudents).Methods(
		"GET")
	fmt.Println("Tutors microservice API --> Listening at port 5001")
	log.Fatal(http.ListenAndServe(":5001", router))

	defer db.Close()
}
