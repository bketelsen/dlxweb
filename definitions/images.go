package definitions

// ImageService manages LXC images
type ImageService interface {
	// Build builds and imports the base image
	Build(ImageBuildRequest) ImageBuildResponse
	Source(ImageSourceRequest) ImageSourceResponse
}

type ImageBuildRequest struct {
	Project string
	Source  string
}

type ImageBuildResponse struct {
}
type ImageSourceRequest struct {
	Project string
}

type ImageSourceResponse struct {
	Source string
}
