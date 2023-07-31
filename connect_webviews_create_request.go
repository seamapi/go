// This file was auto-generated by Fern from our API Definition.

package api

// ConnectWebviewsCreateRequest is an in-lined request used by the Create endpoint.
type ConnectWebviewsCreateRequest struct {
	DeviceSelectionMode      *ConnectWebviewsCreateRequestDeviceSelectionMode             `json:"device_selection_mode,omitempty"`
	CustomRedirectUrl        *string                                                      `json:"custom_redirect_url,omitempty"`
	CustomRedirectFailureUrl *string                                                      `json:"custom_redirect_failure_url,omitempty"`
	AcceptedProviders        *[]ConnectWebviewsCreateRequestAcceptedProvidersItem         `json:"accepted_providers,omitempty"`
	ProviderCategory         *ConnectWebviewsCreateRequestProviderCategory                `json:"provider_category,omitempty"`
	CustomMetadata           *map[string]*ConnectWebviewsCreateRequestCustomMetadataValue `json:"custom_metadata,omitempty"`
}
