// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
	time "time"
)

type LocksGetRequest struct {
	DeviceId *string `json:"device_id,omitempty" url:"device_id,omitempty"`
	Name     *string `json:"name,omitempty" url:"name,omitempty"`
}

type LocksListRequest struct {
	// List all devices owned by this connected account
	ConnectedAccountId  *string                                            `json:"connected_account_id,omitempty" url:"connected_account_id,omitempty"`
	ConnectedAccountIds []string                                           `json:"connected_account_ids,omitempty" url:"connected_account_ids,omitempty"`
	ConnectWebviewId    *string                                            `json:"connect_webview_id,omitempty" url:"connect_webview_id,omitempty"`
	DeviceType          *DeviceType                                        `json:"device_type,omitempty" url:"device_type,omitempty"`
	DeviceTypes         []DeviceType                                       `json:"device_types,omitempty" url:"device_types,omitempty"`
	Manufacturer        *Manufacturer                                      `json:"manufacturer,omitempty" url:"manufacturer,omitempty"`
	DeviceIds           []string                                           `json:"device_ids,omitempty" url:"device_ids,omitempty"`
	Limit               *float64                                           `json:"limit,omitempty" url:"limit,omitempty"`
	CreatedBefore       *time.Time                                         `json:"created_before,omitempty" url:"created_before,omitempty"`
	UserIdentifierKey   *string                                            `json:"user_identifier_key,omitempty" url:"user_identifier_key,omitempty"`
	CustomMetadataHas   map[string]*LocksListRequestCustomMetadataHasValue `json:"custom_metadata_has,omitempty" url:"custom_metadata_has,omitempty"`
}

func (l *LocksListRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler LocksListRequest
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*l = LocksListRequest(body)
	return nil
}

func (l *LocksListRequest) MarshalJSON() ([]byte, error) {
	type embed LocksListRequest
	var marshaler = struct {
		embed
		CreatedBefore *core.DateTime `json:"created_before,omitempty"`
	}{
		embed:         embed(*l),
		CreatedBefore: core.NewOptionalDateTime(l.CreatedBefore),
	}
	return json.Marshal(marshaler)
}

type LocksLockDoorRequest struct {
	DeviceId string `json:"device_id" url:"device_id"`
	Sync     *bool  `json:"sync,omitempty" url:"sync,omitempty"`
}

type LocksGetResponse struct {
	Lock   *Device `json:"lock,omitempty" url:"lock,omitempty"`
	Device *Device `json:"device,omitempty" url:"device,omitempty"`
	Ok     bool    `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (l *LocksGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler LocksGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LocksGetResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LocksGetResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type LocksListRequestCustomMetadataHasValue struct {
	typeName       string
	String         string
	Boolean        bool
	StringOptional *string
}

func NewLocksListRequestCustomMetadataHasValueFromString(value string) *LocksListRequestCustomMetadataHasValue {
	return &LocksListRequestCustomMetadataHasValue{typeName: "string", String: value}
}

func NewLocksListRequestCustomMetadataHasValueFromBoolean(value bool) *LocksListRequestCustomMetadataHasValue {
	return &LocksListRequestCustomMetadataHasValue{typeName: "boolean", Boolean: value}
}

func NewLocksListRequestCustomMetadataHasValueFromStringOptional(value *string) *LocksListRequestCustomMetadataHasValue {
	return &LocksListRequestCustomMetadataHasValue{typeName: "stringOptional", StringOptional: value}
}

func (l *LocksListRequestCustomMetadataHasValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		l.typeName = "string"
		l.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		l.typeName = "boolean"
		l.Boolean = valueBoolean
		return nil
	}
	var valueStringOptional *string
	if err := json.Unmarshal(data, &valueStringOptional); err == nil {
		l.typeName = "stringOptional"
		l.StringOptional = valueStringOptional
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LocksListRequestCustomMetadataHasValue) MarshalJSON() ([]byte, error) {
	switch l.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", l.typeName, l)
	case "string":
		return json.Marshal(l.String)
	case "boolean":
		return json.Marshal(l.Boolean)
	case "stringOptional":
		return json.Marshal(l.StringOptional)
	}
}

type LocksListRequestCustomMetadataHasValueVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
	VisitStringOptional(*string) error
}

func (l *LocksListRequestCustomMetadataHasValue) Accept(visitor LocksListRequestCustomMetadataHasValueVisitor) error {
	switch l.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", l.typeName, l)
	case "string":
		return visitor.VisitString(l.String)
	case "boolean":
		return visitor.VisitBoolean(l.Boolean)
	case "stringOptional":
		return visitor.VisitStringOptional(l.StringOptional)
	}
}

type LocksListResponse struct {
	Locks   []*Device `json:"locks,omitempty" url:"locks,omitempty"`
	Devices []*Device `json:"devices,omitempty" url:"devices,omitempty"`
	Ok      bool      `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (l *LocksListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler LocksListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LocksListResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LocksListResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type LocksLockDoorResponse struct {
	ActionAttempt *ActionAttempt `json:"action_attempt,omitempty" url:"action_attempt,omitempty"`
	Ok            bool           `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (l *LocksLockDoorResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler LocksLockDoorResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LocksLockDoorResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LocksLockDoorResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type LocksUnlockDoorResponse struct {
	ActionAttempt *ActionAttempt `json:"action_attempt,omitempty" url:"action_attempt,omitempty"`
	Ok            bool           `json:"ok" url:"ok"`

	_rawJSON json.RawMessage
}

func (l *LocksUnlockDoorResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler LocksUnlockDoorResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LocksUnlockDoorResponse(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *LocksUnlockDoorResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type LocksUnlockDoorRequest struct {
	DeviceId string `json:"device_id" url:"device_id"`
	Sync     *bool  `json:"sync,omitempty" url:"sync,omitempty"`
}
