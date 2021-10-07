// main.go
package server

import (
	"log"
	"net/http"

	oserver "github.com/bketelsen/dlxweb/generated/server"
	"github.com/bketelsen/dlxweb/server/config"
	"github.com/bketelsen/dlxweb/state"
	"github.com/pacedotdev/oto/otohttp"
)

func Serve() {
	err := CheckDependencies()
	if err != nil {
		log.Fatal(err)
	}
	settings, err := config.Get()
	if err != nil {
		log.Println("Error loading config:", err)
		log.Println("Creating new config file")
		err = config.Create()
		if err != nil {
			log.Fatal(err)
		}
	}
	global := &state.Global{DlxConfig: settings}

	// create services
	instanceService := InstanceService{Global: global}
	imageService := ImageService{Global: global}
	profileService := ProfileService{Global: global}
	projectService := ProjectService{Global: global}

	// create the oto handler
	server := otohttp.NewServer()
	server.Basepath = "/oto/"

	// Register services
	oserver.RegisterInstanceService(server, instanceService)
	oserver.RegisterImageService(server, imageService)
	oserver.RegisterProfileService(server, profileService)
	oserver.RegisterProjectService(server, projectService)

	http.Handle(server.Basepath, server)
	/*s := &http.Server{
		TLSConfig: &tls.Config{
			GetCertificate: tailscale.GetCertificate,
		},
	}
	log.Printf("Running TLS server on :443 ...")
	log.Fatal(s.ListenAndServeTLS("", ""))
	*/
	http.ListenAndServe(":8080", http.DefaultServeMux)

}
