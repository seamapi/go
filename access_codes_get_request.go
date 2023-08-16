// This file was auto-generated by Fern from our API Definition.

package api

// AccessCodesGetRequest is an in-lined request used by the Get endpoint.
type AccessCodesGetRequest struct {
	DeviceId     *string `json:"device_id,omitempty"`
	AccessCodeId *string `json:"access_code_id,omitempty"`
	Code         *string `json:"code,omitempty"`
}