#!/bin/sh
# dockerd-rootless.sh executes dockerd in rootless mode.
#
<<<<<<< HEAD
# Usage: dockerd-rootless.sh --experimental [DOCKERD_OPTIONS]
# Currently, specifying --experimental is mandatory.
=======
# Usage: dockerd-rootless.sh [DOCKERD_OPTIONS]
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
#
# External dependencies:
# * newuidmap and newgidmap needs to be installed.
# * /etc/subuid and /etc/subgid needs to be configured for the current user.
# * Either one of slirp4netns (>= v0.4.0), VPNKit, lxc-user-nic needs to be installed.
<<<<<<< HEAD
#   slirp4netns is used by default if installed. Otherwise fallsback to VPNKit.
#   The default value can be overridden with $DOCKERD_ROOTLESS_ROOTLESSKIT_NET=(slirp4netns|vpnkit|lxc-user-nic)
=======
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
#
# Recognized environment variables:
# * DOCKERD_ROOTLESS_ROOTLESSKIT_NET=(slirp4netns|vpnkit|lxc-user-nic): the rootlesskit network driver. Defaults to "slirp4netns" if slirp4netns (>= v0.4.0) is installed. Otherwise defaults to "vpnkit".
# * DOCKERD_ROOTLESS_ROOTLESSKIT_MTU=NUM: the MTU value for the rootlesskit network driver. Defaults to 65520 for slirp4netns, 1500 for other drivers.
# * DOCKERD_ROOTLESS_ROOTLESSKIT_PORT_DRIVER=(builtin|slirp4netns): the rootlesskit port driver. Defaults to "builtin".
# * DOCKERD_ROOTLESS_ROOTLESSKIT_SLIRP4NETNS_SANDBOX=(auto|true|false): whether to protect slirp4netns with a dedicated mount namespace. Defaults to "auto".
# * DOCKERD_ROOTLESS_ROOTLESSKIT_SLIRP4NETNS_SECCOMP=(auto|true|false): whether to protect slirp4netns with seccomp. Defaults to "auto".
#
# See the documentation for the further information: https://docs.docker.com/engine/security/rootless/

set -e -x
if ! [ -w $XDG_RUNTIME_DIR ]; then
	echo "XDG_RUNTIME_DIR needs to be set and writable"
	exit 1
fi
if ! [ -w $HOME ]; then
	echo "HOME needs to be set and writable"
	exit 1
fi

rootlesskit=""
for f in docker-rootlesskit rootlesskit; do
	if which $f > /dev/null 2>&1; then
		rootlesskit=$f
		break
	fi
done
if [ -z $rootlesskit ]; then
	echo "rootlesskit needs to be installed"
	exit 1
fi

: "${DOCKERD_ROOTLESS_ROOTLESSKIT_NET:=}"
: "${DOCKERD_ROOTLESS_ROOTLESSKIT_MTU:=}"
<<<<<<< HEAD
=======
: "${DOCKERD_ROOTLESS_ROOTLESSKIT_PORT_DRIVER:=builtin}"
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
: "${DOCKERD_ROOTLESS_ROOTLESSKIT_SLIRP4NETNS_SANDBOX:=auto}"
: "${DOCKERD_ROOTLESS_ROOTLESSKIT_SLIRP4NETNS_SECCOMP:=auto}"
net=$DOCKERD_ROOTLESS_ROOTLESSKIT_NET
mtu=$DOCKERD_ROOTLESS_ROOTLESSKIT_MTU
if [ -z $net ]; then
<<<<<<< HEAD
	if which slirp4netns >/dev/null 2>&1; then
=======
	if which slirp4netns > /dev/null 2>&1; then
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
		# If --netns-type is present in --help, slirp4netns is >= v0.4.0.
		if slirp4netns --help | grep -qw -- --netns-type; then
			net=slirp4netns
			if [ -z $mtu ]; then
				mtu=65520
			fi
		else
			echo "slirp4netns found but seems older than v0.4.0. Falling back to VPNKit."
		fi
	fi
	if [ -z $net ]; then
<<<<<<< HEAD
		if which vpnkit >/dev/null 2>&1; then
=======
		if which vpnkit > /dev/null 2>&1; then
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
			net=vpnkit
		else
			echo "Either slirp4netns (>= v0.4.0) or vpnkit needs to be installed"
			exit 1
		fi
	fi
fi
if [ -z $mtu ]; then
	mtu=1500
fi

if [ -z $_DOCKERD_ROOTLESS_CHILD ]; then
	_DOCKERD_ROOTLESS_CHILD=1
	export _DOCKERD_ROOTLESS_CHILD
	# Re-exec the script via RootlessKit, so as to create unprivileged {user,mount,network} namespaces.
	#
	# --copy-up allows removing/creating files in the directories by creating tmpfs and symlinks
	# * /etc: copy-up is required so as to prevent `/etc/resolv.conf` in the
	#         namespace from being unexpectedly unmounted when `/etc/resolv.conf` is recreated on the host
	#         (by either systemd-networkd or NetworkManager)
	# * /run: copy-up is required so that we can create /run/docker (hardcoded for plugins) in our namespace
	exec $rootlesskit \
		--net=$net --mtu=$mtu \
		--slirp4netns-sandbox=$DOCKERD_ROOTLESS_ROOTLESSKIT_SLIRP4NETNS_SANDBOX \
		--slirp4netns-seccomp=$DOCKERD_ROOTLESS_ROOTLESSKIT_SLIRP4NETNS_SECCOMP \
<<<<<<< HEAD
		--disable-host-loopback --port-driver=builtin \
=======
		--disable-host-loopback --port-driver=$DOCKERD_ROOTLESS_ROOTLESSKIT_PORT_DRIVER \
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
		--copy-up=/etc --copy-up=/run \
		--propagation=rslave \
		$DOCKERD_ROOTLESS_ROOTLESSKIT_FLAGS \
		$0 $@
else
	[ $_DOCKERD_ROOTLESS_CHILD = 1 ]
	# remove the symlinks for the existing files in the parent namespace if any,
	# so that we can create our own files in our mount namespace.
	rm -f /run/docker /run/xtables.lock
	exec dockerd $@
fi
