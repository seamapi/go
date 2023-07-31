// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type DevicesUnmanagedListRequestManufacturer uint8

const (
	DevicesUnmanagedListRequestManufacturerAkuvox DevicesUnmanagedListRequestManufacturer = iota + 1
	DevicesUnmanagedListRequestManufacturerAugust
	DevicesUnmanagedListRequestManufacturerBrivo
	DevicesUnmanagedListRequestManufacturerButterflymx
	DevicesUnmanagedListRequestManufacturerDoorking
	DevicesUnmanagedListRequestManufacturerGenie
	DevicesUnmanagedListRequestManufacturerIgloo
	DevicesUnmanagedListRequestManufacturerKeywe
	DevicesUnmanagedListRequestManufacturerKwikset
	DevicesUnmanagedListRequestManufacturerLinear
	DevicesUnmanagedListRequestManufacturerLockly
	DevicesUnmanagedListRequestManufacturerNuki
	DevicesUnmanagedListRequestManufacturerPhilia
	DevicesUnmanagedListRequestManufacturerSalto
	DevicesUnmanagedListRequestManufacturerSamsung
	DevicesUnmanagedListRequestManufacturerSchlage
	DevicesUnmanagedListRequestManufacturerSeam
	DevicesUnmanagedListRequestManufacturerUnknown
	DevicesUnmanagedListRequestManufacturerYale
	DevicesUnmanagedListRequestManufacturerMinut
	DevicesUnmanagedListRequestManufacturerTwoN
	DevicesUnmanagedListRequestManufacturerTtlock
	DevicesUnmanagedListRequestManufacturerNest
	DevicesUnmanagedListRequestManufacturerIgloohome
	DevicesUnmanagedListRequestManufacturerEcobee
	DevicesUnmanagedListRequestManufacturerHubitat
)

func (d DevicesUnmanagedListRequestManufacturer) String() string {
	switch d {
	default:
		return strconv.Itoa(int(d))
	case DevicesUnmanagedListRequestManufacturerAkuvox:
		return "akuvox"
	case DevicesUnmanagedListRequestManufacturerAugust:
		return "august"
	case DevicesUnmanagedListRequestManufacturerBrivo:
		return "brivo"
	case DevicesUnmanagedListRequestManufacturerButterflymx:
		return "butterflymx"
	case DevicesUnmanagedListRequestManufacturerDoorking:
		return "doorking"
	case DevicesUnmanagedListRequestManufacturerGenie:
		return "genie"
	case DevicesUnmanagedListRequestManufacturerIgloo:
		return "igloo"
	case DevicesUnmanagedListRequestManufacturerKeywe:
		return "keywe"
	case DevicesUnmanagedListRequestManufacturerKwikset:
		return "kwikset"
	case DevicesUnmanagedListRequestManufacturerLinear:
		return "linear"
	case DevicesUnmanagedListRequestManufacturerLockly:
		return "lockly"
	case DevicesUnmanagedListRequestManufacturerNuki:
		return "nuki"
	case DevicesUnmanagedListRequestManufacturerPhilia:
		return "philia"
	case DevicesUnmanagedListRequestManufacturerSalto:
		return "salto"
	case DevicesUnmanagedListRequestManufacturerSamsung:
		return "samsung"
	case DevicesUnmanagedListRequestManufacturerSchlage:
		return "schlage"
	case DevicesUnmanagedListRequestManufacturerSeam:
		return "seam"
	case DevicesUnmanagedListRequestManufacturerUnknown:
		return "unknown"
	case DevicesUnmanagedListRequestManufacturerYale:
		return "yale"
	case DevicesUnmanagedListRequestManufacturerMinut:
		return "minut"
	case DevicesUnmanagedListRequestManufacturerTwoN:
		return "two_n"
	case DevicesUnmanagedListRequestManufacturerTtlock:
		return "ttlock"
	case DevicesUnmanagedListRequestManufacturerNest:
		return "nest"
	case DevicesUnmanagedListRequestManufacturerIgloohome:
		return "igloohome"
	case DevicesUnmanagedListRequestManufacturerEcobee:
		return "ecobee"
	case DevicesUnmanagedListRequestManufacturerHubitat:
		return "hubitat"
	}
}

func (d DevicesUnmanagedListRequestManufacturer) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.String())), nil
}

func (d *DevicesUnmanagedListRequestManufacturer) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "akuvox":
		value := DevicesUnmanagedListRequestManufacturerAkuvox
		*d = value
	case "august":
		value := DevicesUnmanagedListRequestManufacturerAugust
		*d = value
	case "brivo":
		value := DevicesUnmanagedListRequestManufacturerBrivo
		*d = value
	case "butterflymx":
		value := DevicesUnmanagedListRequestManufacturerButterflymx
		*d = value
	case "doorking":
		value := DevicesUnmanagedListRequestManufacturerDoorking
		*d = value
	case "genie":
		value := DevicesUnmanagedListRequestManufacturerGenie
		*d = value
	case "igloo":
		value := DevicesUnmanagedListRequestManufacturerIgloo
		*d = value
	case "keywe":
		value := DevicesUnmanagedListRequestManufacturerKeywe
		*d = value
	case "kwikset":
		value := DevicesUnmanagedListRequestManufacturerKwikset
		*d = value
	case "linear":
		value := DevicesUnmanagedListRequestManufacturerLinear
		*d = value
	case "lockly":
		value := DevicesUnmanagedListRequestManufacturerLockly
		*d = value
	case "nuki":
		value := DevicesUnmanagedListRequestManufacturerNuki
		*d = value
	case "philia":
		value := DevicesUnmanagedListRequestManufacturerPhilia
		*d = value
	case "salto":
		value := DevicesUnmanagedListRequestManufacturerSalto
		*d = value
	case "samsung":
		value := DevicesUnmanagedListRequestManufacturerSamsung
		*d = value
	case "schlage":
		value := DevicesUnmanagedListRequestManufacturerSchlage
		*d = value
	case "seam":
		value := DevicesUnmanagedListRequestManufacturerSeam
		*d = value
	case "unknown":
		value := DevicesUnmanagedListRequestManufacturerUnknown
		*d = value
	case "yale":
		value := DevicesUnmanagedListRequestManufacturerYale
		*d = value
	case "minut":
		value := DevicesUnmanagedListRequestManufacturerMinut
		*d = value
	case "two_n":
		value := DevicesUnmanagedListRequestManufacturerTwoN
		*d = value
	case "ttlock":
		value := DevicesUnmanagedListRequestManufacturerTtlock
		*d = value
	case "nest":
		value := DevicesUnmanagedListRequestManufacturerNest
		*d = value
	case "igloohome":
		value := DevicesUnmanagedListRequestManufacturerIgloohome
		*d = value
	case "ecobee":
		value := DevicesUnmanagedListRequestManufacturerEcobee
		*d = value
	case "hubitat":
		value := DevicesUnmanagedListRequestManufacturerHubitat
		*d = value
	}
	return nil
}
