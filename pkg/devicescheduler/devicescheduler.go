package devicescheduler

import (
	"github.com/Microsoft/KubeDevice-API/pkg/types"
)

type PredicateFailureReason interface {
GetReason() string
	GetInfo() (types.ResourceName, int64, int64, int64)
}

// used by scheduler
type DeviceScheduler interface {
	// add node and resources
	AddNode(nodeName string, nodeInfo *types.NodeInfo)
	// remove node
	RemoveNode(nodeName string)
	// see if pod fits on node & return device score
	PodFitsDevice(nodeInfo *types.NodeInfo, podInfo *types.PodInfo, fillAllocateFrom bool) (bool, []PredicateFailureReason, float64)
	// allocate resources
	PodAllocate(nodeInfo *types.NodeInfo, podInfo *types.PodInfo) error
	// take resources from node
	TakePodResources(*types.NodeInfo, *types.PodInfo) error
	// return resources to node
	ReturnPodResources(*types.NodeInfo, *types.PodInfo) error
	// GetName returns the name of a device
	GetName() string
	// Tells whether group scheduler is being used?
	UsingGroupScheduler() bool
}
