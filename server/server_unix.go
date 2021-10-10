//go:build linux
// +build linux

package server

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	oserver "github.com/bketelsen/dlxweb/generated/server"
	"github.com/bketelsen/dlxweb/server/config"
	"github.com/bketelsen/dlxweb/state"
	"github.com/go-chi/chi/v5"
	"github.com/pacedotdev/oto/otohttp"
)

func Register(r *chi.Mux) error {
	err := CheckDependencies()
	if err != nil {
		log.Println(err)
		return err
	}
	settings, err := config.Get()
	if err != nil {
		log.Println("Error loading config:", err)
		log.Println("Creating new config file")
		err = config.Create()
		if err != nil {
			log.Println(err)
			return err
		}
	}
	global := &state.Global{DlxConfig: settings}
	err = ensureProfile(global)
	if err != nil {
		log.Println(err)
		return err
	}
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

	r.Handle(server.Basepath, server)
	fmt.Println(r.Routes())
	return nil
}

func ensureProfile(global *state.Global) error {
	project := config.GetProject("default")
	if project == nil {
		return fmt.Errorf("project %s not found", "default")
	}
	log.Println("project", project.Name)
	global.FlagProject = project.Name
	global.PreRun()
	var err error
	conf := global.Conf

	d, err := conf.GetInstanceServer(conf.DefaultRemote)
	if err != nil {
		return err
	}
	profile, etag, err := d.GetProfile("default")
	if err != nil {
		return err
	}
	var update bool
	_, ok := profile.Config["raw.idmap"]
	if !ok {
		update = true
	}

	_, ok = profile.Devices["keys"]
	if !ok {
		update = true
	}
	if update {
		profilePut := profile.Writable()
		profilePut.Config["raw.idmap"] = "both 1000 1000"
		keys := make(map[string]string)
		homedir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		sshdir := filepath.Join(homedir, ".ssh")
		keys["path"] = sshdir
		keys["source"] = sshdir
		keys["type"] = "disk"
		profilePut.Devices["keys"] = keys
		return d.UpdateProfile("default", profilePut, etag)
	} else {
		log.Println("Profile contains sshkeys and id map, not modifying")
	}
	return nil
}
