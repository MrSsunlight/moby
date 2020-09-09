package oci

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/idtools"
	"github.com/docker/libnetwork/resolvconf"
	"github.com/docker/libnetwork/types"
	"github.com/moby/buildkit/util/flightcontrol"
	"github.com/pkg/errors"
)

var g flightcontrol.Group
var notFirstRun bool
var lastNotEmpty bool

<<<<<<< HEAD
=======
// overridden by tests
var resolvconfGet = resolvconf.Get

>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
type DNSConfig struct {
	Nameservers   []string
	Options       []string
	SearchDomains []string
}

func GetResolvConf(ctx context.Context, stateDir string, idmap *idtools.IdentityMapping, dns *DNSConfig) (string, error) {
	p := filepath.Join(stateDir, "resolv.conf")
	_, err := g.Do(ctx, stateDir, func(ctx context.Context) (interface{}, error) {
		generate := !notFirstRun
		notFirstRun = true

		if !generate {
			fi, err := os.Stat(p)
			if err != nil {
				if !errors.Is(err, os.ErrNotExist) {
					return "", err
				}
				generate = true
			}
			if !generate {
				fiMain, err := os.Stat(resolvconf.Path())
				if err != nil {
					if !errors.Is(err, os.ErrNotExist) {
						return nil, err
					}
					if lastNotEmpty {
						generate = true
						lastNotEmpty = false
					}
				} else {
					if fi.ModTime().Before(fiMain.ModTime()) {
						generate = true
					}
				}
			}
		}

		if !generate {
			return "", nil
		}

		var dt []byte
		f, err := resolvconfGet()
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return "", err
			}
		} else {
			dt = f.Content
		}

		if dns != nil {
			var (
				dnsNameservers   = resolvconf.GetNameservers(dt, types.IP)
				dnsSearchDomains = resolvconf.GetSearchDomains(dt)
				dnsOptions       = resolvconf.GetOptions(dt)
			)
			if len(dns.Nameservers) > 0 {
				dnsNameservers = dns.Nameservers
			}
			if len(dns.SearchDomains) > 0 {
				dnsSearchDomains = dns.SearchDomains
			}
			if len(dns.Options) > 0 {
				dnsOptions = dns.Options
			}

			f, err = resolvconf.Build(p+".tmp", dnsNameservers, dnsSearchDomains, dnsOptions)
			if err != nil {
				return "", err
			}
<<<<<<< HEAD
		} else {
			// Logic seems odd here: why are we filtering localhost IPs
			// only if neither of the DNS configs were specified?
			// Logic comes from https://github.com/docker/libnetwork/blob/164a77ee6d24fb2b1d61f8ad3403a51d8453899e/sandbox_dns_unix.go#L230-L269
			f, err = resolvconf.FilterResolvDNS(f.Content, true)
			if err != nil {
				return "", err
			}
=======
			dt = f.Content
		}

		f, err = resolvconf.FilterResolvDNS(dt, true)
		if err != nil {
			return "", err
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
		}

		tmpPath := p + ".tmp"
		if err := ioutil.WriteFile(tmpPath, f.Content, 0644); err != nil {
			return "", err
		}

		if idmap != nil {
			root := idmap.RootPair()
			if err := os.Chown(tmpPath, root.UID, root.GID); err != nil {
				return "", err
			}
		}

		if err := os.Rename(tmpPath, p); err != nil {
			return "", err
		}
		return "", nil
	})
	if err != nil {
		return "", err
	}
	return p, nil
}
