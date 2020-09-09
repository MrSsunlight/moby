package types // import "github.com/docker/docker/libcontainerd/types"

import (
	"time"

<<<<<<< HEAD
	statsV1 "github.com/containerd/cgroups/stats/v1"
=======
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// Summary is not used on linux
type Summary struct{}

// Stats holds metrics properties as returned by containerd
type Stats struct {
<<<<<<< HEAD
	Read    time.Time
	Metrics *statsV1.Metrics
=======
	Read time.Time
	// Metrics is expected to be either one of:
	// * github.com/containerd/cgroups/stats/v1.Metrics
	// * github.com/containerd/cgroups/stats/v2.Metrics
	Metrics interface{}
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}

// InterfaceToStats returns a stats object from the platform-specific interface.
func InterfaceToStats(read time.Time, v interface{}) *Stats {
	return &Stats{
<<<<<<< HEAD
		Metrics: v.(*statsV1.Metrics),
=======
		Metrics: v,
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
		Read:    read,
	}
}

// Resources defines updatable container resource values. TODO: it must match containerd upcoming API
type Resources specs.LinuxResources

// Checkpoints contains the details of a checkpoint
type Checkpoints struct{}
