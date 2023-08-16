// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type ClimateSettingScheduleHvacModeSetting uint8

const (
	ClimateSettingScheduleHvacModeSettingOff ClimateSettingScheduleHvacModeSetting = iota + 1
	ClimateSettingScheduleHvacModeSettingHeat
	ClimateSettingScheduleHvacModeSettingCool
	ClimateSettingScheduleHvacModeSettingHeatcool
)

func (c ClimateSettingScheduleHvacModeSetting) String() string {
	switch c {
	default:
		return strconv.Itoa(int(c))
	case ClimateSettingScheduleHvacModeSettingOff:
		return "off"
	case ClimateSettingScheduleHvacModeSettingHeat:
		return "heat"
	case ClimateSettingScheduleHvacModeSettingCool:
		return "cool"
	case ClimateSettingScheduleHvacModeSettingHeatcool:
		return "heatcool"
	}
}

func (c ClimateSettingScheduleHvacModeSetting) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", c.String())), nil
}

func (c *ClimateSettingScheduleHvacModeSetting) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "off":
		value := ClimateSettingScheduleHvacModeSettingOff
		*c = value
	case "heat":
		value := ClimateSettingScheduleHvacModeSettingHeat
		*c = value
	case "cool":
		value := ClimateSettingScheduleHvacModeSettingCool
		*c = value
	case "heatcool":
		value := ClimateSettingScheduleHvacModeSettingHeatcool
		*c = value
	}
	return nil
}