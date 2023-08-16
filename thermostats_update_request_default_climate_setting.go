// This file was auto-generated by Fern from our API Definition.

package api

type ThermostatsUpdateRequestDefaultClimateSetting struct {
	AutomaticHeatingEnabled   *bool                                                         `json:"automatic_heating_enabled,omitempty"`
	AutomaticCoolingEnabled   *bool                                                         `json:"automatic_cooling_enabled,omitempty"`
	HvacModeSetting           *ThermostatsUpdateRequestDefaultClimateSettingHvacModeSetting `json:"hvac_mode_setting,omitempty"`
	CoolingSetPointCelsius    *float64                                                      `json:"cooling_set_point_celsius,omitempty"`
	HeatingSetPointCelsius    *float64                                                      `json:"heating_set_point_celsius,omitempty"`
	CoolingSetPointFahrenheit *float64                                                      `json:"cooling_set_point_fahrenheit,omitempty"`
	HeatingSetPointFahrenheit *float64                                                      `json:"heating_set_point_fahrenheit,omitempty"`
	ManualOverrideAllowed     *bool                                                         `json:"manual_override_allowed,omitempty"`
}