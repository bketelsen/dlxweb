package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/bketelsen/dlxweb/server"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/middleware"
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

	serveCmd.PersistentFlags().Bool("tailscale", false, "bind to tailscale IP on port 443 with TLS/Let's Encrypt")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//serveCmd.Flags().BoolP("init", "i", false, "Run the initialization wizard.")
}

func serve(cmd *cobra.Command, args []string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	err := server.Register(r)
	if err != nil {
		log.Error(err.Error())
	}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	workDir, _ := os.Getwd()
	public := http.Dir(filepath.Join(workDir, "./", "frontend", "public"))
	staticHandler(r, "/dashboard", public)
	fmt.Println("listening on http://localhost:3000")
	return http.ListenAndServe(":3000", r)

}

func staticHandler(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
