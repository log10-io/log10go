// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// ChatCompletionRequestToolMessageRole - The role of the messages author, in this case `tool`.
type ChatCompletionRequestToolMessageRole string

const (
	ChatCompletionRequestToolMessageRoleTool ChatCompletionRequestToolMessageRole = "tool"
)

func (e ChatCompletionRequestToolMessageRole) ToPointer() *ChatCompletionRequestToolMessageRole {
	return &e
}
func (e *ChatCompletionRequestToolMessageRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tool":
		*e = ChatCompletionRequestToolMessageRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionRequestToolMessageRole: %v", v)
	}
}

type ChatCompletionRequestToolMessage struct {
	// The role of the messages author, in this case `tool`.
	Role ChatCompletionRequestToolMessageRole `json:"role"`
	// The contents of the tool message.
	Content string `json:"content"`
	// Tool call that this message is responding to.
	ToolCallID string `json:"tool_call_id"`
}

func (o *ChatCompletionRequestToolMessage) GetRole() ChatCompletionRequestToolMessageRole {
	if o == nil {
		return ChatCompletionRequestToolMessageRole("")
	}
	return o.Role
}

func (o *ChatCompletionRequestToolMessage) GetContent() string {
	if o == nil {
		return ""
	}
	return o.Content
}

func (o *ChatCompletionRequestToolMessage) GetToolCallID() string {
	if o == nil {
		return ""
	}
	return o.ToolCallID
}