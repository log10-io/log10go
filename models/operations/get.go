// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/log10-io/log10go/models/components"
)

type GetGlobals struct {
	XLog10Organization *string `header:"style=simple,explode=false,name=X-Log10-Organization"`
}

func (o *GetGlobals) GetXLog10Organization() *string {
	if o == nil {
		return nil
	}
	return o.XLog10Organization
}

type GetRequest struct {
	// The feedback id to fetch.
	FeedbackID         string  `pathParam:"style=simple,explode=false,name=feedbackId"`
	XLog10Organization *string `header:"style=simple,explode=false,name=X-Log10-Organization"`
}

func (o *GetRequest) GetFeedbackID() string {
	if o == nil {
		return ""
	}
	return o.FeedbackID
}

func (o *GetRequest) GetXLog10Organization() *string {
	if o == nil {
		return nil
	}
	return o.XLog10Organization
}

type GetResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Feedback *components.Feedback
}

func (o *GetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetResponse) GetFeedback() *components.Feedback {
	if o == nil {
		return nil
	}
	return o.Feedback
}
