// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Kind - The kind of completion i.e. chat messages or prompt
type Kind string

const (
	KindChat   Kind = "chat"
	KindPrompt Kind = "prompt"
)

func (e Kind) ToPointer() *Kind {
	return &e
}
func (e *Kind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "chat":
		fallthrough
	case "prompt":
		*e = Kind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Kind: %v", v)
	}
}

// Status - The status of this completion.
type Status string

const (
	StatusStarted  Status = "started"
	StatusFinished Status = "finished"
	StatusFailed   Status = "failed"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "started":
		fallthrough
	case "finished":
		fallthrough
	case "failed":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type Stacktrace struct {
	// The file associated with this stacktrace.
	File string `json:"file"`
	// The line associated with this stacktrace.
	Line string `json:"line"`
	// The line number associated with this stacktrace.
	Lineno float64 `json:"lineno"`
	// The function or module associated with this stacktrace.
	Name string `json:"name"`
}

func (o *Stacktrace) GetFile() string {
	if o == nil {
		return ""
	}
	return o.File
}

func (o *Stacktrace) GetLine() string {
	if o == nil {
		return ""
	}
	return o.Line
}

func (o *Stacktrace) GetLineno() float64 {
	if o == nil {
		return 0.0
	}
	return o.Lineno
}

func (o *Stacktrace) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type Completion struct {
	// The unique identifier for this task.
	ID *string `json:"id,omitempty"`
	// The unique identifier for the organization.
	OrganizationID string `json:"organization_id"`
	// The kind of completion i.e. chat messages or prompt
	Kind *Kind `json:"kind,omitempty"`
	// The status of this completion.
	Status *Status `json:"status,omitempty"`
	// The tags for this completion.
	Tags    []string                     `json:"tags,omitempty"`
	Request *CreateChatCompletionRequest `json:"request,omitempty"`
	// Represents a chat completion response returned by model, based on the provided input.
	Response *CreateChatCompletionResponse `json:"response,omitempty"`
	// The stacktrace for this completion.
	Stacktrace []Stacktrace `json:"stacktrace,omitempty"`
	// The session id for this completion.
	SessionID *string `json:"session_id,omitempty"`
	// The duration of this completion in seconds.
	Duration *float64 `json:"duration,omitempty"`
	// The failure kind of this completion.
	FailureKind *string `json:"failure_kind,omitempty"`
	// The failure reason of this completion.
	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o *Completion) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Completion) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *Completion) GetKind() *Kind {
	if o == nil {
		return nil
	}
	return o.Kind
}

func (o *Completion) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *Completion) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Completion) GetRequest() *CreateChatCompletionRequest {
	if o == nil {
		return nil
	}
	return o.Request
}

func (o *Completion) GetResponse() *CreateChatCompletionResponse {
	if o == nil {
		return nil
	}
	return o.Response
}

func (o *Completion) GetStacktrace() []Stacktrace {
	if o == nil {
		return nil
	}
	return o.Stacktrace
}

func (o *Completion) GetSessionID() *string {
	if o == nil {
		return nil
	}
	return o.SessionID
}

func (o *Completion) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *Completion) GetFailureKind() *string {
	if o == nil {
		return nil
	}
	return o.FailureKind
}

func (o *Completion) GetFailureReason() *string {
	if o == nil {
		return nil
	}
	return o.FailureReason
}
