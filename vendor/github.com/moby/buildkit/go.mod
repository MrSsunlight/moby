module github.com/moby/buildkit

<<<<<<< HEAD
go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/Microsoft/go-winio v0.4.13-0.20190408173621-84b4ab48a507
	github.com/Microsoft/hcsshim v0.8.5 // indirect
	github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7 // indirect
	github.com/codahale/hdrhistogram v0.0.0-20160425231609-f8ad88b59a58 // indirect
	github.com/containerd/cgroups v0.0.0-20190226200435-dbea6f2bd416 // indirect
	github.com/containerd/console v0.0.0-20181022165439-0650fd9eeb50
	github.com/containerd/containerd v1.3.0-0.20190507210959-7c1e88399ec0
	github.com/containerd/continuity v0.0.0-20190827140505-75bee3e2ccb6
	github.com/containerd/fifo v0.0.0-20180307165137-3d5202aec260 // indirect
	github.com/containerd/go-cni v0.0.0-20190610170741-5a4663dad645
	github.com/containerd/go-runc v0.0.0-20190911050354-e029b79d8cda
	github.com/containerd/ttrpc v0.0.0-20190411181408-699c4e40d1e7 // indirect
	github.com/containerd/typeurl v0.0.0-20180627222232-a93fcdb778cd // indirect
	github.com/containernetworking/cni v0.6.1-0.20180218032124-142cde0c766c // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
	github.com/docker/cli v0.0.0-20190321234815-f40f9c240ab0
	github.com/docker/distribution v2.7.1-0.20190205005809-0d3efadf0154+incompatible
	github.com/docker/docker v1.14.0-0.20190319215453-e7b5f7dbe98c
	github.com/docker/docker-credential-helpers v0.6.0 // indirect
	github.com/docker/go-connections v0.3.0
	github.com/docker/go-events v0.0.0-20170721190031-9461782956ad // indirect
	github.com/docker/libnetwork v0.8.0-dev.2.0.20190604151032-3c26b4e7495e
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/gofrs/flock v0.7.0
	github.com/gogo/googleapis v1.1.0
	github.com/gogo/protobuf v1.2.0
	github.com/golang/protobuf v1.2.0
	github.com/google/go-cmp v0.2.0
	github.com/google/shlex v0.0.0-20150127133951-6f45313302b9
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/go-immutable-radix v1.0.0
	github.com/hashicorp/golang-lru v0.0.0-20160207214719-a0d98a5f2880
	github.com/hashicorp/uuid v0.0.0-20160311170451-ebb0a03e909c // indirect
	github.com/ishidawataru/sctp v0.0.0-20180213033435-07191f837fed // indirect
	github.com/jaguilar/vt100 v0.0.0-20150826170717-2703a27b14ea
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mitchellh/hashstructure v0.0.0-20170609045927-2bca23e0e452
	github.com/morikuni/aec v0.0.0-20170113033406-39771216ff4c
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/opencontainers/image-spec v1.0.1
	github.com/opencontainers/runc v1.0.0-rc8
	github.com/opencontainers/runtime-spec v0.0.0-20180909173843-eba862dc2470
	github.com/opentracing-contrib/go-stdlib v0.0.0-20171029140428-b1a47cfbdd75
	github.com/opentracing/opentracing-go v0.0.0-20171003133519-1361b9cd60be
	github.com/pkg/errors v0.8.1
	github.com/pkg/profile v1.2.1
	github.com/serialx/hashring v0.0.0-20190422032157-8b2912629002
	github.com/sirupsen/logrus v1.3.0
	github.com/stretchr/testify v1.3.0
	github.com/syndtr/gocapability v0.0.0-20180916011248-d98352740cb2 // indirect
	github.com/tonistiigi/fsutil v0.0.0-20200128191323-6c909ab392c1
	github.com/tonistiigi/units v0.0.0-20180711220420-6950e57a87ea
	github.com/uber/jaeger-client-go v0.0.0-20180103221425-e02c85f9069e
	github.com/uber/jaeger-lib v1.2.1 // indirect
	github.com/urfave/cli v0.0.0-20171014202726-7bc6a0acffa5
	github.com/vishvananda/netlink v1.0.0 // indirect
	github.com/vishvananda/netns v0.0.0-20180720170159-13995c7128cc // indirect
	go.etcd.io/bbolt v1.3.2
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys v0.0.0-20190303122642-d455e41777fc
	golang.org/x/time v0.0.0-20161028155119-f51c12702a4d
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc v1.20.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gotest.tools v2.2.0+incompatible
)

