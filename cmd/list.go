package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/mrcyna/blog/db"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display list of blog posts",
	Long:  "Display list of blog posts in CLI",
	Run: func(cmd *cobra.Command, args []string) {
		posts := db.AllPosts()

		if len(posts) == 0 {
			fmt.Println("No entry!")
			os.Exit(0)
		}

		for i := 0; i < len(posts); i++ {

			published := fmt.Sprintf("%t", posts[i].Published)

			fmt.Println(`
ID:		` + posts[i].ID + `
Title: 		` + posts[i].Title + `
Slug:		` + posts[i].Slug + `
Body:		` + posts[i].Body + `
Author:		` + posts[i].Author + `
Published:	` + published + `
			`)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
