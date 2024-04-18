package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	seam "github.com/seamapi/go"
	seamclient "github.com/seamapi/go/client"
	"github.com/seamapi/go/option"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	t.Run("success (no retry)", func(t *testing.T) {
		handler := newAccessCodeHandler(t)
		server := httptest.NewServer(
			http.HandlerFunc(handler.Handle),
		)
		defer server.Close()

		client := seamclient.NewClient(
			option.WithBaseURL(server.URL),
		)

		_, err := client.AccessCodes.Create(
			context.Background(),
			&seam.AccessCodesCreateRequest{
				DeviceId: "test-device",
			},
		)
		assert.NoError(t, err)
		assert.Equal(t, 1, handler.requestsReceived, "Expected one request upon success")
	})

	t.Run("default retry", func(t *testing.T) {
		handler := newErrorHandler(t)
		server := httptest.NewServer(
			http.HandlerFunc(handler.Handle),
		)
		defer server.Close()

		client := seamclient.NewClient(
			option.WithBaseURL(server.URL),
		)

		_, err := client.AccessCodes.Create(
			context.Background(),
			&seam.AccessCodesCreateRequest{
				DeviceId: "test-device",
			},
		)
		assert.EqualError(t, err, "500: Internal Server Error")
		assert.Equal(t, 2, handler.requestsReceived, "Expected 2 requests to be made by default")

		_, err = client.AccessCodes.Create(
			context.Background(),
			&seam.AccessCodesCreateRequest{
				DeviceId: "test-device",
			},
			option.WithMaxAttempts(1),
		)
		assert.EqualError(t, err, "500: Internal Server Error")
		assert.Equal(t, 3, handler.requestsReceived, "Expected only one additional request to be made")
	})

	t.Run("disable retry", func(t *testing.T) {
		handler := newErrorHandler(t)
		server := httptest.NewServer(
			http.HandlerFunc(handler.Handle),
		)
		defer server.Close()

		client := seamclient.NewClient(
			option.WithBaseURL(server.URL),
			option.WithMaxAttempts(1),
		)

		_, err := client.AccessCodes.Create(
			context.Background(),
			&seam.AccessCodesCreateRequest{
				DeviceId: "test-device",
			},
		)
		assert.EqualError(t, err, "500: Internal Server Error")
		assert.Equal(t, 1, handler.requestsReceived, "Expected only one requests to be made")

		_, err = client.AccessCodes.Create(
			context.Background(),
			&seam.AccessCodesCreateRequest{
				DeviceId: "test-device",
			},
			option.WithMaxAttempts(2),
		)
		assert.EqualError(t, err, "500: Internal Server Error")
		assert.Equal(t, 3, handler.requestsReceived, "Expected two additional requests to be made")
	})
}

type accessCodeHandler struct {
	requestsReceived int
}

func newAccessCodeHandler(*testing.T) *accessCodeHandler {
	return &accessCodeHandler{}
}

func (t *accessCodeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	t.requestsReceived++
	bytes, err := json.Marshal(
		&seam.AccessCodesCreateResponse{
			Ok: true,
		},
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

type errorHandler struct {
	requestsReceived int
}

func newErrorHandler(*testing.T) *errorHandler {
	return &errorHandler{}
}

func (t *errorHandler) Handle(w http.ResponseWriter, r *http.Request) {
	t.requestsReceived++
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
}
