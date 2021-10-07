package config

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func Create() error {
	//make config directory and file
	err := os.MkdirAll(filepath.Join(GetConfigPath()), 0755)
	if err != nil {
		return errors.Wrap(err, "failed to create config directory")
	}
	f, err := os.Create(filepath.Join(GetConfigPath(), "dlxweb.yaml"))
	if err != nil {
		return errors.Wrap(err, "failed to create config file")
	}
	defer f.Close()

	user, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "failed to get current user")
	}

	config := &Settings{
		SSHPrivateKey: "/home/" + user.Username + "/.ssh/id_rsa",
		BaseImage:     "dlxbase",
	}

	bb, err := yaml.Marshal(config)
	if err != nil {
		return errors.Wrap(err, "failed to marshal config")
	}
	_, err = f.Write(bb)
	if err != nil {
		return errors.Wrap(err, "failed to write config")
	}
	baseyaml, err := BaseTemplate()
	if err != nil {
		return errors.Wrap(err, "failed to load base image template")
	}
	err = ioutil.WriteFile(filepath.Join(GetConfigPath(), "base.yaml"), []byte(baseyaml), 0644)

	if err != nil {
		return errors.Wrap(err, "failed to write base image template")
	}
	return err
}

func Check() error {
	_, err := os.Stat(filepath.Join(GetConfigPath(), "dlxweb.yaml"))
	return err
}

func Get() (*Settings, error) {
	var cfg Settings
	err := Check()
	if err != nil {
		return &cfg, errors.Wrap(err, "config not found")
	}
	bb, err := ioutil.ReadFile(filepath.Join(GetConfigPath(), "dlxweb.yaml"))
	if err != nil {
		return &cfg, errors.Wrap(err, "failed to read config")
	}
	err = yaml.Unmarshal(bb, &cfg)
	return &cfg, errors.Wrap(err, "failed to unmarshal config")
}
