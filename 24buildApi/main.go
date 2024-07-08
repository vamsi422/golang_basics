package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db

var courses []Course

// middleware,helper -file
func (c *Course) IsEmpty() bool {

	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("welcome to demo-API")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Vamshi Krishna", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{FullName: "Vamshi Krishna", Website: "go.dev"}})

	// routing

	r.HandleFunc("/", serverHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers -file

// serve home route

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))

}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one course.")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)
	//fmt.Println(params)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return

}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course.")
	w.Header().Set("Content-Type", "application/json")

	// what if entire body is empty.

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}
	// what about -{}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id,string
	// append course into courses

	//rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("update one course.")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req

	params := mux.Vars(r)

	// loop,id ,remove,add with ID

	for index, course := range courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)

			fmt.Printf("course.CourseId is : %v  params['index'] is : %v", course.CourseId, params["index"])
			course.CourseId = params["index"]

			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return

		}
		json.NewEncoder(w).Encode("id not found.")
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("delete one course.")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	response := "Id not found."

	for index, couse := range courses {
		if couse.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			response = fmt.Sprintf("value deleted successfully.", index)
			break
		}
	}
	json.NewEncoder(w).Encode(response)
	return
}
