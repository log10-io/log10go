// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/log10-io/log10go/internal/utils"
)

type ContentType string

const (
	ContentTypeStr                                            ContentType = "str"
	ContentTypeArrayOfChatCompletionRequestMessageContentPart ContentType = "arrayOfChatCompletionRequestMessageContentPart"
)

// Content - The contents of the user message.
type Content struct {
	Str                                            *string
	ArrayOfChatCompletionRequestMessageContentPart []ChatCompletionRequestMessageContentPart

	Type ContentType
}

func CreateContentStr(str string) Content {
	typ := ContentTypeStr

	return Content{
		Str:  &str,
		Type: typ,
	}
}

func CreateContentArrayOfChatCompletionRequestMessageContentPart(arrayOfChatCompletionRequestMessageContentPart []ChatCompletionRequestMessageContentPart) Content {
	typ := ContentTypeArrayOfChatCompletionRequestMessageContentPart

	return Content{
		ArrayOfChatCompletionRequestMessageContentPart: arrayOfChatCompletionRequestMessageContentPart,
		Type: typ,
	}
}

func (u *Content) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ContentTypeStr
		return nil
	}

	var arrayOfChatCompletionRequestMessageContentPart []ChatCompletionRequestMessageContentPart = []ChatCompletionRequestMessageContentPart{}
	if err := utils.UnmarshalJSON(data, &arrayOfChatCompletionRequestMessageContentPart, "", true, true); err == nil {
		u.ArrayOfChatCompletionRequestMessageContentPart = arrayOfChatCompletionRequestMessageContentPart
		u.Type = ContentTypeArrayOfChatCompletionRequestMessageContentPart
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Content", string(data))
}

func (u Content) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfChatCompletionRequestMessageContentPart != nil {
		return utils.MarshalJSON(u.ArrayOfChatCompletionRequestMessageContentPart, "", true)
	}

	return nil, errors.New("could not marshal union type Content: all fields are null")
}

// ChatCompletionRequestUserMessageRole - The role of the messages author, in this case `user`.
type ChatCompletionRequestUserMessageRole string

const (
	ChatCompletionRequestUserMessageRoleUser ChatCompletionRequestUserMessageRole = "user"
)

func (e ChatCompletionRequestUserMessageRole) ToPointer() *ChatCompletionRequestUserMessageRole {
	return &e
}
func (e *ChatCompletionRequestUserMessageRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "user":
		*e = ChatCompletionRequestUserMessageRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChatCompletionRequestUserMessageRole: %v", v)
	}
}

type ChatCompletionRequestUserMessage struct {
	// The contents of the user message.
	//
	Content Content `json:"content"`
	// The role of the messages author, in this case `user`.
	Role ChatCompletionRequestUserMessageRole `json:"role"`
	// An optional name for the participant. Provides the model information to differentiate between participants of the same role.
	Name *string `json:"name,omitempty"`
}

func (o *ChatCompletionRequestUserMessage) GetContent() Content {
	if o == nil {
		return Content{}
	}
	return o.Content
}

func (o *ChatCompletionRequestUserMessage) GetRole() ChatCompletionRequestUserMessageRole {
	if o == nil {
		return ChatCompletionRequestUserMessageRole("")
	}
	return o.Role
}

func (o *ChatCompletionRequestUserMessage) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}