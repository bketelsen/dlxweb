/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/bketelsen/dlxweb/generated/client"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		instanceService := client.NewInstanceService(cl)

		resp, err := instanceService.List(context.Background(), client.InstanceListRequest{
			Project: "",
		})
		if err != nil {
			fmt.Println("remote error:", err)
			return
		}

		header := []string{"Name", "Status", "IPv4"}
		var data [][]string
		for _, instance := range resp.Instances {
			data = append(data, []string{instance.Name, instance.Status, instance.IPV4})
		}

		if len(resp.Instances) < 1 {

			data = append(data, []string{"{None Found}", "", ""})
		}
		RenderTable(TableFormatTable, header, data, nil)

	},
}

func init() {
	instanceCmd.AddCommand(listCmd)
}
