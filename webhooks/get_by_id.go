package webhooks

import (
	context "context"

	seamapigo "github.com/seamapi/go"
)

func (c *Client) GetById(ctx context.Context, webhookId string) (*seamapigo.WebhooksGetResponse, error) {
	return c.Get(
		ctx,
		&seamapigo.WebhooksGetRequest{
			WebhookId: webhookId,
		},
	)
}
