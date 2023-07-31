// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	strconv "strconv"
)

type ConnectWebviewStatus uint8

const (
	ConnectWebviewStatusPending ConnectWebviewStatus = iota + 1
	ConnectWebviewStatusFailed
	ConnectWebviewStatusAuthorized
)

func (c ConnectWebviewStatus) String() string {
	switch c {
	default:
		return strconv.Itoa(int(c))
	case ConnectWebviewStatusPending:
		return "pending"
	case ConnectWebviewStatusFailed:
		return "failed"
	case ConnectWebviewStatusAuthorized:
		return "authorized"
	}
}

func (c ConnectWebviewStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", c.String())), nil
}

func (c *ConnectWebviewStatus) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch raw {
	case "pending":
		value := ConnectWebviewStatusPending
		*c = value
	case "failed":
		value := ConnectWebviewStatusFailed
		*c = value
	case "authorized":
		value := ConnectWebviewStatusAuthorized
		*c = value
	}
	return nil
}
