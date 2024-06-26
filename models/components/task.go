// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

// JSONSchema - The schema of the task. Must be valid JSON Schema.
type JSONSchema struct {
}

// CompletionTagsSelector - The completion tag matching with this task i.e. surfaced as needing feedback.
type CompletionTagsSelector struct {
}

type Task struct {
	// The unique identifier for this task.
	ID *string `json:"id,omitempty"`
	// The epoch this schema was created.
	CreatedAtMs *float64 `json:"created_at_ms,omitempty"`
	// The schema of the task. Must be valid JSON Schema.
	JSONSchema JSONSchema `json:"json_schema"`
	// The name of the task.
	Name string `json:"name"`
	// The instructions for this task.
	Instruction string `json:"instruction"`
	// The completion tag matching with this task i.e. surfaced as needing feedback.
	CompletionTagsSelector CompletionTagsSelector `json:"completion_tags_selector"`
}

func (o *Task) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Task) GetCreatedAtMs() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAtMs
}

func (o *Task) GetJSONSchema() JSONSchema {
	if o == nil {
		return JSONSchema{}
	}
	return o.JSONSchema
}

func (o *Task) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Task) GetInstruction() string {
	if o == nil {
		return ""
	}
	return o.Instruction
}

func (o *Task) GetCompletionTagsSelector() CompletionTagsSelector {
	if o == nil {
		return CompletionTagsSelector{}
	}
	return o.CompletionTagsSelector
}
