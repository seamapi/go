package actionattempts

import (
	context "context"

	seamapigo "github.com/seamapi/go"
)

func (c *Client) GetById(ctx context.Context, actionAttemptId string) (*seamapigo.ActionAttempt, error) {
	return c.Get(
		ctx,
		&seamapigo.ActionAttemptsGetRequest{
			ActionAttemptId: actionAttemptId,
		},
	)
}
