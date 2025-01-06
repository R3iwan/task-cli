package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

var tasks []Task
var nextID int = 1

func main() {
	for {
		var input string
		fmt.Scan(&input)

		if strings.HasPrefix(input, "list ") {
			status := strings.TrimPrefix(input, "list ")
			listStatus(status)
			continue
		}

		switch input {
		case "add":
			addTask()
		case "update":
			updateTask()
		case "delete":
			deleteTask()
		case "mark-in-progress":
			markInProgress()
		case "mark-done":
			markDone()
		case "mark-to-do":
			markToDo()
		case "list":
			list()
		case "exit":
			break
		}
	}
}

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	desc, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading the description:", err)
		return
	}

	desc = strings.TrimSpace(desc)
	var currentTime = time.Now()
	var formatTime = currentTime.Format("2006-01-02 15:04:05")

	task := Task{
		Id:          nextID,
		Description: desc,
		Status:      "in-progress",
		CreatedAt:   formatTime,
		UpdatedAt:   formatTime,
	}

	nextID++
	tasks = append(tasks, task)

	fmt.Println("Task added successfully!")
	fmt.Printf("Task Details: %+v\n", task)

	err = saveTasksToJSON("tasks.json")
	if err != nil {
		fmt.Println("Error saving tasks to file")
	}
}

func updateTask() {
	var taskID int
	var desc string

	fmt.Scan(&taskID)

	reader := bufio.NewReader(os.Stdin)
	desc, _ = reader.ReadString('\n')
	desc = strings.TrimSpace(desc)

	for i, task := range tasks {
		if task.Id == taskID {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			fmt.Println("Task updated successfully!")
			break
		}
	}

	err := saveTasksToJSON("tasks.json")
	if err != nil {
		fmt.Println("Error updating task")
	}
}

func deleteTask() {
	var taskID int
	fmt.Scan(&taskID)

	for i, task := range tasks {
		if task.Id == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			break
		}
	}

	err := saveTasksToJSON("tasks.json")
	if err != nil {
		fmt.Println("Error deleting a task")
	}
}

func markInProgress() {
	var taskID int
	fmt.Scan(&taskID)

	for i, task := range tasks {
		if task.Id == taskID {
			tasks[i].Status = "in-progress"
			fmt.Println("Task status changed successfully!")
			break
		}
	}

	err := saveTasksToJSON("tasks.json")
	if err != nil {
		fmt.Println("Error changing a status of a task")
	}
}

func markDone() {
	var taskID int
	fmt.Scan(&taskID)

	for i, task := range tasks {
		if task.Id == taskID {
			tasks[i].Status = "done"
			fmt.Println("Task status changed successfully!")
			break
		}
	}

	err := saveTasksToJSON("tasks.json")
	if err != nil {
		fmt.Println("Error changing a status of a task")
	}
}

func markToDo() {
	var taskID int
	fmt.Scan(&taskID)

	for i, task := range tasks {
		if task.Id == taskID {
			tasks[i].Status = "to-do"
			fmt.Println("Task status changed successfully!")
			break
		}
	}

	err := saveTasksToJSON("tasks.json")
	if err != nil {
		fmt.Println("Error changing a status of a task")
	}
}

func list() {
	if len(tasks) == 0 {
		fmt.Println("Tasks is empty")
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n", task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

func listStatus(status string) {
	status = strings.ToLower(strings.TrimSpace(status))

	track := false
	for _, task := range tasks {
		if strings.ToLower(task.Status) == status {
			fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated At: %s\nUpdated At: %s\n\n",
				task.Id, strings.TrimSpace(task.Description), task.Status, task.CreatedAt, task.UpdatedAt)
			track = true
		}
	}

	if !track {
		fmt.Printf("No tasks with status '%s' found.\n", status)
	}
}

func saveTasksToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating a file")
	}
	defer file.Close()

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Couldn't convert json")
	}

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Couldn't write to a file")
	}

	return nil
}
