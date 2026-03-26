package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TaskList struct {
	tasks []Task
}

func (tl *TaskList) AddTask(task Task) {
	id := len(tl.tasks) + 1
	newTask := Task{ID: id, Title: task.Title, Done: false}
	tl.tasks = append(tl.tasks, newTask)
}

func (tl *TaskList) ShowTasks() {
	for _, task := range tl.tasks {
		status := "Not Done"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%d. %s - %s\n", task.ID, task.Title, status)
	}
}

func (tl *TaskList) CompleteTask(id int) {
	for i, task := range tl.tasks {
		if task.ID == id {
			tl.tasks[i].Done = true
			break
		}
	}
}

func (tl *TaskList) DeleteTask(id int) {
	var newTasks []Task
	for _, task := range tl.tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	tl.tasks = newTasks
}

func main() {
	taskList := TaskList{}

	for {
		fmt.Println("\n1. Добавить Задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Выполнить задачу")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выход")

		var choice int
		fmt.Print("Выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var title string
			fmt.Print("Введите название задачи: ")
			fmt.Scan(&title)
			taskList.AddTask(Task{Title: title})

		case 2:
			taskList.ShowTasks()
		case 3:
			var id int
			fmt.Print("Введите номер задачи: ")
			fmt.Scan(&id)
			taskList.CompleteTask(id)
		case 4:
			var id int
			fmt.Print("Введите номер задачи: ")
			fmt.Scan(&id)
			taskList.DeleteTask(id)
		case 5:
			return
		default:
			fmt.Println("Неверный выбор")
		}
	}
}
