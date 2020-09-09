// +build !dfrunsecurity

package dockerfile2llb

import (
<<<<<<< HEAD
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
)

func dispatchRunSecurity(d *dispatchState, c *instructions.RunCommand) error {
	return nil
=======
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/dockerfile/instructions"
)

func dispatchRunSecurity(c *instructions.RunCommand) (llb.RunOption, error) {
	return nil, nil
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}
