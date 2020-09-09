package idtools // import "github.com/docker/docker/pkg/idtools"

import (
	"testing"

<<<<<<< HEAD
	"gotest.tools/assert"
=======
	"gotest.tools/v3/assert"
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
)

func TestCreateIDMapOrder(t *testing.T) {
	subidRanges := ranges{
		{100000, 1000},
		{1000, 1},
	}

	idMap := createIDMap(subidRanges)
	assert.DeepEqual(t, idMap, []IDMap{
		{
			ContainerID: 0,
			HostID:      100000,
			Size:        1000,
		},
		{
			ContainerID: 1000,
			HostID:      1000,
			Size:        1,
		},
	})
}
