package todo_test

import (
	"encoding/json"
	"testing"
	"todo"
)

func TestNewTaskList(t *testing.T) {
	// Act
	list := todo.NewTaskList()

	// Assert
	if list == nil {
		t.Error("Expected non-nil TaskList")
	}
	if len(list.Tasks) != 0 {
		t.Errorf("Expected empty task list, got %d tasks", len(list.Tasks))
	}
}

func TestTaskList_Add(t *testing.T) {
	// Arrange
	list := todo.NewTaskList()
	task := todo.NewTask(1, "Test Task")

	// Act
	list.Add(task)

	// Assert
	if len(list.Tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(list.Tasks))
	}
	if list.Tasks[0].ID != task.ID {
		t.Errorf("Expected task ID to be %d, got %d", task.ID, list.Tasks[0].ID)
	}
}

func TestTaskList_Get(t *testing.T) {
	// Arrange
	list := todo.NewTaskList()
	task1 := todo.NewTask(1, "Task 1")
	task2 := todo.NewTask(2, "Task 2")
	list.Add(task1)
	list.Add(task2)

	// Act & Assert
	// Тест 1: Поиск существующей задачи
	foundTask := list.Get(2)
	if foundTask == nil {
		t.Error("Expected to find task with ID 2, got nil")
	} else if foundTask.ID != 2 {
		t.Errorf("Expected task ID to be 2, got %d", foundTask.ID)
	}

	// Тест 2: Поиск несуществующей задачи
	notFoundTask := list.Get(3)
	if notFoundTask != nil {
		t.Errorf("Expected nil for non-existent task, got %+v", notFoundTask)
	}
}

func TestTaskList_Remove(t *testing.T) {
	// Arrange
	list := todo.NewTaskList()
	task1 := todo.NewTask(1, "Task 1")
	task2 := todo.NewTask(2, "Task 2")
	list.Add(task1)
	list.Add(task2)

	// Act
	list.Remove(1)

	// Assert
	if len(list.Tasks) != 1 {
		t.Errorf("Expected 1 task after removal, got %d", len(list.Tasks))
	}
	if list.Tasks[0].ID != 2 {
		t.Errorf("Expected remaining task ID to be 2, got %d", list.Tasks[0].ID)
	}

	// Проверка удаления несуществующего ID (не должно вызывать ошибок)
	list.Remove(3) // Не существующий ID
	if len(list.Tasks) != 1 {
		t.Errorf("Expected still 1 task, got %d", len(list.Tasks))
	}
}

func TestTaskList_UpdateTask(t *testing.T) {
	// Arrange
	list := todo.NewTaskList()
	task := todo.NewTask(1, "Original Title")
	list.Add(task)

	// Act
	newTitle := "Updated Title"
	newCompleted := true
	list.UpdateTask(1, newTitle, newCompleted)

	// Assert
	updatedTask := list.Get(1)
	if updatedTask == nil {
		t.Fatal("Task not found after update")
	}
	if updatedTask.Title != newTitle {
		t.Errorf("Expected title to be %s, got %s", newTitle, updatedTask.Title)
	}
	if updatedTask.Completed != newCompleted {
		t.Errorf("Expected completed to be %v, got %v", newCompleted, updatedTask.Completed)
	}

	// Проверка обновления несуществующего ID (не должно вызывать ошибок)
	list.UpdateTask(999, "Non-existent", false)
	// Состояние не должно измениться
	if len(list.Tasks) != 1 {
		t.Errorf("Expected still 1 task, got %d", len(list.Tasks))
	}
}

func TestTaskList_Clear(t *testing.T) {
	// Arrange
	list := todo.NewTaskList()
	list.Add(todo.NewTask(1, "Task 1"))
	list.Add(todo.NewTask(2, "Task 2"))

	// Act
	list.Clear()

	// Assert
	if len(list.Tasks) != 0 {
		t.Errorf("Expected 0 tasks after clear, got %d", len(list.Tasks))
	}
}

func TestTaskList_GetAll(t *testing.T) {
	// Arrange
	list := todo.NewTaskList()
	task1 := todo.NewTask(1, "Task 1")
	task2 := todo.NewTask(2, "Task 2")
	list.Add(task1)
	list.Add(task2)

	// Act
	tasks := list.GetAll()

	// Assert
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestTaskList_GetAsJSON(t *testing.T) {
	// Arrange
	list := todo.NewTaskList()
	list.Add(todo.NewTask(1, "Task 1"))
	list.Add(todo.NewTask(2, "Task 2"))

	// Act
	jsonStr, err := list.GetAsJSON()

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Проверяем, что JSON можно декодировать обратно в TaskList
	var decodedList todo.TaskList
	err = json.Unmarshal([]byte(jsonStr), &decodedList)
	if err != nil {
		t.Errorf("Failed to unmarshal JSON: %v", err)
	}

	if len(decodedList.Tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(decodedList.Tasks))
	}
}
