package godoit

import (
	"fmt"
	"os"

	"github.com/gretarst/godoit/pkg/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"list"},
	Short:   "list all todos",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.NewConnection()
		if err != nil {
			fmt.Printf("Failed to connect to database")
			os.Exit(1)
		}

		t, err := db.List()
		if err != nil {
			fmt.Printf("Failed to get todos from database")
			os.Exit(1)
		}

		for _, todo := range t {
			fmt.Printf("%d - %s \n", todo.ID, todo.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
