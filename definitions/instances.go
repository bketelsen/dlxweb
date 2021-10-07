// definitions/definitons.go
package definitions

// InstanceService manages LXC Instances.
type InstanceService interface {
	// List returns a list of instances.
	List(InstanceListRequest) InstanceListResponse
	Create(InstanceCreateRequest) InstanceCreateResponse
	Stop(InstanceStopRequest) InstanceStopResponse
	Start(InstanceStartRequest) InstanceStartResponse
	Delete(InstanceDeleteRequest) InstanceDeleteResponse
}

type InstanceDeleteRequest struct {
	Name    string
	Project string
}

type InstanceDeleteResponse struct {
}

type InstanceStartRequest struct {
	Name    string
	Project string
}

type InstanceStartResponse struct {
	Instance InstanceDetails
}

type InstanceStopRequest struct {
	Name    string
	Project string
}

type InstanceStopResponse struct {
	Instance InstanceDetails
}

// CreateRequest is the request to create an instance.
type InstanceCreateRequest struct {
	Name    string
	Project string
}

// CreateResponse is the response from CreateInstance.
type InstanceCreateResponse struct {
	Instance InstanceDetails
}

// ListRequest is the request object for ListService.List.
type InstanceListRequest struct {
	Project string `json:"project"`
}

// ListResponse is the response object containing a
// list of instances.
type InstanceListResponse struct {

	// Instances is a list of LXC instances.
	Instances []InstanceDetails
}

// InstanceDetails is the details of an LXC instance.
type InstanceDetails struct {
	// Name is the name of the LXC instance.
	Name string
	// Status is the status of the Instance
	Status string
	// IPV4 is the IP address of the Instance
	IPV4 string
}
