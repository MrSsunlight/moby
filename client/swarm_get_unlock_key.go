package client // import "github.com/docker/docker/client"

import (
	"context"
	"encoding/json"

	"github.com/docker/docker/api/types"
)

// SwarmGetUnlockKey retrieves the swarm's unlock key.
// Swarm可以在多个服务器或主机上创建容器集群服务；Compose 是一个在单个服务器或主机上创建多个容器的工具
func (cli *Client) SwarmGetUnlockKey(ctx context.Context) (types.SwarmUnlockKeyResponse, error) {
	serverResp, err := cli.get(ctx, "/swarm/unlockkey", nil, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return types.SwarmUnlockKeyResponse{}, err
	}

	var response types.SwarmUnlockKeyResponse
	err = json.NewDecoder(serverResp.body).Decode(&response)
	return response, err
}
