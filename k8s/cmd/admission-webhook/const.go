package main

const (
	emptyBody            = "empty body"
	mutateMethod         = "/mutate"
	invalidContentType   = "invalid Content-Type=%v, expect \"application/json\""
	couldNotEncodeReview = "could not encode response: %v"
	couldNotWriteReview  = "could not write response: %v"
	deployment           = "Deployment"
	pod                  = "Pod"
	nsmAnnotationKey     = "ns.networkservicemesh.io"
	repoEnv              = "REPO"
	initContainerEnv     = "INITCONTAINER"
	tagEnv               = "TAG"
	repoDefault          = "networkservicemesh"
	initContainerDefault = "nsm-init"
	tagDefault           = "latest"
	initContainerName    = "nsm-init-container"
	pathInitContainers   = "/spec/initContainers"
	certFile             = "/etc/webhook/certs/cert.pem"
	keyFile              = "/etc/webhook/certs/key.pem"
	unsupportedKind      = "kind %v is not supported"
	deploymentSubPath    = "/spec/template"
	volumePath           = "/spec/volumes"
	containersPath       = "/spec/containers"
)