// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type DevicesListRequestDeviceType uint8

const (
	DevicesListRequestDeviceTypeAkuvoxLock DevicesListRequestDeviceType = iota + 1
	DevicesListRequestDeviceTypeAugustLock
	DevicesListRequestDeviceTypeBrivoAccessPoint
	DevicesListRequestDeviceTypeButterflymxPanel
	DevicesListRequestDeviceTypeDoorkingLock
	DevicesListRequestDeviceTypeGenieDoor
	DevicesListRequestDeviceTypeIglooLock
	DevicesListRequestDeviceTypeLinearLock
	DevicesListRequestDeviceTypeLocklyLock
	DevicesListRequestDeviceTypeKwiksetLock
	DevicesListRequestDeviceTypeNukiLock
	DevicesListRequestDeviceTypeSaltoLock
	DevicesListRequestDeviceTypeSchlageLock
	DevicesListRequestDeviceTypeSeamRelay
	DevicesListRequestDeviceTypeSmartthingsLock
	DevicesListRequestDeviceTypeYaleLock
	DevicesListRequestDeviceTypeTwoNIntercom
	DevicesListRequestDeviceTypeControlbywebDevice
	DevicesListRequestDeviceTypeTtlockLock
	DevicesListRequestDeviceTypeIgloohomeLock
	DevicesListRequestDeviceTypeHubitatLock
	DevicesListRequestDeviceTypeNoiseawareActivityZone
	DevicesListRequestDeviceTypeMinutSensor
	DevicesListRequestDeviceTypeEcobeeThermostat
	DevicesListRequestDeviceTypeNestThermostat
)

func (d DevicesListRequestDeviceType) String() string {
	switch d {
	default:
		return strconv.Itoa(int(d))
	case DevicesListRequestDeviceTypeAkuvoxLock:
		return "akuvox_lock"
	case DevicesListRequestDeviceTypeAugustLock:
		return "august_lock"
	case DevicesListRequestDeviceTypeBrivoAccessPoint:
		return "brivo_access_point"
	case DevicesListRequestDeviceTypeButterflymxPanel:
		return "butterflymx_panel"
	case DevicesListRequestDeviceTypeDoorkingLock:
		return "doorking_lock"
	case DevicesListRequestDeviceTypeGenieDoor:
		return "genie_door"
	case DevicesListRequestDeviceTypeIglooLock:
		return "igloo_lock"
	case DevicesListRequestDeviceTypeLinearLock:
		return "linear_lock"
	case DevicesListRequestDeviceTypeLocklyLock:
		return "lockly_lock"
	case DevicesListRequestDeviceTypeKwiksetLock:
		return "kwikset_lock"
	case DevicesListRequestDeviceTypeNukiLock:
		return "nuki_lock"
	case DevicesListRequestDeviceTypeSaltoLock:
		return "salto_lock"
	case DevicesListRequestDeviceTypeSchlageLock:
		return "schlage_lock"
	case DevicesListRequestDeviceTypeSeamRelay:
		return "seam_relay"
	case DevicesListRequestDeviceTypeSmartthingsLock:
		return "smartthings_lock"
	case DevicesListRequestDeviceTypeYaleLock:
		return "yale_lock"
	case DevicesListRequestDeviceTypeTwoNIntercom:
		return "two_n_intercom"
	case DevicesListRequestDeviceTypeControlbywebDevice:
		return "controlbyweb_device"
	case DevicesListRequestDeviceTypeTtlockLock:
		return "ttlock_lock"
	case DevicesListRequestDeviceTypeIgloohomeLock:
		return "igloohome_lock"
	case DevicesListRequestDeviceTypeHubitatLock:
		return "hubitat_lock"
	case DevicesListRequestDeviceTypeNoiseawareActivityZone:
		return "noiseaware_activity_zone"
	case DevicesListRequestDeviceTypeMinutSensor:
		return "minut_sensor"
	case DevicesListRequestDeviceTypeEcobeeThermostat:
		return "ecobee_thermostat"
	case DevicesListRequestDeviceTypeNestThermostat:
		return "nest_thermostat"
	}
}

func (d DevicesListRequestDeviceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.String())), nil
}

func (d *DevicesListRequestDeviceType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "akuvox_lock":
		value := DevicesListRequestDeviceTypeAkuvoxLock
		*d = value
	case "august_lock":
		value := DevicesListRequestDeviceTypeAugustLock
		*d = value
	case "brivo_access_point":
		value := DevicesListRequestDeviceTypeBrivoAccessPoint
		*d = value
	case "butterflymx_panel":
		value := DevicesListRequestDeviceTypeButterflymxPanel
		*d = value
	case "doorking_lock":
		value := DevicesListRequestDeviceTypeDoorkingLock
		*d = value
	case "genie_door":
		value := DevicesListRequestDeviceTypeGenieDoor
		*d = value
	case "igloo_lock":
		value := DevicesListRequestDeviceTypeIglooLock
		*d = value
	case "linear_lock":
		value := DevicesListRequestDeviceTypeLinearLock
		*d = value
	case "lockly_lock":
		value := DevicesListRequestDeviceTypeLocklyLock
		*d = value
	case "kwikset_lock":
		value := DevicesListRequestDeviceTypeKwiksetLock
		*d = value
	case "nuki_lock":
		value := DevicesListRequestDeviceTypeNukiLock
		*d = value
	case "salto_lock":
		value := DevicesListRequestDeviceTypeSaltoLock
		*d = value
	case "schlage_lock":
		value := DevicesListRequestDeviceTypeSchlageLock
		*d = value
	case "seam_relay":
		value := DevicesListRequestDeviceTypeSeamRelay
		*d = value
	case "smartthings_lock":
		value := DevicesListRequestDeviceTypeSmartthingsLock
		*d = value
	case "yale_lock":
		value := DevicesListRequestDeviceTypeYaleLock
		*d = value
	case "two_n_intercom":
		value := DevicesListRequestDeviceTypeTwoNIntercom
		*d = value
	case "controlbyweb_device":
		value := DevicesListRequestDeviceTypeControlbywebDevice
		*d = value
	case "ttlock_lock":
		value := DevicesListRequestDeviceTypeTtlockLock
		*d = value
	case "igloohome_lock":
		value := DevicesListRequestDeviceTypeIgloohomeLock
		*d = value
	case "hubitat_lock":
		value := DevicesListRequestDeviceTypeHubitatLock
		*d = value
	case "noiseaware_activity_zone":
		value := DevicesListRequestDeviceTypeNoiseawareActivityZone
		*d = value
	case "minut_sensor":
		value := DevicesListRequestDeviceTypeMinutSensor
		*d = value
	case "ecobee_thermostat":
		value := DevicesListRequestDeviceTypeEcobeeThermostat
		*d = value
	case "nest_thermostat":
		value := DevicesListRequestDeviceTypeNestThermostat
		*d = value
	}
	return nil
}