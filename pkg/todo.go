package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const Todos_file = "./todos.json"

type Todo struct {
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type Todos []Todo

func (t *Todos) ValidateIndex(id int) error {
	idx := id - 1

	if idx < 0 || idx >= len(*t) {
		return errors.New("Todo not found")
	}
	return nil
}

func (t *Todos) Save() error {

	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return errors.New("Error marshalling todos:" + err.Error())
	}

	err = os.WriteFile(Todos_file, data, 0644)
	if err != nil {
		return errors.New("Error writing to file:" + err.Error())
	}

	return nil
}

func (t *Todos) Load() error {
	data, err := os.ReadFile(Todos_file)
	if err != nil {
		return errors.New("Error reading file:" + err.Error())
	}

	if len(data) == 0 {
		return nil
	}

	err = json.Unmarshal(data, t)
	if err != nil {
		return errors.New("Error unmarshalling todos:" + err.Error())
	}

	return nil
}

func (t *Todos) List() Todos {
	for idx, todo := range *t {
		fmt.Printf("ID: %d, Title: %s, Completed: %t, CreatedAt: %s, CompletedAt: %s\n",
			idx+1, todo.Title, todo.Completed, todo.CreatedAt, todo.CompletedAt)
	}

	return (*t)
}

func (t *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*t = append(*t, todo)
	t.Save()
}

func (t *Todos) GetOne(id int) (Todo, error) {
	idx := id - 1

	if err := t.ValidateIndex(id); err != nil {
		return Todo{}, err
	}

	todo := (*t)[idx]
	fmt.Printf("Title: %s, Completed: %t, CreatedAt: %s, CompletedAt: %s\n",
		todo.Title, todo.Completed, todo.CreatedAt, todo.CompletedAt)

	return (*t)[idx], nil
}

func (t *Todos) Edit(id int, title string) error {

	if err := t.ValidateIndex(id); err != nil {
		return err
	}

	(*t)[id-1].Title = title

	t.Save()

	return nil
}

func (t *Todos) Delete(id int) error {
	if err := t.ValidateIndex(id); err != nil {
		return err
	}

	*t = append((*t)[:id-1], (*t)[id:]...)

	t.Save()

	return nil
}

func (t *Todos) ToggleComplete(id int) error {
	idx := id - 1

	if err := t.ValidateIndex(id); err != nil {
		return err
	}

	isComplete := (*t)[idx].Completed

	if !isComplete {
		completionDate := time.Now()
		(*t)[idx].CompletedAt = &completionDate
	} else {
		(*t)[idx].CompletedAt = nil
	}

	(*t)[idx].Completed = !isComplete

	t.Save()
	return nil
}
