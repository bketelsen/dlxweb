package config

import (
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func init() {
	defaultProject := &Project{
		Name:               "default",
		LXDName:            "default",
		ContainerMountName: "projects",
		Profiles:           []string{"default"},
	}

	Projects["default"] = defaultProject

	servicesProject := &Project{
		Name:               "services",
		LXDName:            "services",
		ContainerMountName: "services",
		Profiles:           []string{"default"},
	}

	Projects["services"] = servicesProject
}

var Projects = make(map[string]*Project)

const configDirName = "dlxweb"

type Project struct {
	Name               string
	LXDName            string
	ContainerMountName string
	Profiles           []string
}

func GetProject(name string) *Project {
	if name == "" {
		name = "default"
	}
	if project, ok := Projects[name]; ok {
		return project
	}

	return nil
}

func (p *Project) MountPath() string {
	return filepath.Join(GetHomePath(), configDirName, p.Name)
}

func (p *Project) ContainerMountPath() string {
	return filepath.Join(GetHomePath(), p.ContainerMountName)
}

func (p *Project) CreateMountPath() error {
	if _, err := os.Stat(p.MountPath()); os.IsNotExist(err) {
		err := os.MkdirAll(p.MountPath(), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Project) InstanceMountPath(instanceName string) string {
	return filepath.Join(p.MountPath(), instanceName)
}

func (p *Project) CreateInstanceMountPath(instanceName string) error {
	if _, err := os.Stat(p.InstanceMountPath(instanceName)); os.IsNotExist(err) {
		err := os.MkdirAll(p.InstanceMountPath(instanceName), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetHomePath() string {
	// Find home directory form env
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("Get Home Dir: " + err.Error())
	}
	return home
}

func GetConfigPath() string {

	//Find the default config directory
	configPath := os.Getenv("XDG_CONFIG_HOME")
	if len(configPath) == 0 {
		configPath = filepath.Join(GetHomePath(), ".config")
	}
	//set the dlx config directory
	return filepath.Join(configPath, configDirName)
}
