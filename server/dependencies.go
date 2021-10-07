package server

import (
	"fmt"
	"log"
	"os/exec"
)

func CheckDependencies() error {
	if err := checkBinary("lxd"); err != nil {
		log.Println("lxd was not found")
		log.Println("Install lxd with the following command:")
		log.Println("	sudo snap install lxd --classic")
		return err
	}
	if err := checkBinary("debootstrap"); err != nil {
		log.Println("debootstrap was not found")
		log.Println("Install debootstrap with the following command:")
		log.Println("	sudo apt-get install -y debootstrap")
		return err
	}

	if err := checkBinary("distrobuilder"); err != nil {
		log.Println("distrobuilder was not found")
		log.Println("Install distrobuilder with the following command:")
		log.Println("	sudo snap install distrobuilder --classic")
		return err
	}
	return nil
}

func checkBinary(binary string) error {
	path, err := exec.LookPath(binary)
	if err != nil {
		return fmt.Errorf("%s not found in PATH", binary)
	}
	log.Printf("found %s: %s\n", binary, path)
	return nil
}
