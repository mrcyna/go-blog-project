package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/mrcyna/blog/db"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a blog post",
	Long:  "Delete a blog post within the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Enter the ID:")
		var id int
		fmt.Scanf("%d\n", &id)

		find, _ := db.GetPost(id)

		if !find {
			fmt.Printf("Unable to find post with id %d \n", id)
			os.Exit(0)
		}

		fmt.Println("Are you sure to delete? This action is not undoable! Y/n")
		var confirmation string
		fmt.Scanf("%s\n", &confirmation)

		if confirmation == "Y" || confirmation == "y" {
			r := db.DeletePost(id)

			if r {
				fmt.Println("Delete successfully")
			} else {
				fmt.Println("Unable to delete due to unknown error!")
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
