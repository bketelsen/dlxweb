package server

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	oserver "github.com/bketelsen/dlxweb/generated/server"
	"github.com/bketelsen/dlxweb/server/config"
	"github.com/bketelsen/dlxweb/state"
)

// InstanceService manages instances.
type ImageService struct {
	Global *state.Global
}

func (i ImageService) Build(ctx context.Context, r oserver.ImageBuildRequest) (*oserver.ImageBuildResponse, error) {

	project := config.GetProject(r.Project)
	if project == nil {
		return nil, fmt.Errorf("project %s not found", r.Project)
	}
	log.Println("project", project.Name)
	i.Global.FlagProject = config.GetProject(r.Project).Name
	i.Global.PreRun()

	baseDefinitionPath := filepath.Join(config.GetConfigPath(), "base.yaml")
	_, err := os.Stat(baseDefinitionPath)
	if err != nil {
		return nil, err
	}
	tmpdir, err := os.MkdirTemp(config.GetConfigPath(), "build")
	if err != nil {
		return nil, err
	}

	defer os.RemoveAll(tmpdir)
	log.Println("using temporary build directory", tmpdir)
	var b bytes.Buffer
	command := exec.Command("sudo", "distrobuilder", "build-lxd", baseDefinitionPath, tmpdir)
	command.Stderr = &b
	command.Stdout = &b
	log.Println("running build")
	err = command.Run()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("build complete")
	ioutil.WriteFile(config.GetConfigPath()+"/build.log", b.Bytes(), 0644)
	importCommand := exec.Command("lxc", "image", "import", filepath.Join(tmpdir, "lxd.tar.xz"), filepath.Join(tmpdir, "rootfs.squashfs"), "--alias", "dlxbase")
	var b2 bytes.Buffer
	importCommand.Stderr = &b2
	importCommand.Stdout = &b2
	log.Println("running import")
	err = importCommand.Run()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("import complete")
	ioutil.WriteFile(config.GetConfigPath()+"/import.log", b2.Bytes(), 0644)
	return &oserver.ImageBuildResponse{}, nil
}
func (i ImageService) Source(ctx context.Context, r oserver.ImageSourceRequest) (*oserver.ImageSourceResponse, error) {
	resp := &oserver.ImageSourceResponse{
		Source: config.Base,
	}
	return resp, nil
}
