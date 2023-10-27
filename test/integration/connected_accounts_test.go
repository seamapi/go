package integration

import (
	"context"
	"testing"

	seamgo "github.com/seamapi/go"
	seamclient "github.com/seamapi/go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const email = "user-4@example.com"

func TestConnectedAccounts(t *testing.T) {
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	ctx := context.Background()
	connectedAccount := getTestConnectedAccount(t, seam)
	connectedAccountID := *connectedAccount.ConnectedAccountId

	connectedAccount, err := seam.ConnectedAccounts.Get(
		ctx,
		seamgo.NewConnectedAccountsGetRequestFromConnectedAccountsGetRequestConnectedAccountId(
			&seamgo.ConnectedAccountsGetRequestConnectedAccountId{
				ConnectedAccountId: connectedAccountID,
			},
		),
	)
	require.NoError(t, err)
	assert.NotNil(t, connectedAccount.ConnectedAccountId)
	assert.Equal(t, connectedAccountID, *connectedAccount.ConnectedAccountId)

	// Fake Seam does not like email creation.
	// emailAccount, err := seam.ConnectedAccounts.Get(
	// 	ctx,
	// 	seamgo.NewConnectedAccountsGetRequestFromConnectedAccountsGetRequestEmail(
	// 		&seamgo.ConnectedAccountsGetRequestEmail{
	// 			Email: email,
	// 		},
	// 	),
	// )
	// require.NoError(t, err)
	// assert.Equal(t, connectedAccountID, emailAccount.ConnectedAccountId)

	connectedAccounts, err := seam.ConnectedAccounts.List(ctx)
	require.NoError(t, err)
	assert.Len(t, connectedAccounts, 2)

	_, err = seam.ConnectedAccounts.Delete(
		ctx,
		&seamgo.ConnectedAccountsDeleteRequest{
			ConnectedAccountId: connectedAccountID,
		},
	)
	require.NoError(t, err)

	connectedAccounts, err = seam.ConnectedAccounts.List(ctx)
	require.NoError(t, err)
	assert.Len(t, connectedAccounts, 1)
}

func getTestConnectedAccount(t *testing.T, seam *seamclient.Client) *seamgo.ConnectedAccount {
	connectedAccounts, err := seam.ConnectedAccounts.List(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(connectedAccounts), 1)
	return connectedAccounts[0]
}
