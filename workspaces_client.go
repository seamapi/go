// This file was auto-generated by Fern from our API Definition.

package api

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	core "github.com/seamapi/go/core"
	io "io"
	http "net/http"
)

type WorkspacesClient interface {
	Get(ctx context.Context) (*WorkspacesGetResponse, error)
	List(ctx context.Context) (*WorkspacesListResponse, error)
	ResetSandbox(ctx context.Context) (*WorkspacesResetSandboxResponse, error)
}

func NewWorkspacesClient(opts ...core.ClientOption) WorkspacesClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &workspacesClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type workspacesClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (w *workspacesClient) Get(ctx context.Context) (*WorkspacesGetResponse, error) {
	baseURL := "https://connect.getseam.com"
	if w.baseURL != "" {
		baseURL = w.baseURL
	}
	endpointURL := baseURL + "/" + "workspaces/get"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *WorkspacesGetResponse
	if err := core.DoRequest(
		ctx,
		w.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		w.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (w *workspacesClient) List(ctx context.Context) (*WorkspacesListResponse, error) {
	baseURL := "https://connect.getseam.com"
	if w.baseURL != "" {
		baseURL = w.baseURL
	}
	endpointURL := baseURL + "/" + "workspaces/list"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *WorkspacesListResponse
	if err := core.DoRequest(
		ctx,
		w.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		w.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (w *workspacesClient) ResetSandbox(ctx context.Context) (*WorkspacesResetSandboxResponse, error) {
	baseURL := "https://connect.getseam.com"
	if w.baseURL != "" {
		baseURL = w.baseURL
	}
	endpointURL := baseURL + "/" + "workspaces/reset_sandbox"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 401:
			value := new(UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response *WorkspacesResetSandboxResponse
	if err := core.DoRequest(
		ctx,
		w.httpClient,
		endpointURL,
		http.MethodPost,
		nil,
		&response,
		false,
		w.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}