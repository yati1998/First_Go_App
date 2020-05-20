package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)
// struct to store the student info
type Student struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Subject string `json:"subject"`
	Score string `json:"score"`
}

var Students []Student
//To print all the details of student
func allDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All details endpoint")
	json.NewEncoder(w).Encode(Students)
}
//return student details according to the id
func returnSingleDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, student := range Students {
		if student.Id == key {
			json.NewEncoder(w).Encode(student)
		}
	}
}
//Function to create new details and append to the student struct
func createNewDetails(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request, unmarshal it into new student struct
	// append this to our Students array

	reqBody, _ := ioutil.ReadAll(r.Body)
	var student Student
	json.Unmarshal(reqBody, &student)
	Students = append(Students, student)

	json.NewEncoder(w).Encode(student)

}
//Function to delete the info according to the id
func deleteStudent(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	for index, student := range Students{
		if student.Id == id {
			Students = append(Students[:index],Students[index+1:]...)
		}
		}
	}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to Home Page!")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/details", allDetails)
	myRouter.HandleFunc("/create-student", createNewDetails).Methods("POST")
	myRouter.HandleFunc("/student/{id}", deleteStudent).Methods("DELETE")
	myRouter.HandleFunc("/student/{id}", returnSingleDetail)
	log.Fatal(http.ListenAndServe(":5000",myRouter))
}

func main() {
	Students = []Student{
		Student{Id: "1", Name: "Yati", Subject: "Maths", Score: "99"},
		Student{Id: "2", Name: "Aditi", Subject: "English", Score: "95"},
	}
	handleRequests()
}


