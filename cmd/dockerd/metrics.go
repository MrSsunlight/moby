package main

import (
	"net"
	"net/http"
	"strings"

<<<<<<< HEAD
	"github.com/docker/go-metrics"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (cli *DaemonCli) startMetricsServer(addr string) error {
	if addr == "" {
		return nil
	}

	if !cli.d.HasExperimental() {
		return errors.New("metrics-addr is only supported when experimental is enabled")
	}

=======
	metrics "github.com/docker/go-metrics"
	"github.com/sirupsen/logrus"
)

func startMetricsServer(addr string) error {
	if addr == "" {
		return nil
	}
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
	if err := allocateDaemonPort(addr); err != nil {
		return err
	}
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/metrics", metrics.Handler())
	go func() {
		logrus.Infof("metrics API listening on %s", l.Addr())
		if err := http.Serve(l, mux); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			logrus.WithError(err).Error("error serving metrics API")
		}
	}()
	return nil
}
