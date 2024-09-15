package cmd

import (
	"os"
	"todoCLI/pkg"

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
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the todos",
	RunE:  listTodos,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Run:   addTodo,
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
