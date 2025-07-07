package todo_test

import (
	"encoding/json"
	"testing"
	"todo"
)

func TestNewTask(t *testing.T) {
	// Arrange
	id := 1
	title := "Test Task"

	// Act
	task := todo.NewTask(id, title)

	// Assert
	if task.ID != id {
		t.Errorf("Expected task ID to be %d, got %d", id, task.ID)
	}
	if task.Title != title {
		t.Errorf("Expected task title to be %s, got %s", title, task.Title)
	}
	if task.Completed != false {
		t.Errorf("Expected task completion to be false, got %v", task.Completed)
	}
}

func TestTaskAsJSON(t *testing.T) {
	// Arrange
	task := todo.NewTask(1, "Test Task")

	// Act
	jsonStr, err := task.AsJSON()

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Проверяем, что JSON можно декодировать обратно в Task
	var decodedTask todo.Task
	err = json.Unmarshal([]byte(jsonStr), &decodedTask)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if decodedTask.ID != task.ID {
		t.Errorf("Expected task ID to be %d, got %d", task.ID, decodedTask.ID)
	}
	if decodedTask.Title != task.Title {
		t.Errorf("Expected task title to be %s, got %s", task.Title, decodedTask.Title)
	}
	if decodedTask.Completed != task.Completed {
		t.Errorf("Expected task completion to be %v, got %v", task.Completed, decodedTask.Completed)
	}

}

func TestDeleteTask(t *testing.T) {
	// Arrange
	taskList := todo.NewTaskList()
	task := todo.NewTask(1, "Test Task")
	taskList.Add(task)
	// Act
	taskList.Remove(task.ID)
	// Assert
	if len(taskList.GetAll()) != 0 {
		t.Errorf("Expected task list to be empty, got %d tasks", len(taskList.GetAll()))
	}
}
