// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"addi/models"
)

// GetDSPRecipeOKCode is the HTTP code returned for type GetDSPRecipeOK
const GetDSPRecipeOKCode int = 200

/*GetDSPRecipeOK OK

swagger:response getDSPRecipeOK
*/
type GetDSPRecipeOK struct {

	/*
	  In: Body
	*/
	Payload []*models.DSPRecipeResponse `json:"body,omitempty"`
}

// NewGetDSPRecipeOK creates GetDSPRecipeOK with default headers values
func NewGetDSPRecipeOK() *GetDSPRecipeOK {

	return &GetDSPRecipeOK{}
}

// WithPayload adds the payload to the get d s p recipe o k response
func (o *GetDSPRecipeOK) WithPayload(payload []*models.DSPRecipeResponse) *GetDSPRecipeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get d s p recipe o k response
func (o *GetDSPRecipeOK) SetPayload(payload []*models.DSPRecipeResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDSPRecipeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.DSPRecipeResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetDSPRecipeMethodNotAllowedCode is the HTTP code returned for type GetDSPRecipeMethodNotAllowed
const GetDSPRecipeMethodNotAllowedCode int = 405

/*GetDSPRecipeMethodNotAllowed Invalid input

swagger:response getDSPRecipeMethodNotAllowed
*/
type GetDSPRecipeMethodNotAllowed struct {
}

// NewGetDSPRecipeMethodNotAllowed creates GetDSPRecipeMethodNotAllowed with default headers values
func NewGetDSPRecipeMethodNotAllowed() *GetDSPRecipeMethodNotAllowed {

	return &GetDSPRecipeMethodNotAllowed{}
}

// WriteResponse to the client
func (o *GetDSPRecipeMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
