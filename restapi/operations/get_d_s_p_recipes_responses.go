// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"addi/models"
)

// GetDSPRecipesOKCode is the HTTP code returned for type GetDSPRecipesOK
const GetDSPRecipesOKCode int = 200

/*GetDSPRecipesOK OK

swagger:response getDSPRecipesOK
*/
type GetDSPRecipesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.DSPRecipe `json:"body,omitempty"`
}

// NewGetDSPRecipesOK creates GetDSPRecipesOK with default headers values
func NewGetDSPRecipesOK() *GetDSPRecipesOK {

	return &GetDSPRecipesOK{}
}

// WithPayload adds the payload to the get d s p recipes o k response
func (o *GetDSPRecipesOK) WithPayload(payload []*models.DSPRecipe) *GetDSPRecipesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get d s p recipes o k response
func (o *GetDSPRecipesOK) SetPayload(payload []*models.DSPRecipe) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDSPRecipesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.DSPRecipe, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
