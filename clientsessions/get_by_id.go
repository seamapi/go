package clientsessions

import (
	context "context"

	seamapigo "github.com/seamapi/go"
)

func (c *Client) GetById(ctx context.Context, clientSessionId string) (*seamapigo.ClientSession, error) {
	return c.Get(
		ctx,
		&seamapigo.ClientSessionsGetRequest{
			ClientSessionId: seamapigo.String(clientSessionId),
		},
	)
}
