package leaseutil

import (
	"context"
	"time"

	"github.com/containerd/containerd/leases"
<<<<<<< HEAD
	"github.com/containerd/containerd/metadata"
	"github.com/containerd/containerd/namespaces"
	bolt "go.etcd.io/bbolt"
=======
	"github.com/containerd/containerd/namespaces"
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
)

func WithLease(ctx context.Context, ls leases.Manager, opts ...leases.Opt) (context.Context, func(context.Context) error, error) {
	_, ok := leases.FromContext(ctx)
	if ok {
		return ctx, func(context.Context) error {
			return nil
		}, nil
	}

	l, err := ls.Create(ctx, append([]leases.Opt{leases.WithRandomID(), leases.WithExpiration(time.Hour)}, opts...)...)
	if err != nil {
		return nil, nil, err
	}

	ctx = leases.WithLease(ctx, l.ID)
	return ctx, func(ctx context.Context) error {
		return ls.Delete(ctx, l)
	}, nil
}

<<<<<<< HEAD
func NewManager(mdb *metadata.DB) leases.Manager {
	return &local{db: mdb}
}

type local struct {
	db *metadata.DB
}

func (l *local) Create(ctx context.Context, opts ...leases.Opt) (leases.Lease, error) {
	var lease leases.Lease
	if err := l.db.Update(func(tx *bolt.Tx) error {
		var err error
		lease, err = metadata.NewLeaseManager(tx).Create(ctx, opts...)
		return err
	}); err != nil {
		return leases.Lease{}, err
	}
	return lease, nil
}

func (l *local) Delete(ctx context.Context, lease leases.Lease, opts ...leases.DeleteOpt) error {
	var do leases.DeleteOptions
	for _, opt := range opts {
		if err := opt(ctx, &do); err != nil {
			return err
		}
	}

	if err := l.db.Update(func(tx *bolt.Tx) error {
		return metadata.NewLeaseManager(tx).Delete(ctx, lease)
	}); err != nil {
		return err
	}

	return nil

}

func (l *local) List(ctx context.Context, filters ...string) ([]leases.Lease, error) {
	var ll []leases.Lease
	if err := l.db.View(func(tx *bolt.Tx) error {
		var err error
		ll, err = metadata.NewLeaseManager(tx).List(ctx, filters...)
		return err
	}); err != nil {
		return nil, err
	}
	return ll, nil
}

func WithNamespace(lm leases.Manager, ns string) leases.Manager {
	return &nsLM{Manager: lm, ns: ns}
}

type nsLM struct {
	leases.Manager
	ns string
=======
func MakeTemporary(l *leases.Lease) error {
	if l.Labels == nil {
		l.Labels = map[string]string{}
	}
	l.Labels["buildkit/lease.temporary"] = time.Now().UTC().Format(time.RFC3339Nano)
	return nil
}

func WithNamespace(lm leases.Manager, ns string) leases.Manager {
	return &nsLM{manager: lm, ns: ns}
}

type nsLM struct {
	manager leases.Manager
	ns      string
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}

func (l *nsLM) Create(ctx context.Context, opts ...leases.Opt) (leases.Lease, error) {
	ctx = namespaces.WithNamespace(ctx, l.ns)
<<<<<<< HEAD
	return l.Manager.Create(ctx, opts...)
=======
	return l.manager.Create(ctx, opts...)
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}

func (l *nsLM) Delete(ctx context.Context, lease leases.Lease, opts ...leases.DeleteOpt) error {
	ctx = namespaces.WithNamespace(ctx, l.ns)
<<<<<<< HEAD
	return l.Manager.Delete(ctx, lease, opts...)
=======
	return l.manager.Delete(ctx, lease, opts...)
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}

func (l *nsLM) List(ctx context.Context, filters ...string) ([]leases.Lease, error) {
	ctx = namespaces.WithNamespace(ctx, l.ns)
<<<<<<< HEAD
	return l.Manager.List(ctx, filters...)
=======
	return l.manager.List(ctx, filters...)
}

func (l *nsLM) AddResource(ctx context.Context, lease leases.Lease, resource leases.Resource) error {
	ctx = namespaces.WithNamespace(ctx, l.ns)
	return l.manager.AddResource(ctx, lease, resource)
}

func (l *nsLM) DeleteResource(ctx context.Context, lease leases.Lease, resource leases.Resource) error {
	ctx = namespaces.WithNamespace(ctx, l.ns)
	return l.manager.DeleteResource(ctx, lease, resource)
}

func (l *nsLM) ListResources(ctx context.Context, lease leases.Lease) ([]leases.Resource, error) {
	ctx = namespaces.WithNamespace(ctx, l.ns)
	return l.manager.ListResources(ctx, lease)
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}
