// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type AccessCodesUpdateRequestType uint8

const (
	AccessCodesUpdateRequestTypeOngoing AccessCodesUpdateRequestType = iota + 1
	AccessCodesUpdateRequestTypeTimeBound
)

func (a AccessCodesUpdateRequestType) String() string {
	switch a {
	default:
		return strconv.Itoa(int(a))
	case AccessCodesUpdateRequestTypeOngoing:
		return "ongoing"
	case AccessCodesUpdateRequestTypeTimeBound:
		return "time_bound"
	}
}

func (a AccessCodesUpdateRequestType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

func (a *AccessCodesUpdateRequestType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "ongoing":
		value := AccessCodesUpdateRequestTypeOngoing
		*a = value
	case "time_bound":
		value := AccessCodesUpdateRequestTypeTimeBound
		*a = value
	}
	return nil
}
