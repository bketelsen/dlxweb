//go:build !linux
// +build !linux

package server

import (
	"errors"
	"github.com/go-chi/chi/v5"
)

func Register(r *chi.Mux) error {
	return errors.New("Serve only works on linux")
}
