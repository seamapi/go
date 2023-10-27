// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

type LocksGetRequest struct {
	DeviceId *string `json:"device_id,omitempty"`
	Name     *string `json:"name,omitempty"`
}

type LocksListRequest struct {
	ConnectedAccountId  *string       `json:"connected_account_id,omitempty"`
	ConnectedAccountIds []string      `json:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string       `json:"connect_webview_id,omitempty"`
	DeviceType          *DeviceType   `json:"device_type,omitempty"`
	DeviceTypes         []DeviceType  `json:"device_types,omitempty"`
	Manufacturer        *Manufacturer `json:"manufacturer,omitempty"`
	DeviceIds           []string      `json:"device_ids,omitempty"`
	Limit               *float64      `json:"limit,omitempty"`
	CreatedBefore       *time.Time    `json:"created_before,omitempty"`
}

type LocksLockDoorRequest struct {
	DeviceId string `json:"device_id"`
	Sync     *bool  `json:"sync,omitempty"`
}

type LocksGetResponse struct {
	Lock   interface{} `json:"lock,omitempty"`
	Device *Device     `json:"device,omitempty"`
	Ok     bool        `json:"ok"`
}

type LocksListResponse struct {
	Locks   interface{} `json:"locks,omitempty"`
	Devices []*Device   `json:"devices,omitempty"`
	Ok      bool        `json:"ok"`
}

type LocksLockDoorResponse struct {
	ActionAttempt *ActionAttempt `json:"action_attempt,omitempty"`
	Ok            bool           `json:"ok"`
}

type LocksUnlockDoorResponse struct {
	ActionAttempt *ActionAttempt `json:"action_attempt,omitempty"`
	Ok            bool           `json:"ok"`
}

type LocksUnlockDoorRequest struct {
	DeviceId string `json:"device_id"`
	Sync     *bool  `json:"sync,omitempty"`
}