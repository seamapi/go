// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type ActionAttempt struct {
	Status  string
	Success *ActionAttemptSuccess
	Pending *ActionAttemptPending
	Error   *ActionAttemptError
}

func NewActionAttemptFromSuccess(value *ActionAttemptSuccess) *ActionAttempt {
	return &ActionAttempt{Status: "success", Success: value}
}

func NewActionAttemptFromPending(value *ActionAttemptPending) *ActionAttempt {
	return &ActionAttempt{Status: "pending", Pending: value}
}

func NewActionAttemptFromError(value *ActionAttemptError) *ActionAttempt {
	return &ActionAttempt{Status: "error", Error: value}
}

func (a *ActionAttempt) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Status = unmarshaler.Status
	switch unmarshaler.Status {
	case "success":
		value := new(ActionAttemptSuccess)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Success = value
	case "pending":
		value := new(ActionAttemptPending)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Pending = value
	case "error":
		value := new(ActionAttemptError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Error = value
	}
	return nil
}

func (a ActionAttempt) MarshalJSON() ([]byte, error) {
	switch a.Status {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", a.Status, a)
	case "success":
		var marshaler = struct {
			Status string `json:"status"`
			*ActionAttemptSuccess
		}{
			Status:               a.Status,
			ActionAttemptSuccess: a.Success,
		}
		return json.Marshal(marshaler)
	case "pending":
		var marshaler = struct {
			Status string `json:"status"`
			*ActionAttemptPending
		}{
			Status:               a.Status,
			ActionAttemptPending: a.Pending,
		}
		return json.Marshal(marshaler)
	case "error":
		var marshaler = struct {
			Status string `json:"status"`
			*ActionAttemptError
		}{
			Status:             a.Status,
			ActionAttemptError: a.Error,
		}
		return json.Marshal(marshaler)
	}
}

type ActionAttemptVisitor interface {
	VisitSuccess(*ActionAttemptSuccess) error
	VisitPending(*ActionAttemptPending) error
	VisitError(*ActionAttemptError) error
}

func (a *ActionAttempt) Accept(v ActionAttemptVisitor) error {
	switch a.Status {
	default:
		return fmt.Errorf("invalid type %s in %T", a.Status, a)
	case "success":
		return v.VisitSuccess(a.Success)
	case "pending":
		return v.VisitPending(a.Pending)
	case "error":
		return v.VisitError(a.Error)
	}
}
