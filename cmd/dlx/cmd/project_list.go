package cmd

import (
	"context"
	"fmt"

	"github.com/bketelsen/dlxweb/generated/client"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var projectListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		projectService := client.NewProjectService(cl)

		resp, err := projectService.List(context.Background(), client.ProjectListRequest{})
		if err != nil {
			fmt.Println("remote error:", err)
			return
		}
		fmt.Println(resp)
	},
}

func init() {
	projectCmd.AddCommand(projectListCmd)
}
