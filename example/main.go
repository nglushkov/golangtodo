package main

import (
	"fmt"
	"log"

	"todo"
)

func main() {

	taskList := todo.NewTaskList()

	task1 := todo.NewTask(1, "Learn Go")
	task2 := todo.NewTask(2, "Build a web app")

	taskList.Add(task1)
	taskList.Add(task2)
	taskList.Remove(task1.ID)
	taskList.UpdateTask(task2.ID, "Build a web application", true)

	json, err := taskList.GetAsJSON()
	if err != nil {
		log.Fatalf("Error getting JSON: %v", err)
	}
	fmt.Println(json)

	taskFromList := taskList.Get(task2.ID)
	if taskFromList != nil {
		fmt.Printf("Task: %s, Completed: %v\n", taskFromList.Title, taskFromList.Completed)
	}

	taskList.Clear()

	json, err = taskList.GetAsJSON()
	if err != nil {
		log.Fatalf("Error getting JSON: %v", err)
	}
	fmt.Println(json)
}