replace github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe

replace github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
=======
go 1.13

require (
	github.com/AkihiroSuda/containerd-fuse-overlayfs v0.0.0-20200512015515-32086ef23a5a
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/Microsoft/go-winio v0.4.15-0.20190919025122-fc70bd9a86b5
	github.com/apache/thrift v0.0.0-20161221203622-b2a4d4ae21c7 // indirect
	github.com/codahale/hdrhistogram v0.0.0-20160425231609-f8ad88b59a58 // indirect
	github.com/containerd/cgroups v0.0.0-20200327175542-b44481373989 // indirect
	github.com/containerd/console v1.0.0
	github.com/containerd/containerd v1.4.0-0
	github.com/containerd/continuity v0.0.0-20200413184840-d3ef23f19fbb
	github.com/containerd/fifo v0.0.0-20200410184934-f15a3290365b // indirect
	github.com/containerd/go-cni v0.0.0-20200107172653-c154a49e2c75
	github.com/containerd/go-runc v0.0.0-20200220073739-7016d3ce2328
	github.com/coreos/go-systemd/v22 v22.0.0
	github.com/docker/cli v0.0.0-20200227165822-2298e6a3fe24
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v0.0.0
	github.com/docker/docker-credential-helpers v0.6.0 // indirect
	github.com/docker/go-connections v0.3.0
	github.com/docker/libnetwork v0.8.0-dev.2.0.20200226230617-d8334ccdb9be
	github.com/gofrs/flock v0.7.0
	github.com/gogo/googleapis v1.3.2
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.3.3
	github.com/google/go-cmp v0.4.0
	github.com/google/shlex v0.0.0-20150127133951-6f45313302b9
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/go-immutable-radix v1.0.0
	github.com/hashicorp/golang-lru v0.5.1
	github.com/hashicorp/uuid v0.0.0-20160311170451-ebb0a03e909c // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/ishidawataru/sctp v0.0.0-20191218070446-00ab2ac2db07 // indirect
	github.com/jaguilar/vt100 v0.0.0-20150826170717-2703a27b14ea
	github.com/mitchellh/hashstructure v1.0.0
	github.com/morikuni/aec v0.0.0-20170113033406-39771216ff4c
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.1
	github.com/opencontainers/runc v1.0.0-rc10
	github.com/opencontainers/runtime-spec v1.0.2
	github.com/opencontainers/selinux v1.5.1 // indirect
	github.com/opentracing-contrib/go-stdlib v0.0.0-20171029140428-b1a47cfbdd75
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/pkg/profile v1.2.1
	github.com/serialx/hashring v0.0.0-20190422032157-8b2912629002
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.5.1
	github.com/syndtr/gocapability v0.0.0-20180916011248-d98352740cb2 // indirect
	github.com/tonistiigi/fsutil v0.0.0-20200512175118-ae3a8d753069
	github.com/tonistiigi/units v0.0.0-20180711220420-6950e57a87ea
	github.com/uber/jaeger-client-go v2.11.2+incompatible
	github.com/uber/jaeger-lib v1.2.1 // indirect
	github.com/urfave/cli v1.22.2
	github.com/vishvananda/netlink v1.1.0 // indirect
	go.etcd.io/bbolt v1.3.3
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20200302150141-5c8b2ff67527
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0
	google.golang.org/genproto v0.0.0-20200227132054-3f1135a288c9
	google.golang.org/grpc v1.27.1
)

replace (
	github.com/containerd/containerd => github.com/containerd/containerd v1.3.1-0.20200512144102-f13ba8f2f2fd
	github.com/docker/docker => github.com/docker/docker v17.12.0-ce-rc1.0.20200310163718-4634ce647cf2+incompatible
	github.com/hashicorp/go-immutable-radix => github.com/tonistiigi/go-immutable-radix v0.0.0-20170803185627-826af9ccf0fe
	github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
)
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
