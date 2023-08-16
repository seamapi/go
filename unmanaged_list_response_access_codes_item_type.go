// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type UnmanagedListResponseAccessCodesItemType uint8

const (
	UnmanagedListResponseAccessCodesItemTypeTimeBound UnmanagedListResponseAccessCodesItemType = iota + 1
	UnmanagedListResponseAccessCodesItemTypeOngoing
)

func (u UnmanagedListResponseAccessCodesItemType) String() string {
	switch u {
	default:
		return strconv.Itoa(int(u))
	case UnmanagedListResponseAccessCodesItemTypeTimeBound:
		return "time_bound"
	case UnmanagedListResponseAccessCodesItemTypeOngoing:
		return "ongoing"
	}
}

func (u UnmanagedListResponseAccessCodesItemType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", u.String())), nil
}

func (u *UnmanagedListResponseAccessCodesItemType) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "time_bound":
		value := UnmanagedListResponseAccessCodesItemTypeTimeBound
		*u = value
	case "ongoing":
		value := UnmanagedListResponseAccessCodesItemTypeOngoing
		*u = value
	}
	return nil
}