//go:build !linux
// +build !linux

package server

import (
	"fmt"
	"os"
)

func Serve(port, bind string, useTailscale bool) {
	fmt.Println("Serve only works on linux")
	os.Exit(-1)
}
