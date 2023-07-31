// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type DeviceDeviceTypeDeviceDeviceType uint8

const (
	DeviceDeviceTypeDeviceDeviceTypeEcobeeThermostat DeviceDeviceTypeDeviceDeviceType = iota + 1
	DeviceDeviceTypeDeviceDeviceTypeNestThermostat
)

func (d DeviceDeviceTypeDeviceDeviceType) String() string {
	switch d {
	default:
		return strconv.Itoa(int(d))
	case DeviceDeviceTypeDeviceDeviceTypeEcobeeThermostat:
		return "ecobee_thermostat"
	case DeviceDeviceTypeDeviceDeviceTypeNestThermostat:
		return "nest_thermostat"
	}
}

func (d DeviceDeviceTypeDeviceDeviceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.String())), nil
}

func (d *DeviceDeviceTypeDeviceDeviceType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "ecobee_thermostat":
		value := DeviceDeviceTypeDeviceDeviceTypeEcobeeThermostat
		*d = value
	case "nest_thermostat":
		value := DeviceDeviceTypeDeviceDeviceTypeNestThermostat
		*d = value
	}
	return nil
}
