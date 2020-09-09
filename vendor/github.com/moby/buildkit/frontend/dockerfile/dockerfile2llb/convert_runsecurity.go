// +build dfrunsecurity

package dockerfile2llb

import (
	"github.com/pkg/errors"

<<<<<<< HEAD
=======
	"github.com/moby/buildkit/client/llb"
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
	"github.com/moby/buildkit/solver/pb"
)

<<<<<<< HEAD
func dispatchRunSecurity(d *dispatchState, c *instructions.RunCommand) error {
	security := instructions.GetSecurity(c)

	for _, sec := range security {
		switch sec {
		case instructions.SecurityInsecure:
			d.state = d.state.Security(pb.SecurityMode_INSECURE)
		case instructions.SecuritySandbox:
			d.state = d.state.Security(pb.SecurityMode_SANDBOX)
		default:
			return errors.Errorf("unsupported security mode %q", sec)
		}
	}

	return nil
=======
func dispatchRunSecurity(c *instructions.RunCommand) (llb.RunOption, error) {
	security := instructions.GetSecurity(c)

	switch security {
	case instructions.SecurityInsecure:
		return llb.Security(pb.SecurityMode_INSECURE), nil
	case instructions.SecuritySandbox:
		return llb.Security(pb.SecurityMode_SANDBOX), nil
	default:
		return nil, errors.Errorf("unsupported security mode %q", security)
	}
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}
