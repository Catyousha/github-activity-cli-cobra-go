package cmd

import (
	"fmt"
	"tenessine/github-activity/internal"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "github-activity [username]",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := internal.GetGithubActivity(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
		},
	}
}
