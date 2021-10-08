package server

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	oserver "github.com/bketelsen/dlxweb/generated/server"
	"github.com/bketelsen/dlxweb/server/config"
	"github.com/bketelsen/dlxweb/state"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/i18n"
	"github.com/pkg/errors"
)

// InstanceService manages instances.
type InstanceService struct {
	Global *state.Global
}

func (i InstanceService) Delete(ctx context.Context, r oserver.InstanceDeleteRequest) (*oserver.InstanceDeleteResponse, error) {

	project := config.GetProject(r.Project)
	if project == nil {
		return nil, fmt.Errorf("project %s not found", r.Project)
	}
	log.Println("project", project.Name)
	i.Global.FlagProject = config.GetProject(r.Project).Name
	i.Global.PreRun()
	var err error
	conf := i.Global.Conf

	d, err := conf.GetInstanceServer(conf.DefaultRemote)
	if err != nil {
		return nil, err
	}
	name := r.Name
	op, err := d.DeleteInstance(name)

	if err != nil {
		return nil, errors.Wrap(err, "deleting container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "wait for delete container")
	}

	return &oserver.InstanceDeleteResponse{}, nil
}
func (i InstanceService) Create(ctx context.Context, r oserver.InstanceCreateRequest) (*oserver.InstanceCreateResponse, error) {
	project := config.GetProject(r.Project)
	if project == nil {
		return nil, fmt.Errorf("project %s not found", r.Project)
	}
	log.Println("project", project.Name)
	i.Global.FlagProject = config.GetProject(r.Project).Name
	i.Global.PreRun()
	var err error
	conf := i.Global.Conf

	d, err := conf.GetInstanceServer(conf.DefaultRemote)
	if err != nil {
		return nil, err
	}
	name := r.Name

	bi := i.Global.DlxConfig.BaseImage

	source := api.ContainerSource{
		Type: "image",
		//Server:   "https://cloud-images.ubuntu.com/daily",
		//Alias:    getImage(),
		Alias: bi,
		//Protocol: "simplestreams",
	}

	req := api.ContainersPost{
		Name: name,
		ContainerPut: api.ContainerPut{
			Profiles: []string{"default"}, // TODO: ? support command line adding profiles
		},
		Source: source,
	}

	// Get LXD to create the container (background operation)
	op, err := d.CreateContainer(req)
	if err != nil {
		return nil, errors.Wrap(err, "creating container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "wait for create container")
	}

	// Get LXD to start the container (background operation)
	reqState := api.ContainerStatePut{
		Action:  "start",
		Timeout: -1,
	}

	op, err = d.UpdateContainerState(name, reqState, "")
	if err != nil {
		return nil, errors.Wrap(err, "starting container")
	}

	// Wait for the operation to complete
	err = op.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "waiting for container start")
	}

	// Mount the project directory into container FS

	devname := "persist"
	devSource := "source=" + project.InstanceMountPath(name)
	devPath := "path=" + project.ContainerMountPath()
	log.Println("Source: ", devSource)
	log.Println("Path: ", devPath)

	err = project.CreateInstanceMountPath(name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create host mount path")
	}
	err = addDevice(d, name, []string{devname, "disk", devSource, devPath})
	if err != nil {
		return nil, errors.Wrap(err, "failed to mount project directory")
	}

	inst, err := i.getInstanceFull(d, name)
	if err != nil {
		return nil, errors.Wrap(err, "get instance details")
	}
	resp := &oserver.InstanceCreateResponse{
		Instance: oserver.InstanceDetails{
			Name:   inst.Name,
			IPV4:   i.IP4ColumnData(*inst),
			Status: inst.Status,
		},
	}
	return resp, nil

}
func addDevice(d lxd.InstanceServer, name string, args []string) error {

	// Add the device
	devname := args[0]
	device := map[string]string{}
	device["type"] = args[1]
	if len(args) > 2 {
		for _, prop := range args[2:] {
			results := strings.SplitN(prop, "=", 2)
			if len(results) != 2 {
				return fmt.Errorf(i18n.G("No value found in %q"), prop)
			}
			k := results[0]
			v := results[1]
			device[k] = v
		}
	}

	inst, etag, err := d.GetInstance(name)
	if err != nil {
		return err
	}

	_, ok := inst.Devices[devname]
	if ok {
		return fmt.Errorf(i18n.G("The device already exists"))
	}

	inst.Devices[devname] = device

	op, err := d.UpdateInstance(name, inst.Writable(), etag)
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return err
	}

	return nil
}
func (i InstanceService) Start(ctx context.Context, r oserver.InstanceStartRequest) (*oserver.InstanceStartResponse, error) {
	project := config.GetProject(r.Project)
	if project == nil {
		return nil, fmt.Errorf("project %s not found", r.Project)
	}
	log.Println("project", project.Name)
	i.Global.FlagProject = config.GetProject(r.Project).Name
	i.Global.PreRun()
	var err error
	conf := i.Global.Conf

	d, err := conf.GetInstanceServer(conf.DefaultRemote)
	if err != nil {
		return nil, err
	}
	err = i.putState(d, r.Name, "start")
	if err != nil {
		return nil, err
	}

	inst, err := i.getInstanceFull(d, r.Name)
	if err != nil {
		return nil, errors.Wrap(err, "get instance details")
	}
	resp := &oserver.InstanceStartResponse{
		Instance: oserver.InstanceDetails{
			Name:   inst.Name,
			IPV4:   i.IP4ColumnData(*inst),
			Status: inst.Status,
		},
	}
	return resp, nil
}
func (i InstanceService) Stop(ctx context.Context, r oserver.InstanceStopRequest) (*oserver.InstanceStopResponse, error) {
	project := config.GetProject(r.Project)
	if project == nil {
		return nil, fmt.Errorf("project %s not found", r.Project)
	}
	log.Println("project", project.Name)
	i.Global.FlagProject = config.GetProject(r.Project).Name
	i.Global.PreRun()
	var err error
	conf := i.Global.Conf

	d, err := conf.GetInstanceServer(conf.DefaultRemote)
	if err != nil {
		return nil, err
	}
	err = i.putState(d, r.Name, "stop")
	if err != nil {
		return nil, err
	}
	inst, err := i.getInstanceFull(d, r.Name)
	if err != nil {
		return nil, errors.Wrap(err, "get instance details")
	}
	resp := &oserver.InstanceStopResponse{
		Instance: oserver.InstanceDetails{
			Name:   inst.Name,
			IPV4:   i.IP4ColumnData(*inst),
			Status: inst.Status,
		},
	}
	return resp, nil
}

