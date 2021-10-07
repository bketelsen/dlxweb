package definitions

import "github.com/lxc/lxd/shared/api"

// ProfileService manages LXC profiles
type ProfileService interface {
	// List returns a list of LXC profiles
	List(ProfileListRequest) ProfileListResponse
}

type ProfileListRequest struct {
	Project string
}

type ProfileListResponse struct {
	Profiles []api.Profile
}
