// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type ThermostatsClimateSettingSchedulesCreateRequestHvacModeSetting uint8

const (
	ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingOff ThermostatsClimateSettingSchedulesCreateRequestHvacModeSetting = iota + 1
	ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingHeat
	ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingCool
	ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingHeatcool
)

func (t ThermostatsClimateSettingSchedulesCreateRequestHvacModeSetting) String() string {
	switch t {
	default:
		return strconv.Itoa(int(t))
	case ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingOff:
		return "off"
	case ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingHeat:
		return "heat"
	case ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingCool:
		return "cool"
	case ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingHeatcool:
		return "heatcool"
	}
}

func (t ThermostatsClimateSettingSchedulesCreateRequestHvacModeSetting) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", t.String())), nil
}

func (t *ThermostatsClimateSettingSchedulesCreateRequestHvacModeSetting) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "off":
		value := ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingOff
		*t = value
	case "heat":
		value := ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingHeat
		*t = value
	case "cool":
		value := ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingCool
		*t = value
	case "heatcool":
		value := ThermostatsClimateSettingSchedulesCreateRequestHvacModeSettingHeatcool
		*t = value
	}
	return nil
}
