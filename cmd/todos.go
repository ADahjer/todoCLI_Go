package cmd

import (
	"errors"
	"os"
	"strconv"
	"todos/pkg"

	"github.com/spf13/cobra"
)

var todos = pkg.Todos{}

func init() {
	if _, err := os.Stat(pkg.Todos_file); err == nil {
		todos.Load()
	}

	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntP("index", "i", 0, "Return the todo with the specified ID")

	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("title", "t", "", "Title of the new todo")
	addCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringP("title", "t", "", "New todo title")
	editCmd.MarkFlagRequired("title")

	rootCmd.AddCommand(toggleCmd)

	rootCmd.AddCommand(deleteCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List all the todos",
	RunE:    listTodos,
}

var addCmd = &cobra.Command{
	Use:     "add [--title|-t] <todo_title>",
	Aliases: []string{"a"},
	Short:   "Add a new todo",
	Run:     addTodo,
}

var editCmd = &cobra.Command{
	Use:          "edit <todo_id> [--title|-t] <new_todo_title>",
	Short:        "Edit a todo",
	Aliases:      []string{"e"},
	Args:         cobra.ExactArgs(1),
	RunE:         editTodo,
	SilenceUsage: true,
}

var toggleCmd = &cobra.Command{
	Use:          "toggle <todo_id>",
	Short:        "Toggle the status of completion of a todo",
	Args:         cobra.ExactArgs(1),
	RunE:         toggleTodo,
	SilenceUsage: true,
}

var deleteCmd = &cobra.Command{
	Use:          "delete <todo_id>",
	Aliases:      []string{"rm", "del"},
	Short:        "Delete the todo based on the given id",
	Args:         cobra.ExactArgs(1),
	RunE:         deleteTodo,
	SilenceUsage: true,
}

func listTodos(cmd *cobra.Command, args []string) error {
	id, _ := cmd.Flags().GetInt("index")

	err := todos.ValidateIndex(id)

	if id != 0 && err != nil {
		return err
	} else if err == nil {
		todos.GetOne(id)
		return nil
	}

	todos.List()

	return nil
}

func addTodo(cmd *cobra.Command, args []string) {
	title, _ := cmd.Flags().GetString("title")

	todos.Add(title)
}

func editTodo(cmd *cobra.Command, args []string) error {
	title, _ := cmd.Flags().GetString("title")
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("you should pass the id of the todo as a number")
	}

	if title == "" {
		return errors.New("the title cannot be empty")
	}

	if err := todos.ValidateIndex(id); err != nil {
		return err
	}

	todos.Edit(id, title)

	return nil
}

func toggleTodo(cmd *cobra.Command, args []string) error {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("you should pass the id of the todo as a number")
	}

	if err := todos.ValidateIndex(id); err != nil {
		return err
	}

	todos.ToggleComplete(id)

	return nil
}

func deleteTodo(cmd *cobra.Command, args []string) error {
	id, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.New("you should pass the id of the todo as a number")
	}

	if err := todos.ValidateIndex(id); err != nil {
		return err
	}

	todos.Delete(id)

	return nil
}
