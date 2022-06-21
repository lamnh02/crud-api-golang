package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "io/ioutil"
)

type Task struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
}


var Tasks []Task

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllTasks(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint: returnAllTasks")
    json.NewEncoder(w).Encode(Tasks)
}

func returnSingleTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    
    for _, task := range Tasks {
        if task.Id == key {
            json.NewEncoder(w).Encode(task)
			
        }
    }
}

func createNewTask(w http.ResponseWriter, r *http.Request) {    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var task Task 
    json.Unmarshal(reqBody, &task)
    Tasks = append(Tasks, task)
    json.NewEncoder(w).Encode(task)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, task := range Tasks {
        if task.Id == id {
            Tasks = append(Tasks[:index], Tasks[index+1:]...)
        }
    }
}

func updateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    body, _ := ioutil.ReadAll(r.Body)

    var taskUpdated Task
	json.Unmarshal(body, &taskUpdated)

	taskUpdated.Id = id

	for index, task := range Tasks {
		if task.Id == id {
            fmt.Println("Endpoint: updated")
			Tasks[index] = taskUpdated
            json.NewEncoder(w).Encode("updated") 
			return
        }
	}

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/tasks", returnAllTasks).Methods("GET")
    myRouter.HandleFunc("/task", createNewTask).Methods("POST")
    myRouter.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
    myRouter.HandleFunc("/task/{id}", returnSingleTask).Methods("GET")
    myRouter.HandleFunc("/task/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    Tasks = []Task{
        Task{Id: "1", Title: "Do homework", Desc: "not responding"},
        Task{Id: "2", Title: "Play sthing", Desc: "test dsc"},
    }
    handleRequests()
}

