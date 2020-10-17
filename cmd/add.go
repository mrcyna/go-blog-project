package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/mrcyna/blog/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new blog post",
	Long:  "You can add blog post within the command line",
	Run: func(cmd *cobra.Command, args []string) {

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter Title:")
		title, _ := reader.ReadString('\n')

		fmt.Println("Enter Body:")
		body, _ := reader.ReadString('\n')

		fmt.Println("Enter Author:")
		author, _ := reader.ReadString('\n')

		fmt.Println("Publish Right Now? Y/n")
		var publishInput string
		fmt.Scanf("%s\n", &publishInput)
		var publish bool = false

		if publishInput == "Y" || publishInput == "y" {
			publish = true
		}

		r := db.CreatePost(title, body, author, publish)

		if r {
			fmt.Println("Post created successfully!")
		} else {
			fmt.Println("Unable to create the post")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
