package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

var tasks []Task
var lastID int

const tasksFile = "tasks.json"

func main() {
	loadTasks()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Update Task Status")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			addTask(scanner)
		case "2":
			listTasks()
		case "3":
			deleteTask(scanner)
		case "4":
			updateTaskStatus(scanner)
		case "5":
			saveTasks()
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := scanner.Text()
	lastID++
	task := Task{ID: lastID, Title: title, Done: false}
	tasks = append(tasks, task)
	fmt.Println("Task added.")
	saveTasks()
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Title, status)
	}
}

func deleteTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to delete: ")
	scanner.Scan()
	var id int
	fmt.Sscanf(scanner.Text(), "%d", &id)
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted.")
			saveTasks()
			return
		}
	}
	fmt.Println("Task not found.")
}

func updateTaskStatus(scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to update: ")
	scanner.Scan()
	var id int
	fmt.Sscanf(scanner.Text(), "%d", &id)
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = !tasks[i].Done
			fmt.Println("Task status updated.")
			saveTasks()
			return
		}
	}
	fmt.Println("Task not found.")
}

func loadTasks() {
	file, err := os.Open(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error loading tasks:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		fmt.Println("Error decoding tasks:", err)
		return
	}

	for _, task := range tasks {
		if task.ID > lastID {
			lastID = task.ID
		}
	}
}

func saveTasks() {
	file, err := os.Create(tasksFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
	}
}
