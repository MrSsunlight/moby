package auth

import (
	"context"

	"github.com/moby/buildkit/session"
<<<<<<< HEAD
	"github.com/pkg/errors"
=======
	"github.com/moby/buildkit/util/grpcerrors"
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
	"google.golang.org/grpc/codes"
)

func CredentialsFunc(sm *session.Manager, g session.Group) func(string) (string, string, error) {
	return func(host string) (string, string, error) {
		var user, secret string
		err := sm.Any(context.TODO(), g, func(ctx context.Context, _ string, c session.Caller) error {
			client := NewAuthClient(c.Conn())

			resp, err := client.Credentials(ctx, &CredentialsRequest{
				Host: host,
			})
			if err != nil {
				if grpcerrors.Code(err) == codes.Unimplemented {
					return nil
				}
				return err
			}
			user = resp.Username
			secret = resp.Secret
			return nil
		})
		if err != nil {
<<<<<<< HEAD
			if st, ok := status.FromError(errors.Cause(err)); ok && st.Code() == codes.Unimplemented {
				return "", "", nil
			}
			return "", "", errors.WithStack(err)
=======
			return "", "", err
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
		}
		return user, secret, nil
	}
}
