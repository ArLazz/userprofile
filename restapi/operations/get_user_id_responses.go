// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"userprofile/models"
)

// GetUserIDOKCode is the HTTP code returned for type GetUserIDOK
const GetUserIDOKCode int = 200

/*
GetUserIDOK Successful operation

swagger:response getUserIdOK
*/
type GetUserIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserIDOK creates GetUserIDOK with default headers values
func NewGetUserIDOK() *GetUserIDOK {

	return &GetUserIDOK{}
}

// WithPayload adds the payload to the get user Id o k response
func (o *GetUserIDOK) WithPayload(payload *models.User) *GetUserIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user Id o k response
func (o *GetUserIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserIDNotFoundCode is the HTTP code returned for type GetUserIDNotFound
const GetUserIDNotFoundCode int = 404

/*
GetUserIDNotFound User not found

swagger:response getUserIdNotFound
*/
type GetUserIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewGetUserIDNotFound creates GetUserIDNotFound with default headers values
func NewGetUserIDNotFound() *GetUserIDNotFound {

	return &GetUserIDNotFound{}
}

// WithPayload adds the payload to the get user Id not found response
func (o *GetUserIDNotFound) WithPayload(payload *models.ErrorResponse) *GetUserIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user Id not found response
func (o *GetUserIDNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
