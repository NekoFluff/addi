// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"addi/models"
)

// GetInArtOKCode is the HTTP code returned for type GetInArtOK
const GetInArtOKCode int = 200

/*GetInArtOK OK

swagger:response getInArtOK
*/
type GetInArtOK struct {

	/*
	  In: Body
	*/
	Payload []*models.InArt `json:"body,omitempty"`
}

// NewGetInArtOK creates GetInArtOK with default headers values
func NewGetInArtOK() *GetInArtOK {

	return &GetInArtOK{}
}

// WithPayload adds the payload to the get in art o k response
func (o *GetInArtOK) WithPayload(payload []*models.InArt) *GetInArtOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get in art o k response
func (o *GetInArtOK) SetPayload(payload []*models.InArt) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInArtOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.InArt, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
