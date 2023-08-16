// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared uint8

const (
	AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedThrow AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared = iota + 1
	AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedCreateRandomCode
)

func (a AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared) String() string {
	switch a {
	default:
		return strconv.Itoa(int(a))
	case AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedThrow:
		return "throw"
	case AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedCreateRandomCode:
		return "create_random_code"
	}
}

func (a AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", a.String())), nil
}

func (a *AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeShared) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "throw":
		value := AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedThrow
		*a = value
	case "create_random_code":
		value := AccessCodesCreateMultipleRequestBehaviorWhenCodeCannotBeSharedCreateRandomCode
		*a = value
	}
	return nil
}