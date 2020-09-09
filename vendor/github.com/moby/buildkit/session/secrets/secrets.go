package secrets

import (
	"context"

	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/util/grpcerrors"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

type SecretStore interface {
	GetSecret(context.Context, string) ([]byte, error)
}

var ErrNotFound = errors.Errorf("not found")

func GetSecret(ctx context.Context, c session.Caller, id string) ([]byte, error) {
	client := NewSecretsClient(c.Conn())
	resp, err := client.GetSecret(ctx, &GetSecretRequest{
		ID: id,
	})
	if err != nil {
<<<<<<< HEAD
		if st, ok := status.FromError(errors.Cause(err)); ok && (st.Code() == codes.Unimplemented || st.Code() == codes.NotFound) {
=======
		if code := grpcerrors.Code(err); code == codes.Unimplemented || code == codes.NotFound {
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
			return nil, errors.Wrapf(ErrNotFound, "secret %s not found", id)
		}
		return nil, errors.WithStack(err)
	}
	return resp.Data, nil
}
