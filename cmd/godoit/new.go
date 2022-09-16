package godoit

import (
	"fmt"
	"os"
	"time"

	"github.com/gretarst/godoit/pkg/database"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"new"},
	Short:   "create a new todo",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.NewConnection()
		if err != nil {
			fmt.Printf("Failed to connect to database")
			os.Exit(1)
		}

		t := args[0]

		id, err := db.Insert(database.Todo{
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
