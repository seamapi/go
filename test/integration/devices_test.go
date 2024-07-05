package integration

import (
	"context"
	"testing"

	seamgo "github.com/seamapi/go"
	seamclient "github.com/seamapi/go/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDevices(t *testing.T) {
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	ctx := context.Background()
	device := getTestDevice(t, seam)
	assert.NotNil(t, device.Properties)

	devices, err := seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			ConnectedAccountId: &device.ConnectedAccountId,
		},
	)
	require.NoError(t, err)
	assert.Len(t, devices, 3)

	devices, err = seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			ConnectedAccountIds: []string{device.ConnectedAccountId},
		},
	)
	require.NoError(t, err)
	assert.Len(t, devices, 3)

	devices, err = seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			DeviceType: seamgo.DeviceTypeAugustLock.Ptr(),
		},
	)
	require.NoError(t, err)
	assert.Len(t, devices, 2)

	devices, err = seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			DeviceTypes: []seamgo.DeviceType{
				seamgo.DeviceTypeAugustLock,
			},
		},
	)
	require.NoError(t, err)
	assert.Len(t, devices, 2)

	// query device with id
	deviceById, err := seam.Devices.Get(
		ctx,
		&seamgo.DevicesGetRequest{
				DeviceId: &device.DeviceId,
		},
)
require.NoError(t, err)
assert.NotNil(t, deviceById, "Device queried by ID should not be nil")
assert.Equal(t, device.DisplayName, deviceById.DisplayName)
assert.Equal(t, device.DeviceId, deviceById.DeviceId)

	// query device by name
	deviceByName, err := seam.Devices.Get(
		ctx,
		&seamgo.DevicesGetRequest{
			Name: &device.DisplayName,
		},
	)
	require.NoError(t, err)
	assert.NotNil(t, deviceByName, "Device queried by name should not be nil")
	assert.Equal(t, device.DisplayName, deviceByName.DisplayName)
	assert.Equal(t, device.DeviceId, deviceByName.DeviceId)

	// querying with both name and ID returns an error
	_, err = seam.Devices.Get(
		ctx,
		&seamgo.DevicesGetRequest{
				Name:     &device.DisplayName,
				DeviceId: &device.DeviceId,
		},
)
assert.Error(t, err, "Expected an error when querying with both name and ID")
assert.Contains(t, err.Error(), "Either 'device_id' or 'name' is required")

	locks, err := seam.Locks.List(
		ctx,
		nil,
	)
	require.NoError(t, err)
	assert.Len(t, locks, 3)

	lock, err := seam.Locks.Get(
		ctx,
		&seamgo.LocksGetRequest{
			DeviceId: &device.DeviceId,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, device.DeviceId, lock.DeviceId)
}

func getTestDevice(t *testing.T, seam *seamclient.Client) *seamgo.Device {
	devices, err := seam.Devices.List(context.Background(), nil)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(devices), 1)
	return devices[0]
}
