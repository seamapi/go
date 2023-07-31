// This file was auto-generated by Fern from our API Definition.

package devices

// DevicesUnmanagedUpdateRequest is an in-lined request used by the Update endpoint.
type DevicesUnmanagedUpdateRequest struct {
	DeviceId  string `json:"device_id"`
	IsManaged string `json:"is_managed"`
}
