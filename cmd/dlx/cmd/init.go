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
	"os"
	"os/exec"
	"strings"

	"github.com/bketelsen/dlxweb/server/config"
	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Run initialization wizard",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Output("\n\nWelcome to the dlx init wizard!")

		log.Output("\n\nThis wizard will configure your server for dlx usage.")
		log.Output("If you don't yet have the prerequisites installed, we'll get them installed now.\n\n")

		log.Output("This wizard assumes Ubuntu Server 18.04 or higher.")

		log.Output("Here's what we'll install:")

		log.Output("\t * lxd - container management daemon")
		log.Output("\t * distrobuilder - container image builder")
		log.Output("\t * debootstrap - linux filesystem installer")

		log.Output("\nAfter installing the prerequisites we'll ask some questions, create")
		log.Output("a configuration file with your settings, and then build your base")
		log.Output("container image.")
		log.Output("\nShall we continue? [y/n]")

		if !askYN() {
			log.Output("You can re-run this wizard later to continue.")
			os.Exit(0)
		}
		log.Output("Let's get started!")
		log.Output("\n*Installing lxd*")
		err := installIt("sudo", "snap", "install", "lxd", "--classic")
		if err != nil {
			log.Error(err.Error())
			//os.Exit(1)
		}
		log.Output("\n*Installing debootstrap*")
		err = installIt("sudo", "apt", "install", "debootstrap", "-y")
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		log.Output("\n*Installing distrobuilder*")
		err = installIt("sudo", "snap", "install", "distrobuilder", "--classic")
		if err != nil {
			log.Error(err.Error())
			//os.Exit(1)
		}

		_, err = config.Get()
		if err != nil {
			log.Info("Creating new config file")
			err = config.Create()
			if err != nil {
				log.Error(err.Error())
				os.Exit(1)
			}
		} else {
			log.Warn("Configuration file already exists")
		}
		log.Info("The configuration file is at " + config.GetConfigPath())

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func askYN() bool {
	response, err := log.Ask("", "")
	if err != nil {
		log.Output("error: " + err.Error())
	}
	if !strings.Contains(strings.ToLower(response), "y") {
		return false
	}
	return true
}

func installIt(commmand string, args ...string) error {
	cmd := exec.Command(commmand, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
