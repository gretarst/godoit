package godoit

import (
	"fmt"
	"os"
	"time"

	"github.com/gretarst/godoit/pkg/todos"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"new"},
	Short:   "create a new todo",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tdb, err := todos.NewTodos()
		if err != nil {
			fmt.Printf("Failed to connect to database")
			os.Exit(1)
		}

		t := args[0]

		id, err := tdb.Insert(todos.Todo{
			Time:  time.Now(),
			Title: t,
		})

		if err != nil {
			fmt.Printf("Failed to write to database")
			os.Exit(1)
		}

		fmt.Printf("New todo with id %d \n", id)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
