package cmd

import (
	"context"
	"fmt"

	"github.com/bketelsen/dlxweb/generated/client"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		imageService := client.NewImageService(cl)
		resp, err := imageService.Build(context.Background(), client.ImageBuildRequest{
			Project: "",
		})
		if err != nil {
			fmt.Println("remote error:", err)
			return
		}
		fmt.Println(resp)
	},
}

func init() {
	imageCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
