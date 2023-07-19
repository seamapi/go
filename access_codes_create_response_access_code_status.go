// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type AccessCodesCreateResponseAccessCodeStatus uint8

const (
	AccessCodesCreateResponseAccessCodeStatusSetting AccessCodesCreateResponseAccessCodeStatus = iota + 1
	AccessCodesCreateResponseAccessCodeStatusSet
	AccessCodesCreateResponseAccessCodeStatusUnset
	AccessCodesCreateResponseAccessCodeStatusRemoving
	AccessCodesCreateResponseAccessCodeStatusUnknown
)

func (a AccessCodesCreateResponseAccessCodeStatus) String() string {
	switch a {
	default:
		return strconv.Itoa(int(a))
	case AccessCodesCreateResponseAccessCodeStatusSetting:
		return "setting"
	case AccessCodesCreateResponseAccessCodeStatusSet:
		return "set"
	case AccessCodesCreateResponseAccessCodeStatusUnset:
		return "unset"
	case AccessCodesCreateResponseAccessCodeStatusRemoving:
		return "removing"
	case AccessCodesCreateResponseAccessCodeStatusUnknown:
		return "unknown"
	}
}

func (a AccessCodesCreateResponseAccessCodeStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

func (a *AccessCodesCreateResponseAccessCodeStatus) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "setting":
		value := AccessCodesCreateResponseAccessCodeStatusSetting
		*a = value
	case "set":
		value := AccessCodesCreateResponseAccessCodeStatusSet
		*a = value
	case "unset":
		value := AccessCodesCreateResponseAccessCodeStatusUnset
		*a = value
	case "removing":
		value := AccessCodesCreateResponseAccessCodeStatusRemoving
		*a = value
	case "unknown":
		value := AccessCodesCreateResponseAccessCodeStatusUnknown
		*a = value
	}
	return nil
}
