package device

import (
	"github.com/Microsoft/KubeDevice-API/pkg/types"
)

type Mount struct {
	ContainerPath string
	HostPath      string
	ReadOnly      bool
}

// Device is a device to use
type Device interface {
	// New creates the device and initializes it
	New() error
	// Start logically initializes the device
	Start() error
	// UpdateNodeInfo - updates a node info structure by writing capacity, allocatable, used, scorer
	UpdateNodeInfo(*types.NodeInfo) error
	// Allocate attempst to allocate the devices
	// Returns list of Mounts, and list of Devices to use
	// Returns an error on failure.
	Allocate(*types.PodInfo, *types.ContainerInfo) ([]Mount, []string, error)
	// GetName returns the name of a device
	GetName() string
}
