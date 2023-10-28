package client

import (
	context "context"

	seamapigo "github.com/seamapi/go"
)

func (c *Client) GetById(ctx context.Context, accessCodeId string) (*seamapigo.AccessCode, error) {
	return c.Get(
		ctx,
		&seamapigo.AccessCodesGetRequest{
			AccessCodeId: seamapigo.String(accessCodeId),
		},
	)
}
