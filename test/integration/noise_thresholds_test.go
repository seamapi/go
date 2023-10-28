package integration

import (
	"context"
	"testing"

	seamgo "github.com/seamapi/go"
	seamclient "github.com/seamapi/go/client"
	"github.com/seamapi/go/noisesensors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	quietHoursThresholdName = "builtin_quiet_hours"
	startsDailyAt           = "20:00:00[America/Los_Angeles]"
	endsDailyAt             = "08:00:00[America/Los_Angeles]"
	noiseThresholdDecibels  = 75.0
)

func TestNoiseThresholds(t *testing.T) {
	t.Skip("Fake Seam does not support noise thresholds endpoints")
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	ctx := context.Background()
	device := getTestDevice(t, seam)

	noiseThresholds := getTestNoiseThresholdsForDevice(t, seam, device.DeviceId)
	assert.Len(t, noiseThresholds, 2)

	var quietHoursThreshold *seamgo.NoiseThreshold
	for _, noiseThreshold := range noiseThresholds {
		if noiseThreshold.Name == quietHoursThresholdName {
			quietHoursThreshold = noiseThreshold
			break
		}
	}
	require.NotNil(t, quietHoursThreshold)

	_, err := seam.NoiseSensors.NoiseThresholds.Delete(
		ctx,
		&noisesensors.NoiseThresholdsDeleteRequest{
			NoiseThresholdId: quietHoursThreshold.NoiseThresholdId,
			DeviceId:         device.DeviceId,
		},
	)
	require.NoError(t, err)

	noiseThresholds = getTestNoiseThresholdsForDevice(t, seam, device.DeviceId)
	assert.Len(t, noiseThresholds, 1)

	_, err = seam.NoiseSensors.NoiseThresholds.Create(
		context.Background(),
		&noisesensors.NoiseThresholdsCreateRequest{
			DeviceId:               device.DeviceId,
			StartsDailyAt:          startsDailyAt,
			EndsDailyAt:            endsDailyAt,
			NoiseThresholdDecibels: seamgo.Float64(noiseThresholdDecibels),
		},
	)
	require.NoError(t, err)

	noiseThresholds = getTestNoiseThresholdsForDevice(t, seam, device.DeviceId)
	assert.Len(t, noiseThresholds, 2)
}

func getTestNoiseThresholdsForDevice(
	t *testing.T,
	seam *seamclient.Client,
	deviceID string,
) []*seamgo.NoiseThreshold {
	noiseThresholds, err := seam.NoiseSensors.NoiseThresholds.List(
		context.Background(),
		&noisesensors.NoiseThresholdsListRequest{
			DeviceId: deviceID,
		},
	)
	require.NoError(t, err)
	return noiseThresholds
}
