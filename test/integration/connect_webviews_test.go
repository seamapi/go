package integration

import (
	"context"
	"testing"

	seamgo "github.com/seamapi/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnectWebviews(t *testing.T) {
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	ctx := context.Background()
	createdWebview, err := seam.ConnectWebviews.Create(
		ctx,
		&seamgo.ConnectWebviewsCreateRequest{
			AcceptedProviders: []seamgo.AcceptedProvider{
				seamgo.AcceptedProviderIgloo,
			},
		},
	)
	require.NoError(t, err)
	assert.NotEmpty(t, createdWebview.ConnectWebviewId)

	gotWebview, err := seam.ConnectWebviews.Get(
		ctx,
		&seamgo.ConnectWebviewsGetRequest{
			ConnectWebviewId: createdWebview.ConnectWebviewId,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, createdWebview.ConnectWebviewId, gotWebview.ConnectWebviewId)

	newWebview, err := seam.ConnectWebviews.Create(
		ctx,
		&seamgo.ConnectWebviewsCreateRequest{
			ProviderCategory: seamgo.ProviderCategoryStable.Ptr(),
		},
	)
	require.NoError(t, err)
	assert.NotEmpty(t, newWebview.ConnectWebviewId)

	gotWebview, err = seam.ConnectWebviews.Get(
		ctx,
		&seamgo.ConnectWebviewsGetRequest{
			ConnectWebviewId: newWebview.ConnectWebviewId,
		},
	)
	require.NoError(t, err)
	assert.Len(t, gotWebview.AcceptedProviders, 1)
}
