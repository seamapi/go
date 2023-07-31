// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type LocksListRequestDeviceTypesItem struct {
	typeName                                                       string
	LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
	LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
	LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
}

func NewLocksListRequestDeviceTypesItemFromLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(value LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem) *LocksListRequestDeviceTypesItem {
	return &LocksListRequestDeviceTypesItem{typeName: "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem", LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem: value}
}

func NewLocksListRequestDeviceTypesItemFromLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(value LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem) *LocksListRequestDeviceTypesItem {
	return &LocksListRequestDeviceTypesItem{typeName: "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem", LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem: value}
}

func NewLocksListRequestDeviceTypesItemFromLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(value LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem) *LocksListRequestDeviceTypesItem {
	return &LocksListRequestDeviceTypesItem{typeName: "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem", LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem: value}
}

func (l *LocksListRequestDeviceTypesItem) UnmarshalJSON(data []byte) error {
	var valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
	if err := json.Unmarshal(data, &valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem); err == nil {
		l.typeName = "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem"
		l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem = valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
		return nil
	}
	var valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
	if err := json.Unmarshal(data, &valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem); err == nil {
		l.typeName = "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem"
		l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem = valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
		return nil
	}
	var valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
	if err := json.Unmarshal(data, &valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem); err == nil {
		l.typeName = "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem"
		l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem = valueLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LocksListRequestDeviceTypesItem) MarshalJSON() ([]byte, error) {
	switch l.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", l.typeName, l)
	case "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem":
		return json.Marshal(l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem)
	case "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem":
		return json.Marshal(l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem)
	case "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem":
		return json.Marshal(l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem)
	}
}

type LocksListRequestDeviceTypesItemVisitor interface {
	VisitLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem) error
	VisitLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem) error
	VisitLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem) error
}

func (l *LocksListRequestDeviceTypesItem) Accept(v LocksListRequestDeviceTypesItemVisitor) error {
	switch l.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", l.typeName, l)
	case "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem":
		return v.VisitLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem)
	case "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem":
		return v.VisitLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem)
	case "locksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem":
		return v.VisitLocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem(l.LocksListRequestDeviceTypesItemLocksListRequestDeviceTypesItem)
	}
}
