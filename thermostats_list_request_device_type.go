// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type ThermostatsListRequestDeviceType struct {
	typeName                                                         string
	ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
	ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
	ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
}

func NewThermostatsListRequestDeviceTypeFromThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(value ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType) *ThermostatsListRequestDeviceType {
	return &ThermostatsListRequestDeviceType{typeName: "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType", ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType: value}
}

func NewThermostatsListRequestDeviceTypeFromThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(value ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType) *ThermostatsListRequestDeviceType {
	return &ThermostatsListRequestDeviceType{typeName: "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType", ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType: value}
}

func NewThermostatsListRequestDeviceTypeFromThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(value ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType) *ThermostatsListRequestDeviceType {
	return &ThermostatsListRequestDeviceType{typeName: "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType", ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType: value}
}

func (t *ThermostatsListRequestDeviceType) UnmarshalJSON(data []byte) error {
	var valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
	if err := json.Unmarshal(data, &valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType); err == nil {
		t.typeName = "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType"
		t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType = valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
		return nil
	}
	var valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
	if err := json.Unmarshal(data, &valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType); err == nil {
		t.typeName = "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType"
		t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType = valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
		return nil
	}
	var valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
	if err := json.Unmarshal(data, &valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType); err == nil {
		t.typeName = "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType"
		t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType = valueThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, t)
}

func (t ThermostatsListRequestDeviceType) MarshalJSON() ([]byte, error) {
	switch t.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", t.typeName, t)
	case "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType":
		return json.Marshal(t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType)
	case "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType":
		return json.Marshal(t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType)
	case "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType":
		return json.Marshal(t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType)
	}
}

type ThermostatsListRequestDeviceTypeVisitor interface {
	VisitThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType) error
	VisitThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType) error
	VisitThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType) error
}

func (t *ThermostatsListRequestDeviceType) Accept(v ThermostatsListRequestDeviceTypeVisitor) error {
	switch t.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", t.typeName, t)
	case "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType":
		return v.VisitThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType)
	case "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType":
		return v.VisitThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType)
	case "thermostatsListRequestDeviceTypeThermostatsListRequestDeviceType":
		return v.VisitThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType(t.ThermostatsListRequestDeviceTypeThermostatsListRequestDeviceType)
	}
}
