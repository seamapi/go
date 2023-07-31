// This file was auto-generated by Fern from our API Definition.

package thermostats

import (
	seamapigo "github.com/seamapi/go"
)

// ThermostatsClimateSettingSchedulesUpdateRequest is an in-lined request used by the Update endpoint.
type ThermostatsClimateSettingSchedulesUpdateRequest struct {
	ClimateSettingScheduleId  string                                                                    `json:"climate_setting_schedule_id"`
	ScheduleType              *string                                                                   `json:"schedule_type,omitempty"`
	Name                      *string                                                                   `json:"name,omitempty"`
	ScheduleStartsAt          *string                                                                   `json:"schedule_starts_at,omitempty"`
	ScheduleEndsAt            *string                                                                   `json:"schedule_ends_at,omitempty"`
	AutomaticHeatingEnabled   *bool                                                                     `json:"automatic_heating_enabled,omitempty"`
	AutomaticCoolingEnabled   *bool                                                                     `json:"automatic_cooling_enabled,omitempty"`
	HvacModeSetting           *seamapigo.ThermostatsClimateSettingSchedulesUpdateRequestHvacModeSetting `json:"hvac_mode_setting,omitempty"`
	CoolingSetPointCelsius    *float64                                                                  `json:"cooling_set_point_celsius,omitempty"`
	HeatingSetPointCelsius    *float64                                                                  `json:"heating_set_point_celsius,omitempty"`
	CoolingSetPointFahrenheit *float64                                                                  `json:"cooling_set_point_fahrenheit,omitempty"`
	HeatingSetPointFahrenheit *float64                                                                  `json:"heating_set_point_fahrenheit,omitempty"`
	ManualOverrideAllowed     *bool                                                                     `json:"manual_override_allowed,omitempty"`
}
