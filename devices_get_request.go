// This file was auto-generated by Fern from our API Definition.

package api

// DevicesGetRequest is an in-lined request used by the Get endpoint.
type DevicesGetRequest struct {
	DeviceId *string `json:"device_id,omitempty"`
	Name     *string `json:"name,omitempty"`
}