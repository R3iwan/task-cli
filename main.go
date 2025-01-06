package main

import "time"

type Task struct {
	id          int       `json:"id"`
	description string    `json:"description"`
	status      string    `json:"string"`
	createdAt   time.Time `json:"created_at"`
	updatedAt   time.Time `json:"updated_at"`
}

var tasks []Task

func main() {

}

func addTask() {

}
