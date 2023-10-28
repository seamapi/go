package connectwebviews

import (
	context "context"

	seamapigo "github.com/seamapi/go"
)

func (c *Client) GetById(ctx context.Context, connectWebviewId string) (*seamapigo.ConnectWebview, error) {
	return c.Get(
		ctx,
		&seamapigo.ConnectWebviewsGetRequest{
			ConnectWebviewId: connectWebviewId,
		},
	)
}
