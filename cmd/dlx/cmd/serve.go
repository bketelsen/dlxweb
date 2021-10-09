package cmd

import (
	"github.com/bketelsen/dlxweb/server"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// instanceCmd.PersistentFlags().String("foo", "", "A help for foo")
	serveCmd.PersistentFlags().String("port", "8080", "Listen port")
	serveCmd.PersistentFlags().String("bind", "", "Listen address (default all interfaces)")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//serveCmd.Flags().BoolP("init", "i", false, "Run the initialization wizard.")
}

func serve(cmd *cobra.Command, args []string) error {
	port, err := cmd.PersistentFlags().GetString("port")
	if err != nil {
		return errors.Wrap(err, "getting port flag")
	}
	bind, err := cmd.PersistentFlags().GetString("bind")
	if err != nil {
		return errors.Wrap(err, "getting bind flag")
	}
	server.Serve(port, bind)
	return nil
}
