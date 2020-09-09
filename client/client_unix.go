<<<<<<< HEAD
// +build linux freebsd openbsd netbsd darwin dragonfly
=======
// +build linux freebsd openbsd netbsd darwin solaris illumos dragonfly
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375

package client // import "github.com/docker/docker/client"

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///var/run/docker.sock"

const defaultProto = "unix"
const defaultAddr = "/var/run/docker.sock"
