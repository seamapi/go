// This file was auto-generated by Fern from our API Definition.

package api

import (
	time "time"
)

type UnmanagedDevice struct {
	DeviceId              string                                     `json:"device_id"`
	DeviceType            UnmanagedDeviceDeviceType                  `json:"device_type,omitempty"`
	ConnectedAccountId    string                                     `json:"connected_account_id"`
	CapabilitiesSupported []UnmanagedDeviceCapabilitiesSupportedItem `json:"capabilities_supported,omitempty"`
	WorkspaceId           string                                     `json:"workspace_id"`
	Errors                []*UnmanagedDeviceErrorsItem               `json:"errors,omitempty"`
	Warnings              []*UnmanagedDeviceWarningsItem             `json:"warnings,omitempty"`
	CreatedAt             time.Time                                  `json:"created_at"`
	IsManaged             string                                     `json:"is_managed"`
	Properties            *UnmanagedDeviceProperties                 `json:"properties,omitempty"`
}