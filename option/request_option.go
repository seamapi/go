// This file was auto-generated by Fern from our API Definition.

package option

import (
	core "github.com/seamapi/go/core"
	http "net/http"
)

// RequestOption adapts the behavior of an indivdual request.
type RequestOption = core.RequestOption

// WithBaseURL sets the base URL, overriding the default
// environment, if any.
func WithBaseURL(baseURL string) *core.BaseURLOption {
	return &core.BaseURLOption{
		BaseURL: baseURL,
	}
}

// WithHTTPClient uses the given HTTPClient to issue the request.
func WithHTTPClient(httpClient core.HTTPClient) *core.HTTPClientOption {
	return &core.HTTPClientOption{
		HTTPClient: httpClient,
	}
}

// WithHTTPHeader adds the given http.Header to the request.
func WithHTTPHeader(httpHeader http.Header) *core.HTTPHeaderOption {
	return &core.HTTPHeaderOption{
		// Clone the headers so they can't be modified after the option call.
		HTTPHeader: httpHeader.Clone(),
	}
}

// WithMaxAttempts configures the maximum number of retry attempts.
func WithMaxAttempts(attempts uint) *core.MaxAttemptsOption {
	return &core.MaxAttemptsOption{
		MaxAttempts: attempts,
	}
}

// WithApiKey sets the 'Authorization: Bearer <apiKey>' request header.
func WithApiKey(apiKey string) *core.ApiKeyOption {
	return &core.ApiKeyOption{
		ApiKey: apiKey,
	}
}

// WithSeamWorkspace sets the seamWorkspace request header.
func WithSeamWorkspace(seamWorkspace string) *core.SeamWorkspaceOption {
	return &core.SeamWorkspaceOption{
		SeamWorkspace: seamWorkspace,
	}
}

// WithSeamClientSessionToken sets the seamClientSessionToken request header.
func WithSeamClientSessionToken(seamClientSessionToken string) *core.SeamClientSessionTokenOption {
	return &core.SeamClientSessionTokenOption{
		SeamClientSessionToken: seamClientSessionToken,
	}
}

// WithClientSessionToken sets the clientSessionToken request header.
func WithClientSessionToken(clientSessionToken string) *core.ClientSessionTokenOption {
	return &core.ClientSessionTokenOption{
		ClientSessionToken: clientSessionToken,
	}
}
