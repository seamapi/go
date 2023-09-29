// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type ConnectWebviewsCreateRequestCustomMetadataValue struct {
	typeName       string
	String         string
	Double         float64
	StringOptional *string
	Boolean        bool
}

func NewConnectWebviewsCreateRequestCustomMetadataValueFromString(value string) *ConnectWebviewsCreateRequestCustomMetadataValue {
	return &ConnectWebviewsCreateRequestCustomMetadataValue{typeName: "string", String: value}
}

func NewConnectWebviewsCreateRequestCustomMetadataValueFromDouble(value float64) *ConnectWebviewsCreateRequestCustomMetadataValue {
	return &ConnectWebviewsCreateRequestCustomMetadataValue{typeName: "double", Double: value}
}

func NewConnectWebviewsCreateRequestCustomMetadataValueFromStringOptional(value *string) *ConnectWebviewsCreateRequestCustomMetadataValue {
	return &ConnectWebviewsCreateRequestCustomMetadataValue{typeName: "stringOptional", StringOptional: value}
}

func NewConnectWebviewsCreateRequestCustomMetadataValueFromBoolean(value bool) *ConnectWebviewsCreateRequestCustomMetadataValue {
	return &ConnectWebviewsCreateRequestCustomMetadataValue{typeName: "boolean", Boolean: value}
}

func (c *ConnectWebviewsCreateRequestCustomMetadataValue) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		c.typeName = "string"
		c.String = valueString
		return nil
	}
	var valueDouble float64
	if err := json.Unmarshal(data, &valueDouble); err == nil {
		c.typeName = "double"
		c.Double = valueDouble
		return nil
	}
	var valueStringOptional *string
	if err := json.Unmarshal(data, &valueStringOptional); err == nil {
		c.typeName = "stringOptional"
		c.StringOptional = valueStringOptional
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		c.typeName = "boolean"
		c.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ConnectWebviewsCreateRequestCustomMetadataValue) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return json.Marshal(c.String)
	case "double":
		return json.Marshal(c.Double)
	case "stringOptional":
		return json.Marshal(c.StringOptional)
	case "boolean":
		return json.Marshal(c.Boolean)
	}
}

type ConnectWebviewsCreateRequestCustomMetadataValueVisitor interface {
	VisitString(string) error
	VisitDouble(float64) error
	VisitStringOptional(*string) error
	VisitBoolean(bool) error
}

func (c *ConnectWebviewsCreateRequestCustomMetadataValue) Accept(v ConnectWebviewsCreateRequestCustomMetadataValueVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "string":
		return v.VisitString(c.String)
	case "double":
		return v.VisitDouble(c.Double)
	case "stringOptional":
		return v.VisitStringOptional(c.StringOptional)
	case "boolean":
		return v.VisitBoolean(c.Boolean)
	}
}
