package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bketelsen/dlxweb/generated/client"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm", "d"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		instanceService := client.NewInstanceService(cl)
		name := args[0]
		if name == "" {
			log.Error("<container name> required argument is missing")
			os.Exit(1)
		}
		_, err := instanceService.Delete(context.Background(), client.InstanceDeleteRequest{
			Project: "",
			Name:    name,
		})
		if err != nil {
			fmt.Println("remote error:", err)
			return
		}
		log.Info("Instance deleted")
	},
}

func init() {
	instanceCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
