package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = slices.Delete(t, index, index+1)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed

	if !isCompleted {
		completedTime := time.Now()
		t[index].CompletedAt = &completedTime
	}
	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")

	for i, todo := range *todos {
		completed := "No"
		completedAt := ""

		if todo.Completed {
			completed = "Yes"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
