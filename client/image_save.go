package client // import "github.com/docker/docker/client"

import (
	"context"
	"io"
	"net/url"
)

// ImageSave retrieves one or more images from the docker host as an io.ReadCloser.
// It's up to the caller to store the images and close the stream.
// 从docker主机中检索一个或多个图像作为io.ReadCloser
func (cli *Client) ImageSave(ctx context.Context, imageIDs []string) (io.ReadCloser, error) {
	query := url.Values{
		"names": imageIDs,
	}

	resp, err := cli.get(ctx, "/images/get", query, nil)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}
