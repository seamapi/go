// This file was auto-generated by Fern from our API Definition.

package api

// ClientSessionsCreateRequest is an in-lined request used by the Create endpoint.
type ClientSessionsCreateRequest struct {
	// <span style="white-space: nowrap">`non-empty`</span>
	UserIdentifierKey   *string   `json:"user_identifier_key,omitempty"`
	ConnectWebviewIds   *[]string `json:"connect_webview_ids,omitempty"`
	ConnectedAccountIds *[]string `json:"connected_account_ids,omitempty"`
}