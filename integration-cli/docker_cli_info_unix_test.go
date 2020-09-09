// +build !windows

package main

import (
<<<<<<< HEAD
	"strings"
	"testing"

	"gotest.tools/assert"
)

func (s *DockerSuite) TestInfoSecurityOptions(c *testing.T) {
	testRequires(c, testEnv.IsLocalDaemon, seccompEnabled, Apparmor, DaemonIsLinux)

	out, _ := dockerCmd(c, "info")
	assert.Assert(c, strings.Contains(out, "Security Options:\n apparmor\n seccomp\n  Profile: default\n"))
=======
	"context"
	"testing"

	"github.com/docker/docker/client"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func (s *DockerSuite) TestInfoSecurityOptions(c *testing.T) {
	testRequires(c, testEnv.IsLocalDaemon, DaemonIsLinux)
	if !seccompEnabled() && !Apparmor() {
		c.Skip("test requires Seccomp and/or AppArmor")
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	assert.NilError(c, err)
	defer cli.Close()
	info, err := cli.Info(context.Background())
	assert.NilError(c, err)

	if Apparmor() {
		assert.Check(c, is.Contains(info.SecurityOptions, "name=apparmor"))
	}
	if seccompEnabled() {
		assert.Check(c, is.Contains(info.SecurityOptions, "name=seccomp,profile=default"))
	}
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}
