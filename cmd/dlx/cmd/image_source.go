package cmd

import (
	"context"
	"fmt"

	"github.com/bketelsen/dlxweb/generated/client"
	"github.com/spf13/cobra"
)

// sourceCmd represents the source command
var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		imageService := client.NewImageService(cl)

		resp, err := imageService.Source(context.Background(), client.ImageSourceRequest{
			Project: "",
		})
		if err != nil {
			fmt.Println("remote error:", err)
			return
		}
		fmt.Println(resp.Source)
	},
}

func init() {
	imageCmd.AddCommand(sourceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sourceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sourceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
