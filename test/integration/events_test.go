package integration

import (
	"context"
	"testing"

	seamgo "github.com/seamapi/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	since                    = "2021-01-01T00:00:00.000Z"
	eventTypeDeviceConnected = "device.connected"
	fakeEventID              = "00000000-0000-0000-0000-000000000000"
)

func TestEvents(t *testing.T) {
	t.Skip("Fake Seam does not support events endpoints")
	t.Parallel()

	seam, cleanup := newFakeSeam(t)
	defer cleanup()

	ctx := context.Background()
	events, err := seam.Events.List(
		ctx,
		&seamgo.EventsListRequest{
			Since: seamgo.String(since),
		},
	)
	require.NoError(t, err)
	assert.Len(t, events, 1)

	var connectedEvent *seamgo.Event
	for _, event := range events {
		if event.EventType == eventTypeDeviceConnected {
			connectedEvent = event
			break
		}
	}
	require.NotNil(t, connectedEvent)

	eventByID, err := seam.Events.Get(
		ctx,
		&seamgo.EventsGetRequest{
			EventId: &connectedEvent.EventId,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, connectedEvent.EventId, eventByID.EventId)

	eventByType, err := seam.Events.Get(
		ctx,
		&seamgo.EventsGetRequest{
			EventType: &connectedEvent.EventType,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, connectedEvent.EventType, eventByType.EventType)

	eventByDeviceID, err := seam.Events.Get(
		ctx,
		&seamgo.EventsGetRequest{
			DeviceId: connectedEvent.DeviceId,
		},
	)
	require.NoError(t, err)
	assert.Equal(t, connectedEvent.DeviceId, eventByDeviceID.DeviceId)

	emptyEvent, err := seam.Events.Get(
		ctx,
		&seamgo.EventsGetRequest{
			EventId: seamgo.String(fakeEventID),
		},
	)
	require.NoError(t, err)
	assert.Empty(t, emptyEvent)
}
