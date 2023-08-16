// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

// DevicesListRequest is an in-lined request used by the List endpoint.
type DevicesListRequest struct {
	ConnectedAccountId  *string                              `json:"connected_account_id,omitempty"`
	ConnectedAccountIds *[]string                            `json:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string                              `json:"connect_webview_id,omitempty"`
	DeviceType          *DevicesListRequestDeviceType        `json:"device_type,omitempty"`
	DeviceTypes         *[]DevicesListRequestDeviceTypesItem `json:"device_types,omitempty"`
	Manufacturer        *DevicesListRequestManufacturer      `json:"manufacturer,omitempty"`
	DeviceIds           *[]string                            `json:"device_ids,omitempty"`
	Limit               *float64                             `json:"limit,omitempty"`
	CreatedBefore       *time.Time                           `json:"created_before,omitempty"`
}