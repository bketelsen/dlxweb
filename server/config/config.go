package config

import (
	_ "embed"
	"os/user"
	"strings"
)

type Settings struct {
	SSHPrivateKey string `yaml:"ssh_private_key"`
	BaseImage     string `yaml:"baseimage"`
}

//go:embed base.yaml
var Base string

func BaseTemplate() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	Base = strings.Replace(Base, "dlxuser", user.Username, -1)
	return Base, nil
}
