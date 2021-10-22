package server

import (
	"context"
	"fmt"
	"log"

	oserver "github.com/bketelsen/dlxweb/generated/server"
	"github.com/bketelsen/dlxweb/server/config"
	"github.com/bketelsen/dlxweb/state"
	lxd "github.com/lxc/lxd/client"
	"github.com/pkg/errors"
)

// rofileService manages profiles
type ProfileService struct {
	Global *state.Global
}

// List returns a list of profiles
func (i ProfileService) List(ctx context.Context, r oserver.ProfileListRequest) (*oserver.ProfileListResponse, error) {
	project := config.GetProject(r.Project)
	if project == nil {
		return nil, fmt.Errorf("project %s not found", r.Project)
	}
	log.Println("project", project.Name)
	/*
		i.Global.FlagProject = config.GetProject(r.Project).Name
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
	profiles, err := d.GetProfiles()
	if err != nil {
		return nil, errors.Wrap(err, "could not get profiles")
	}
	return &oserver.ProfileListResponse{
		Profiles: profiles,
	}, nil
}
