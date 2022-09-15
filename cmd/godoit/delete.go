package godoit

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gretarst/godoit/pkg/todos"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"delete"},
	Short:   "delete a todo",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tdb, err := todos.NewTodos()
		if err != nil {
			fmt.Printf("Failed to connect to database")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Printf("Argument needs to be an integer type")
			os.Exit(1)
		}

		id, err = tdb.Delete(id)

		if err != nil {
			fmt.Printf("Failed to delete %d from database", id)
			os.Exit(1)
		}

		fmt.Printf("Deleted todo with id %d \n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
