// This file was auto-generated by Fern from our API Definition.

package api

type ServiceHealth struct {
	Service     string              `json:"service"`
	Status      ServiceHealthStatus `json:"status,omitempty"`
	Description string              `json:"description"`
}
