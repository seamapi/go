// This file was auto-generated by Fern from our API Definition.

package noisesensors

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
)

type SimulateTriggerNoiseThresholdRequest struct {
	DeviceId string `json:"device_id" url:"device_id"`
}

type SimulateTriggerNoiseThresholdResponse struct {
	Ok bool `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (s *SimulateTriggerNoiseThresholdResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SimulateTriggerNoiseThresholdResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SimulateTriggerNoiseThresholdResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SimulateTriggerNoiseThresholdResponse) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
