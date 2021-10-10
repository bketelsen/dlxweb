//go:build !linux
// +build !linux

package server

import (
	"errors"
	"net/http"
)

func Register() (http.Handler, error) {
	return nil, errors.New("Serve only works on linux")
}
