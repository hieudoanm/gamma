package tasks_service

import (
	"encoding/json"
	"fmt"
	"fullstack-graphql/src/libs/http"
	"fullstack-graphql/src/utils"
	"log"
	"time"
)

var API = utils.Getenv("API", "http://localhost:8080")

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Createdat time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
}

func GetTasks() ([]Task, error) {
	var url = fmt.Sprintf("%s/tasks", API)
	body, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetTasks")
		return nil, httpGetError
	}
	// Parse JSON
	var tasks []Task
	jsonUnmarshalError := json.Unmarshal(body, &tasks)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetTasks")
		return nil, jsonUnmarshalError
	}

	return tasks, nil
}

func GetTask(id string) (Task, error) {
	var url = fmt.Sprintf("%s/tasks/%s", API, id)
	body, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetTasks")
		return Task{}, httpGetError
	}
	// Parse JSON
	var task Task
	jsonUnmarshalError := json.Unmarshal(body, &task)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetTasks")
		return Task{}, jsonUnmarshalError
	}

	return task, nil
}