func (i InstanceService) putState(d lxd.InstanceServer, name string, action string) error {

	req := api.InstanceStatePut{
		Action:   action,
		Timeout:  10,
		Force:    false,
		Stateful: false,
	}

	op, err := d.UpdateInstanceState(name, req, "")
	if err != nil {
		return err
	}

	err = op.Wait()
	if err != nil {
		return errors.Wrap(err, "waiting for container action")
	}
	return nil
}

func (i InstanceService) getInstanceFull(d lxd.InstanceServer, name string) (*api.InstanceFull, error) {

	instances, err := d.GetInstancesFull(api.InstanceTypeAny)
	if err != nil {
		errors.Wrap(err, "get container names")
		return nil, err
	}
	for _, inst := range instances {
		if inst.Name == name {

			return &inst, nil
		}
	}
	return nil, fmt.Errorf("instance %s not found", name)
}

// List returns a list of instances.
func (i InstanceService) List(ctx context.Context, r oserver.InstanceListRequest) (*oserver.InstanceListResponse, error) {
	log.Println("request", r.Project)
	project := config.GetProject(r.Project)
	if project == nil {
		return nil, fmt.Errorf("project %s not found", r.Project)
	}
	log.Println("project", project.Name)
	i.Global.FlagProject = config.GetProject(r.Project).Name

	i.Global.PreRun()
	var err error
	conf := i.Global.Conf

	d, err := conf.GetInstanceServer(conf.DefaultRemote)
	if err != nil {
		return nil, err
	}

	instances, err := d.GetInstancesFull(api.InstanceTypeAny)
	if err != nil {
		errors.Wrap(err, "get container names")
		return nil, err
	}

	var details []oserver.InstanceDetails
	for _, instance := range instances {

		details = append(details, oserver.InstanceDetails{
			Name:   instance.Name,
			Status: instance.Status,
			IPV4:   i.IP4ColumnData(instance),
		})

	}
	resp := &oserver.InstanceListResponse{
		Instances: details,
	}
	return resp, nil
}

// IP4ColumnData returns the IPV4 column data for an instance.
//
// Copied from github.com/lxc/lxd
func (c InstanceService) IP4ColumnData(cInfo api.InstanceFull) string {
	if cInfo.IsActive() && cInfo.State != nil && cInfo.State.Network != nil {
		ipv4s := []string{}
		for netName, net := range cInfo.State.Network {
			if net.Type == "loopback" {
				continue
			}

			for _, addr := range net.Addresses {
				if shared.StringInSlice(addr.Scope, []string{"link", "local"}) {
					continue
				}

				if addr.Family == "inet" {
					ipv4s = append(ipv4s, fmt.Sprintf("%s (%s)", addr.Address, netName))
				}
			}
		}
		sort.Sort(sort.Reverse(sort.StringSlice(ipv4s)))
		return strings.Join(ipv4s, "\n")
	}

	return ""
}
