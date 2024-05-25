// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Type - The type of the content part.
type Type string

const (
	TypeText Type = "text"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type ChatCompletionRequestMessageContentPartText struct {
	// The type of the content part.
	Type Type `json:"type"`
	// The text content.
	Text string `json:"text"`
}

func (o *ChatCompletionRequestMessageContentPartText) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

func (o *ChatCompletionRequestMessageContentPartText) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}