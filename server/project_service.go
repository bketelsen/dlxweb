package server

import (
	"context"
	"log"

	oserver "github.com/bketelsen/dlxweb/generated/server"
	"github.com/bketelsen/dlxweb/state"
	lxd "github.com/lxc/lxd/client"
	"github.com/pkg/errors"
)

// ProjectService manages projects
type ProjectService struct {
	Global *state.Global
}

// List returns a list of profiles
func (i ProjectService) List(ctx context.Context, r oserver.ProjectListRequest) (*oserver.ProjectListResponse, error) {
	/*
		i.Global.PreRun()
		var err error
		conf := i.Global.Conf

		d, err := conf.GetInstanceServer(conf.DefaultRemote)
		if err != nil {
			return nil, errors.Wrap(err, "could not get instance server")
		}
	*/
	d, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	projects, err := d.GetProjects()
	if err != nil {
		return nil, errors.Wrap(err, "could not get projects")
	}
	return &oserver.ProjectListResponse{
		Projects: projects,
	}, nil
}
