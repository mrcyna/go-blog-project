package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/mrcyna/blog/db"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit the blog post within CLI",
	Long:  "Edit the blog post within CLI",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Enter the ID:")
		var id int
		fmt.Scanf("%d\n", &id)

		find, post := db.GetPost(id)

		if !find {
			fmt.Printf("Unable to find post with id %d \n", id)
			os.Exit(0)
		}

		reader := bufio.NewReader(os.Stdin)

		// Title
		fmt.Printf("Enter Title: (Default is '%s' if left empty)\n", post.Title)
		title, _ := reader.ReadString('\n')
		if title == "\n" {
			title = post.Title
		}

		// Body
		fmt.Printf("Enter Body: (Default is '%s' if left empty)\n", post.Body)
		body, _ := reader.ReadString('\n')
		if body == "\n" {
			body = post.Body
		}

		// Author
		fmt.Printf("Enter Author: (Default is '%s' if left empty)\n", post.Author)
		author, _ := reader.ReadString('\n')
		if author == "\n" {
			author = post.Author
		}

		// Published
		fmt.Printf("Enter Published: (Default is '%t' if left empty) Y/n\n", post.Published)
		publishedInput, _ := reader.ReadString('\n')
		var published bool
		if publishedInput == "\n" {
			published = post.Published
		} else if publishedInput == "y\n" || publishedInput == "Y\n" {
			published = true
		} else {
			published = false
		}

		r := db.UpdatePost(id, title, body, author, published)

		if r {
			fmt.Println("Post updated successfully!")
		} else {
			fmt.Println("Unable to update the post")
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
