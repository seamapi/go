// This file was auto-generated by Fern from our API Definition.

package api

import (
	fmt "fmt"
)

type AccessCodesCreateRequest struct {
	DeviceId                string  `json:"device_id"`
	Name                    *string `json:"name,omitempty"`
	StartsAt                *string `json:"starts_at,omitempty"`
	EndsAt                  *string `json:"ends_at,omitempty"`
	Code                    *string `json:"code,omitempty"`
	Sync                    *bool   `json:"sync,omitempty"`
	AttemptForOfflineDevice *bool   `json:"attempt_for_offline_device,omitempty"`
	CommonCodeKey           *string `json:"common_code_key,omitempty"`
	PreferNativeScheduling  *bool   `json:"prefer_native_scheduling,omitempty"`
	UseBackupAccessCodePool *bool   `json:"use_backup_access_code_pool,omitempty"`
}

type AccessCodesCreateMultipleRequest struct {
	DeviceIds                      []string                                                        `json:"device_ids,omitempty"`
	BehaviorWhenCodeCannotBeShared *AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared `json:"behavior_when_code_cannot_be_shared,omitempty"`
	Name                           *string                                                         `json:"name,omitempty"`
	StartsAt                       *string                                                         `json:"starts_at,omitempty"`
	EndsAt                         *string                                                         `json:"ends_at,omitempty"`
	Code                           *string                                                         `json:"code,omitempty"`
	AttemptForOfflineDevice        *bool                                                           `json:"attempt_for_offline_device,omitempty"`
	PreferNativeScheduling         *bool                                                           `json:"prefer_native_scheduling,omitempty"`
	UseBackupAccessCodePool        *bool                                                           `json:"use_backup_access_code_pool,omitempty"`
}

type AccessCodesDeleteRequest struct {
	DeviceId     *string `json:"device_id,omitempty"`
	AccessCodeId string  `json:"access_code_id"`
	Sync         *bool   `json:"sync,omitempty"`
}

type AccessCodesGetRequest struct {
	DeviceId     *string `json:"device_id,omitempty"`
	AccessCodeId *string `json:"access_code_id,omitempty"`
	Code         *string `json:"code,omitempty"`
}

type AccessCodesListRequest struct {
	DeviceId      string   `json:"device_id"`
	AccessCodeIds []string `json:"access_code_ids,omitempty"`
}

type AccessCodesPullBackupAccessCodeRequest struct {
	AccessCodeId string `json:"access_code_id"`
}

type AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared string

const (
	AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedThrow            AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared = "throw"
	AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedCreateRandomCode AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared = "create_random_code"
)

func NewAccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedFromString(s string) (AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared, error) {
	switch s {
	case "throw":
		return AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedThrow, nil
	case "create_random_code":
		return AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedCreateRandomCode, nil
	}
	var t AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared) Ptr() *AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared {
	return &a
}

type AccessCodesCreateMultipleResponse struct {
	AccessCodes []*AccessCode `json:"access_codes,omitempty"`
	Ok          bool          `json:"ok"`
}

type AccessCodesCreateResponse struct {
	ActionAttempt *ActionAttempt `json:"action_attempt,omitempty"`
	AccessCode    *AccessCode    `json:"access_code,omitempty"`
	Ok            bool           `json:"ok"`
}

type AccessCodesDeleteResponse struct {
	ActionAttempt *ActionAttempt `json:"action_attempt,omitempty"`
	Ok            bool           `json:"ok"`
}

type AccessCodesGetResponse struct {
	AccessCode *AccessCode `json:"access_code,omitempty"`
	Ok         bool        `json:"ok"`
}

type AccessCodesListResponse struct {
	AccessCodes []*AccessCode `json:"access_codes,omitempty"`
	Ok          bool          `json:"ok"`
}

type AccessCodesPullBackupAccessCodeResponse struct {
	BackupAccessCode *AccessCode `json:"backup_access_code,omitempty"`
	Ok               bool        `json:"ok"`
}

type AccessCodesUpdateRequestType string

const (
	AccessCodesUpdateRequestTypeOngoing   AccessCodesUpdateRequestType = "ongoing"
	AccessCodesUpdateRequestTypeTimeBound AccessCodesUpdateRequestType = "time_bound"
)

func NewAccessCodesUpdateRequestTypeFromString(s string) (AccessCodesUpdateRequestType, error) {
	switch s {
	case "ongoing":
		return AccessCodesUpdateRequestTypeOngoing, nil
	case "time_bound":
		return AccessCodesUpdateRequestTypeTimeBound, nil
	}
	var t AccessCodesUpdateRequestType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a AccessCodesUpdateRequestType) Ptr() *AccessCodesUpdateRequestType {
	return &a
}

type AccessCodesUpdateResponse struct {
	ActionAttempt *ActionAttempt `json:"action_attempt,omitempty"`
	Ok            bool           `json:"ok"`
}

type AccessCodesUpdateRequest struct {
	Name                    *string                       `json:"name,omitempty"`
	StartsAt                *string                       `json:"starts_at,omitempty"`
	EndsAt                  *string                       `json:"ends_at,omitempty"`
	Code                    *string                       `json:"code,omitempty"`
	Sync                    *bool                         `json:"sync,omitempty"`
	AttemptForOfflineDevice *bool                         `json:"attempt_for_offline_device,omitempty"`
	PreferNativeScheduling  *bool                         `json:"prefer_native_scheduling,omitempty"`
	UseBackupAccessCodePool *bool                         `json:"use_backup_access_code_pool,omitempty"`
	AccessCodeId            string                        `json:"access_code_id"`
	DeviceId                *string                       `json:"device_id,omitempty"`
	Type                    *AccessCodesUpdateRequestType `json:"type,omitempty"`
}