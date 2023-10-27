package integration

import (
	"context"
	"testing"

	seamgo "github.com/seamapi/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThermostats(t *testing.T) {
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	ctx := context.Background()
	thermostats, err := seam.Thermostats.List(ctx, nil)
	require.NoError(t, err)
	assert.Len(t, thermostats, 1)

	thermostat := thermostats[0]
	assert.Equal(t, seamgo.DeviceTypeEcobeeThermostat, thermostat.DeviceType)

	gotThermostat, err := seam.Thermostats.Get(
		ctx,
		&seamgo.ThermostatsGetRequest{
			DeviceId: &thermostat.DeviceId,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, thermostat.DeviceType, gotThermostat.DeviceType)

	// Update endpoint not served by Fake Seam.
	// _, err = seam.Thermostats.Update(
	// 	ctx,
	// 	&seamgo.ThermostatsUpdateRequest{
	// 		DeviceId: thermostat.DeviceId,
	// 		DefaultClimateSetting: &seamgo.ThermostatsUpdateRequestDefaultClimateSetting{
	// 			HvacModeSetting:        seamgo.HvacModeSettingCool.Ptr(),
	// 			CoolingSetPointCelsius: seamgo.Float64(20.0),
	// 			ManualOverrideAllowed:  seamgo.Bool(true),
	// 		},
	// 	},
	// )
	// require.NoError(t, err)

	// updatedThermostat, err := seam.Thermostats.Get(
	// 	ctx,
	// 	&seamgo.ThermostatsGetRequest{
	// 		DeviceId: &thermostat.DeviceId,
	// 	},
	// )
	// require.NoError(t, err)
	// assert.Equal(t, gotThermostat.DeviceType, updatedThermostat.DeviceType)
	// assert.NotNil(t, updatedThermostat)
	// assert.NotNil(t, updatedThermostat.Properties)
	// assert.NotNil(t, updatedThermostat.Properties.CurrentClimateSetting.CoolingSetPointCelsius)
	// assert.Equal(t, 20.0, *updatedThermostat.Properties.CurrentClimateSetting.CoolingSetPointCelsius)

	// Heat endpoint not served by Fake Seam.
	// _, err = seam.Thermostats.Heat(
	// 	ctx,
	// 	&seamgo.ThermostatsHeatRequest{
	// 		DeviceId:               thermostat.DeviceId,
	// 		HeatingSetPointCelsius: seamgo.Float64(18.0),
	// 	},
	// )
	// require.NoError(t, err)

	// heatedThermostat, err := seam.Thermostats.Get(
	// 	ctx,
	// 	&seamgo.ThermostatsGetRequest{
	// 		DeviceId: &thermostat.DeviceId,
	// 	},
	// )
	// require.NoError(t, err)
	// assert.NotNil(t, heatedThermostat.Properties)
	// assert.NotNil(t, heatedThermostat.Properties.CurrentClimateSetting)
	// assert.NotNil(t, heatedThermostat.Properties.CurrentClimateSetting.HvacModeSetting)
	// assert.Equal(t, "heat_cool", *heatedThermostat.Properties.CurrentClimateSetting.HvacModeSetting)
}
