// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/log10-io/log10go/models/components"
)

type ListUngradedGlobals struct {
	XLog10Organization *string `header:"style=simple,explode=false,name=X-Log10-Organization"`
}

func (o *ListUngradedGlobals) GetXLog10Organization() *string {
	if o == nil {
		return nil
	}
	return o.XLog10Organization
}

type ListUngradedRequest struct {
	XLog10Organization *string `header:"style=simple,explode=false,name=X-Log10-Organization"`
}

func (o *ListUngradedRequest) GetXLog10Organization() *string {
	if o == nil {
		return nil
	}
	return o.XLog10Organization
}

// ListUngradedResponseBody - OK
type ListUngradedResponseBody struct {
	Completions []components.Completion `json:"completions,omitempty"`
}

func (o *ListUngradedResponseBody) GetCompletions() []components.Completion {
	if o == nil {
		return nil
	}
	return o.Completions
}

type ListUngradedResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Object *ListUngradedResponseBody
}

func (o *ListUngradedResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListUngradedResponse) GetObject() *ListUngradedResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
