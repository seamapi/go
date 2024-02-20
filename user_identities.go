// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/seamapi/go/core"
)

type UserIdentitiesAddAcsUserRequest struct {
	UserIdentityId string `json:"user_identity_id"`
	AcsUserId      string `json:"acs_user_id"`
}

type UserIdentitiesCreateRequest struct {
	UserIdentityKey *string `json:"user_identity_key,omitempty"`
	EmailAddress    *string `json:"email_address,omitempty"`
	PhoneNumber     *string `json:"phone_number,omitempty"`
	FullName        *string `json:"full_name,omitempty"`
}

type UserIdentitiesDeleteRequest struct {
	UserIdentityId string `json:"user_identity_id"`
}

type UserIdentitiesGrantAccessToDeviceRequest struct {
	UserIdentityId string `json:"user_identity_id"`
	DeviceId       string `json:"device_id"`
}

type UserIdentitiesListAccessibleDevicesRequest struct {
	UserIdentityId string `json:"user_identity_id"`
}

type UserIdentitiesListAcsSystemsRequest struct {
	UserIdentityId string `json:"user_identity_id"`
}

type UserIdentitiesListAcsUsersRequest struct {
	UserIdentityId string `json:"user_identity_id"`
}

type UserIdentitiesRemoveAcsUserRequest struct {
	UserIdentityId string `json:"user_identity_id"`
	AcsUserId      string `json:"acs_user_id"`
}

type UserIdentitiesRevokeAccessToDeviceRequest struct {
	UserIdentityId string `json:"user_identity_id"`
	DeviceId       string `json:"device_id"`
}

type UserIdentitiesAddAcsUserResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesAddAcsUserResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesAddAcsUserResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesAddAcsUserResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesAddAcsUserResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesCreateResponse struct {
	UserIdentity *UserIdentitiesCreateResponseUserIdentity `json:"user_identity,omitempty"`
	Ok           bool                                      `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesCreateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesCreateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesCreateResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesCreateResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesDeleteResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesDeleteResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesDeleteResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesDeleteResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesDeleteResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesGetRequest struct {
	typeName                                string
	UserIdentitiesGetRequestUserIdentityId  *UserIdentitiesGetRequestUserIdentityId
	UserIdentitiesGetRequestUserIdentityKey *UserIdentitiesGetRequestUserIdentityKey
}

func NewUserIdentitiesGetRequestFromUserIdentitiesGetRequestUserIdentityId(value *UserIdentitiesGetRequestUserIdentityId) *UserIdentitiesGetRequest {
	return &UserIdentitiesGetRequest{typeName: "userIdentitiesGetRequestUserIdentityId", UserIdentitiesGetRequestUserIdentityId: value}
}

func NewUserIdentitiesGetRequestFromUserIdentitiesGetRequestUserIdentityKey(value *UserIdentitiesGetRequestUserIdentityKey) *UserIdentitiesGetRequest {
	return &UserIdentitiesGetRequest{typeName: "userIdentitiesGetRequestUserIdentityKey", UserIdentitiesGetRequestUserIdentityKey: value}
}

func (u *UserIdentitiesGetRequest) UnmarshalJSON(data []byte) error {
	valueUserIdentitiesGetRequestUserIdentityId := new(UserIdentitiesGetRequestUserIdentityId)
	if err := json.Unmarshal(data, &valueUserIdentitiesGetRequestUserIdentityId); err == nil {
		u.typeName = "userIdentitiesGetRequestUserIdentityId"
		u.UserIdentitiesGetRequestUserIdentityId = valueUserIdentitiesGetRequestUserIdentityId
		return nil
	}
	valueUserIdentitiesGetRequestUserIdentityKey := new(UserIdentitiesGetRequestUserIdentityKey)
	if err := json.Unmarshal(data, &valueUserIdentitiesGetRequestUserIdentityKey); err == nil {
		u.typeName = "userIdentitiesGetRequestUserIdentityKey"
		u.UserIdentitiesGetRequestUserIdentityKey = valueUserIdentitiesGetRequestUserIdentityKey
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, u)
}

func (u UserIdentitiesGetRequest) MarshalJSON() ([]byte, error) {
	switch u.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", u.typeName, u)
	case "userIdentitiesGetRequestUserIdentityId":
		return json.Marshal(u.UserIdentitiesGetRequestUserIdentityId)
	case "userIdentitiesGetRequestUserIdentityKey":
		return json.Marshal(u.UserIdentitiesGetRequestUserIdentityKey)
	}
}

type UserIdentitiesGetRequestVisitor interface {
	VisitUserIdentitiesGetRequestUserIdentityId(*UserIdentitiesGetRequestUserIdentityId) error
	VisitUserIdentitiesGetRequestUserIdentityKey(*UserIdentitiesGetRequestUserIdentityKey) error
}

func (u *UserIdentitiesGetRequest) Accept(visitor UserIdentitiesGetRequestVisitor) error {
	switch u.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", u.typeName, u)
	case "userIdentitiesGetRequestUserIdentityId":
		return visitor.VisitUserIdentitiesGetRequestUserIdentityId(u.UserIdentitiesGetRequestUserIdentityId)
	case "userIdentitiesGetRequestUserIdentityKey":
		return visitor.VisitUserIdentitiesGetRequestUserIdentityKey(u.UserIdentitiesGetRequestUserIdentityKey)
	}
}

type UserIdentitiesGetResponse struct {
	UserIdentity *UserIdentitiesGetResponseUserIdentity `json:"user_identity,omitempty"`
	Ok           bool                                   `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesGetResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesGetResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesGrantAccessToDeviceResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesGrantAccessToDeviceResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesGrantAccessToDeviceResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesGrantAccessToDeviceResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesGrantAccessToDeviceResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesListAccessibleDevicesResponse struct {
	AccessibleDevices []*Device `json:"accessible_devices,omitempty"`
	Ok                bool      `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesListAccessibleDevicesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesListAccessibleDevicesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesListAccessibleDevicesResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesListAccessibleDevicesResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesListAcsSystemsResponse struct {
	AcsSystems []*AcsSystem `json:"acs_systems,omitempty"`
	Ok         bool         `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesListAcsSystemsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesListAcsSystemsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesListAcsSystemsResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesListAcsSystemsResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesListAcsUsersResponse struct {
	AcsUsers []*AcsUser `json:"acs_users,omitempty"`
	Ok       bool       `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesListAcsUsersResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesListAcsUsersResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesListAcsUsersResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesListAcsUsersResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesListResponse struct {
	UserIdentities []*UserIdentitiesListResponseUserIdentitiesItem `json:"user_identities,omitempty"`
	Ok             bool                                            `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesListResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesListResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesRemoveAcsUserResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesRemoveAcsUserResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesRemoveAcsUserResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesRemoveAcsUserResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesRemoveAcsUserResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesRevokeAccessToDeviceResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesRevokeAccessToDeviceResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesRevokeAccessToDeviceResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesRevokeAccessToDeviceResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesRevokeAccessToDeviceResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesUpdateResponse struct {
	Ok bool `json:"ok"`

	_rawJSON json.RawMessage
}

func (u *UserIdentitiesUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserIdentitiesUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserIdentitiesUpdateResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserIdentitiesUpdateResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserIdentitiesUpdateRequest struct {
	UserIdentityId  string  `json:"user_identity_id"`
	UserIdentityKey *string `json:"user_identity_key,omitempty"`
	EmailAddress    *string `json:"email_address,omitempty"`
	PhoneNumber     *string `json:"phone_number,omitempty"`
	FullName        *string `json:"full_name,omitempty"`
}
