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

//go:embed ubuntu.yaml
var UbuntuBase string

//go:embed opensuse.yaml
var SuseBase string

func BaseTemplate() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	UbuntuBase = strings.Replace(UbuntuBase, "dlxuser", user.Username, -1)
	return UbuntuBase, nil
}
func SuseBaseTemplate() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	SuseBase = strings.Replace(SuseBase, "dlxuser", user.Username, -1)
	return SuseBase, nil
}
