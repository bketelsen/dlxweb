package cmd

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/bketelsen/dlxweb/server"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
)

//go:embed public
var public embed.FS

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

	serveCmd.Flags().Bool("local", false, "use local filesystem instead of embedded one")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//serveCmd.Flags().BoolP("init", "i", false, "Run the initialization wizard.")
}

func serve(cmd *cobra.Command, args []string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	oto, err := server.Register()
	if err != nil {
		log.Error(err.Error())
	}
	//workDir, _ := os.Getwd()
	//public := http.Dir(filepath.Join(workDir, "./", "frontend", "public"))
	local, err := cmd.Flags().GetBool("local")

	if err != nil {
		log.Error(err.Error())
	}

	staticHandler(r, "/dashboard/", public, local)

	http.Handle("/oto/", oto)
	http.Handle("/", r)
	fmt.Println("listening on http://localhost:8080")
	return http.ListenAndServe(":8080", nil)

}

func staticHandler(r chi.Router, path string, root embed.FS, local bool) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	path += "*"
	log.Info(path)

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		log.Info(pathPrefix)
		fs := http.StripPrefix(pathPrefix, http.FileServer(getFileSystem(local)))
		fs.ServeHTTP(w, r)
	})
}
func getFileSystem(useOS bool) http.FileSystem {
	if useOS {
		log.Info("using live mode")
		return http.FS(os.DirFS("frontend/public"))
	}

	log.Info("using embed mode")
	fsys, err := fs.Sub(public, "public")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
