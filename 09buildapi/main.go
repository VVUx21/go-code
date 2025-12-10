package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)

type Course struct {
	Id    string `json:"courseid"`
	Name  string `json:"coursename"`
	Price int    `json:"price"`
	Author Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.Name == "" && c.Id == "" && c.Price == 0
}

func main() {
	fmt.Println("welcome to build api")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{Id: "1", Name: "ReactJS", Price: 299, Author: Author{Fullname: "John Doe", Website: "johndoe.com"}})
	courses = append(courses, Course{Id: "2", Name: "MERN Stack", Price: 199, Author: Author{Fullname: "Jane Doe", Website: "janedoe.com"}})
	courses = append(courses, Course{Id: "3", Name: "Angular", Price: 299, Author: Author{Fullname: "Jim Beam", Website: "jimbeam.com"}})
	courses = append(courses, Course{Id: "4", Name: "VueJS", Price: 299, Author: Author{Fullname: "Jack Daniels", Website: "jackdaniels.com"}})
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//listen to port
	http.ListenAndServe(":4000", r)
}

// controllers - file
// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to the Home Page </h1>"))
}

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //explain- json encoder converts Go objects into JSON format and writes it to the response writer
}

// get one course
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //grab id from request
	for _, course := range courses {
		if course.Id == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with given id")
}

// create one course
// update one course
// delete one course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course) //decode the request body into course struct
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("no data inside json")
		return
	}
	//generate unique id, string
	rand.Seed(time.Now().UnixNano())
	course.Id=strconv.Itoa(rand.Intn(100)) //mock id - not safe
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop, find, remove, add with my ID
	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course //meaning new course
			_ = json.NewDecoder(r.Body).Decode(&course) //update the course
			//id should be same
			course.Id = params["id"]
			courses = append(courses, course) //add the course
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with given id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop, find, remove , return
	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //explain in detail - this line removes the course at the specified index by appending the slice of courses before the index with the slice of courses after the index, effectively deleting the course from the slice
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
}