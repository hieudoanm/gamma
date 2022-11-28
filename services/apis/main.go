package main

import (
	"fmt"
	"log"
	"net/http"

	health_controller "apis/src/controller/health"
	tasks_controller "apis/src/controller/tasks"
	"apis/src/utils"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// Router
	router.GET("/health", health_controller.GetHealth)
	// Tasks
	router.GET("/tasks", tasks_controller.GetTasks)
	router.POST("/tasks", tasks_controller.CreateTask)
	router.GET("/tasks/:id", tasks_controller.GetTaskById)
	router.PUT("/tasks/:id", tasks_controller.UpdateTaskById)
	router.PATCH("/tasks/:id", tasks_controller.PatchTaskById)
	router.DELETE("/tasks/:id", tasks_controller.DeleteTaskById)
	// Start
	var PORT string = utils.Getenv("PORT", "8080")
	log.Printf("🚀 Server is listening on port %s", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))

}
