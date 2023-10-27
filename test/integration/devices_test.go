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
	assert.Nil(t, device.Properties.AugustMetadata)

	devices, err := seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			ConnectedAccountId: &device.ConnectedAccountId,
		},
	)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(devices), 1)

	devices, err = seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			ConnectedAccountIds: []string{device.ConnectedAccountId},
		},
	)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(devices), 1)

	devices, err = seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			DeviceType: seamgo.DeviceTypeAugustLock.Ptr(),
		},
	)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(devices), 1)

	devices, err = seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			DeviceTypes: []seamgo.DeviceType{
				seamgo.DeviceTypeAugustLock,
			},
		},
	)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(devices), 1)

	devices, err = seam.Devices.List(
		ctx,
		&seamgo.DevicesListRequest{
			Manufacturer: seamgo.ManufacturerAugust.Ptr(),
		},
	)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(devices), 1)

	deviceWithName, err := seam.Devices.Get(
		ctx,
		&seamgo.DevicesGetRequest{
			Name:     &device.Properties.Name,
			DeviceId: &device.DeviceId,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, device.Properties.Name, deviceWithName.Properties.Name)

	locks, err := seam.Locks.List(
		ctx,
		nil,
	)
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(locks), 1)

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
