package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bketelsen/dlxweb/generated/client"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new <name>",
	Aliases: []string{"create", "c", "n"},
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
		resp, err := instanceService.Create(context.Background(), client.InstanceCreateRequest{
			Project: "",
			Name:    name,
		})
		if err != nil {
			fmt.Println("remote error:", err)
			return
		}

		header := []string{"Name", "Status", "IPv4"}
		var data [][]string
		data = append(data, []string{resp.Instance.Name, resp.Instance.Status, resp.Instance.IPV4})

		RenderTable(TableFormatTable, header, data, nil)
	},
}

func init() {
	instanceCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
