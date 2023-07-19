// This file was auto-generated by Fern from our API Definition.

package api

// AccessCodesUpdatePutRequest is an in-lined request used by the AccessCodesUpdatePut endpoint.
type AccessCodesUpdatePutRequest struct {
	Name     *string `json:"name,omitempty"`
	StartsAt *string `json:"starts_at,omitempty"`
	EndsAt   *string `json:"ends_at,omitempty"`
	// <span style="white-space: nowrap">`<= 8 characters`</span>
	Code                    *string                          `json:"code,omitempty"`
	Sync                    *bool                            `json:"sync,omitempty"`
	AttemptForOfflineDevice *bool                            `json:"attempt_for_offline_device,omitempty"`
	AccessCodeId            string                           `json:"access_code_id"`
	DeviceId                *string                          `json:"device_id,omitempty"`
	Type                    *AccessCodesUpdatePutRequestType `json:"type,omitempty"`
}
