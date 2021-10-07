package server

import (
	"context"

	oserver "github.com/bketelsen/dlxweb/generated/server"
	"github.com/bketelsen/dlxweb/state"
	"github.com/pkg/errors"
)

// ProjectService manages projects
type ProjectService struct {
	Global *state.Global
}

// List returns a list of profiles
func (i ProjectService) List(ctx context.Context, r oserver.ProjectListRequest) (*oserver.ProjectListResponse, error) {
	i.Global.PreRun()
	var err error
	conf := i.Global.Conf

	d, err := conf.GetInstanceServer(conf.DefaultRemote)
	if err != nil {
		return nil, errors.Wrap(err, "could not get instance server")
	}
	projects, err := d.GetProjects()
	if err != nil {
		return nil, errors.Wrap(err, "could not get profiles")
	}
	return &oserver.ProjectListResponse{
		Projects: projects,
	}, nil
}
