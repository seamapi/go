// This file was auto-generated by Fern from our API Definition.

package noisethresholds

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	seamapigo "github.com/seamapi/go"
	core "github.com/seamapi/go/core"
	noisesensors "github.com/seamapi/go/noisesensors"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

func (c *Client) Create(ctx context.Context, request *noisesensors.NoiseThresholdsCreateRequest) (*seamapigo.ActionAttempt, error) {
	baseURL := "https://connect.getseam.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "noise_sensors/noise_thresholds/create"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(seamapigo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(seamapigo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *seamapigo.NoiseThresholdsCreateResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return nil, err
	}
	return response.ActionAttempt, nil
}

func (c *Client) Delete(ctx context.Context, request *noisesensors.NoiseThresholdsDeleteRequest) (*seamapigo.ActionAttempt, error) {
	baseURL := "https://connect.getseam.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "noise_sensors/noise_thresholds/delete"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(seamapigo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(seamapigo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *seamapigo.NoiseThresholdsDeleteResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodDelete,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return nil, err
	}
	return response.ActionAttempt, nil
}

func (c *Client) Get(ctx context.Context, request *noisesensors.NoiseThresholdsGetRequest) (*seamapigo.NoiseThreshold, error) {
	baseURL := "https://connect.getseam.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "noise_sensors/noise_thresholds/get"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(seamapigo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(seamapigo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *seamapigo.NoiseThresholdsGetResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return nil, err
	}
	return response.NoiseThreshold, nil
}

func (c *Client) List(ctx context.Context, request *noisesensors.NoiseThresholdsListRequest) ([]*seamapigo.NoiseThreshold, error) {
	baseURL := "https://connect.getseam.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "noise_sensors/noise_thresholds/list"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(seamapigo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(seamapigo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *seamapigo.NoiseThresholdsListResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return nil, err
	}
	return response.NoiseThresholds, nil
}

func (c *Client) Update(ctx context.Context, request *noisesensors.NoiseThresholdsUpdateRequest) (*seamapigo.ActionAttempt, error) {
	baseURL := "https://connect.getseam.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "noise_sensors/noise_thresholds/update"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(seamapigo.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 401:
			value := new(seamapigo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *seamapigo.NoiseThresholdsUpdateResponse
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return nil, err
	}
	return response.ActionAttempt, nil
}