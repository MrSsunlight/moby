package buildkit

import (
	"context"
	"errors"

	"github.com/docker/docker/daemon/config"
	"github.com/docker/docker/pkg/idtools"
	"github.com/docker/libnetwork"
	"github.com/moby/buildkit/cache"
	"github.com/moby/buildkit/executor"
	"github.com/moby/buildkit/executor/oci"
)

func newExecutor(_, _ string, _ libnetwork.NetworkController, _ *oci.DNSConfig, _ bool, _ *idtools.IdentityMapping) (executor.Executor, error) {
	return &winExecutor{}, nil
}

type winExecutor struct {
}

func (w *winExecutor) Run(ctx context.Context, id string, root cache.Mountable, mounts []executor.Mount, process executor.ProcessInfo, started chan<- struct{}) (err error) {
	return errors.New("buildkit executor not implemented for windows")
}

<<<<<<< HEAD
=======
func (w *winExecutor) Exec(ctx context.Context, id string, process executor.ProcessInfo) error {
	return errors.New("buildkit executor not implemented for windows")
}

>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
func getDNSConfig(config.DNSConfig) *oci.DNSConfig {
	return nil
}
