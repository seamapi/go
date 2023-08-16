// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type DevicesListDeviceProvidersRequestProviderCategory uint8

const (
	DevicesListDeviceProvidersRequestProviderCategoryStable DevicesListDeviceProvidersRequestProviderCategory = iota + 1
	DevicesListDeviceProvidersRequestProviderCategoryConsumerSmartlocks
)

func (d DevicesListDeviceProvidersRequestProviderCategory) String() string {
	switch d {
	default:
		return strconv.Itoa(int(d))
	case DevicesListDeviceProvidersRequestProviderCategoryStable:
		return "stable"
	case DevicesListDeviceProvidersRequestProviderCategoryConsumerSmartlocks:
		return "consumer_smartlocks"
	}
}

func (d DevicesListDeviceProvidersRequestProviderCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", d.String())), nil
}

func (d *DevicesListDeviceProvidersRequestProviderCategory) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "stable":
		value := DevicesListDeviceProvidersRequestProviderCategoryStable
		*d = value
	case "consumer_smartlocks":
		value := DevicesListDeviceProvidersRequestProviderCategoryConsumerSmartlocks
		*d = value
	}
	return nil
}