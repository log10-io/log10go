// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type Session struct {
	// The unique identifier for this session.
	ID *string `json:"id,omitempty"`
}

func (o *Session) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
