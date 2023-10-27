package integration

import (
	"context"
	"testing"

	seamgo "github.com/seamapi/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestActionAttempts(t *testing.T) {
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	ctx := context.Background()
	device := getTestDevice(t, seam)
	createdAccessCode, err := seam.AccessCodes.Create(
		ctx,
		&seamgo.AccessCodesCreateRequest{
			DeviceId: device.DeviceId,
			Name:     seamgo.String("Test code"),
			Code:     seamgo.String("4444"),
		},
	)
	require.NoError(t, err)

	actionAttempt, err := seam.AccessCodes.Delete(
		ctx,
		&seamgo.AccessCodesDeleteRequest{
			AccessCodeId: createdAccessCode.AccessCodeId,
		},
	)
	require.NoError(t, err)
	assert.NotNil(t, actionAttempt)
}
