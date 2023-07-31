// This file was auto-generated by Fern from our API Definition.

package api

// AccessCodesListRequest is an in-lined request used by the List endpoint.
type AccessCodesListRequest struct {
	DeviceId      string    `json:"device_id"`
	AccessCodeIds *[]string `json:"access_code_ids,omitempty"`
}
