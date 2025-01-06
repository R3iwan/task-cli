# Task Tracker

A simple command-line task management tool implemented in Go. This tool allows you to create, update, delete, and manage tasks efficiently. Each task is stored in a JSON file, making it easy to persist data across sessions.
URL : https://roadmap.sh/projects/task-tracker

## Features

- Add new tasks with descriptions.
- Update the description of an existing task.
- Delete tasks by ID.
- Mark tasks as:
  - `in-progress`
  - `done`
  - `to-do`
- List all tasks or filter by status (`list <status>`).
- Data persistence in a JSON file (`tasks.json`).

## Usage

### Running the Application

1. Clone this repository.
2. Navigate to the project directory.
3. Run the application using:
   ```bash
   go run main.go
