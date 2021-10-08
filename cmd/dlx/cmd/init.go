package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
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
		log.Info("The default image template is in the same directory called 'base.yaml'")
		log.Info("A variation based on openSUSE Tumbleweed is there too, named 'suse.yaml'")

		log.Info("\nNow let's build your base image, which will be used as a template for all your containers.")

		tmpdir, err := os.MkdirTemp(config.GetConfigPath(), "build")
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		baseDefinitionPath := filepath.Join(config.GetConfigPath(), "base.yaml")
		_, err = os.Stat(baseDefinitionPath)
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		defer os.RemoveAll(tmpdir)
		log.Info("* using temporary build directory" + tmpdir)
		command := exec.Command("sudo", "distrobuilder", "build-lxd", baseDefinitionPath, tmpdir)
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		log.Info("Starting build...")
		err = command.Run()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
		log.Info("Build complete")
		importCommand := exec.Command("lxc", "image", "import", filepath.Join(tmpdir, "lxd.tar.xz"), filepath.Join(tmpdir, "rootfs.squashfs"), "--alias", "dlxbase")
		importCommand.Stderr = os.Stderr
		importCommand.Stdout = os.Stdout
		log.Info("Starting import...")
		err = importCommand.Run()
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}

		log.Info("Import complete.")
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
