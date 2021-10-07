package definitions

import "github.com/lxc/lxd/shared/api"

// ProjectService manages LXC projects
type ProjectService interface {
	// List returns a list of LXC projects
	List(ProjectListRequest) ProjectListResponse
}

type ProjectListRequest struct {
}

type ProjectListResponse struct {
	Projects []api.Project
}
