// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type ActionAttemptsGetResponseActionAttempt struct {
	Status  string
	Success *ActionAttemptsGetResponseActionAttemptSuccess
	Pending *ActionAttemptsGetResponseActionAttemptPending
	Error   *ActionAttemptsGetResponseActionAttemptError
}

func NewActionAttemptsGetResponseActionAttemptFromSuccess(value *ActionAttemptsGetResponseActionAttemptSuccess) *ActionAttemptsGetResponseActionAttempt {
	return &ActionAttemptsGetResponseActionAttempt{Status: "success", Success: value}
}

func NewActionAttemptsGetResponseActionAttemptFromPending(value *ActionAttemptsGetResponseActionAttemptPending) *ActionAttemptsGetResponseActionAttempt {
	return &ActionAttemptsGetResponseActionAttempt{Status: "pending", Pending: value}
}

func NewActionAttemptsGetResponseActionAttemptFromError(value *ActionAttemptsGetResponseActionAttemptError) *ActionAttemptsGetResponseActionAttempt {
	return &ActionAttemptsGetResponseActionAttempt{Status: "error", Error: value}
}

func (a *ActionAttemptsGetResponseActionAttempt) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Status = unmarshaler.Status
	switch unmarshaler.Status {
	case "success":
		value := new(ActionAttemptsGetResponseActionAttemptSuccess)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Success = value
	case "pending":
		value := new(ActionAttemptsGetResponseActionAttemptPending)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Pending = value
	case "error":
		value := new(ActionAttemptsGetResponseActionAttemptError)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Error = value
	}
	return nil
}

func (a ActionAttemptsGetResponseActionAttempt) MarshalJSON() ([]byte, error) {
	switch a.Status {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", a.Status, a)
	case "success":
		var marshaler = struct {
			Status string `json:"status"`
			*ActionAttemptsGetResponseActionAttemptSuccess
		}{
			Status: a.Status,
			ActionAttemptsGetResponseActionAttemptSuccess: a.Success,
		}
		return json.Marshal(marshaler)
	case "pending":
		var marshaler = struct {
			Status string `json:"status"`
			*ActionAttemptsGetResponseActionAttemptPending
		}{
			Status: a.Status,
			ActionAttemptsGetResponseActionAttemptPending: a.Pending,
		}
		return json.Marshal(marshaler)
	case "error":
		var marshaler = struct {
			Status string `json:"status"`
			*ActionAttemptsGetResponseActionAttemptError
		}{
			Status: a.Status,
			ActionAttemptsGetResponseActionAttemptError: a.Error,
		}
		return json.Marshal(marshaler)
	}
}

type ActionAttemptsGetResponseActionAttemptVisitor interface {
	VisitSuccess(*ActionAttemptsGetResponseActionAttemptSuccess) error
	VisitPending(*ActionAttemptsGetResponseActionAttemptPending) error
	VisitError(*ActionAttemptsGetResponseActionAttemptError) error
}

func (a *ActionAttemptsGetResponseActionAttempt) Accept(v ActionAttemptsGetResponseActionAttemptVisitor) error {
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
