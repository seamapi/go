// This file was auto-generated by Fern from our API Definition.

package health

import (
	json "encoding/json"
	fmt "fmt"
	seamapigo "github.com/seamapi/go"
	core "github.com/seamapi/go/core"
)

type ServiceByServiceNameRequest struct {
	ServiceName string `json:"service_name" url:"service_name"`
}

type ServiceByServiceNameResponse struct {
	Ok                      bool                     `json:"ok" url:"ok"`
	LastServiceEvaluationAt string                   `json:"last_service_evaluation_at" url:"last_service_evaluation_at"`
	ServiceHealth           *seamapigo.ServiceHealth `json:"service_health,omitempty" url:"service_health,omitempty"`

	_rawJSON json.RawMessage
}

func (s *ServiceByServiceNameResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ServiceByServiceNameResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ServiceByServiceNameResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ServiceByServiceNameResponse) String() string {
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
