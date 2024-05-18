// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"userprofile/models"
)

// DeleteUserIDNoContentCode is the HTTP code returned for type DeleteUserIDNoContent
const DeleteUserIDNoContentCode int = 204

/*
DeleteUserIDNoContent User deleted

swagger:response deleteUserIdNoContent
*/
type DeleteUserIDNoContent struct {
}

// NewDeleteUserIDNoContent creates DeleteUserIDNoContent with default headers values
func NewDeleteUserIDNoContent() *DeleteUserIDNoContent {

	return &DeleteUserIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteUserIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteUserIDUnauthorizedCode is the HTTP code returned for type DeleteUserIDUnauthorized
const DeleteUserIDUnauthorizedCode int = 401

/*
DeleteUserIDUnauthorized Unauthorized

swagger:response deleteUserIdUnauthorized
*/
type DeleteUserIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUserIDUnauthorized creates DeleteUserIDUnauthorized with default headers values
func NewDeleteUserIDUnauthorized() *DeleteUserIDUnauthorized {

	return &DeleteUserIDUnauthorized{}
}

// WithPayload adds the payload to the delete user Id unauthorized response
func (o *DeleteUserIDUnauthorized) WithPayload(payload *models.ErrorResponse) *DeleteUserIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user Id unauthorized response
func (o *DeleteUserIDUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserIDForbiddenCode is the HTTP code returned for type DeleteUserIDForbidden
const DeleteUserIDForbiddenCode int = 403

/*
DeleteUserIDForbidden Forbidden

swagger:response deleteUserIdForbidden
*/
type DeleteUserIDForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUserIDForbidden creates DeleteUserIDForbidden with default headers values
func NewDeleteUserIDForbidden() *DeleteUserIDForbidden {

	return &DeleteUserIDForbidden{}
}

// WithPayload adds the payload to the delete user Id forbidden response
func (o *DeleteUserIDForbidden) WithPayload(payload *models.ErrorResponse) *DeleteUserIDForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user Id forbidden response
func (o *DeleteUserIDForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteUserIDNotFoundCode is the HTTP code returned for type DeleteUserIDNotFound
const DeleteUserIDNotFoundCode int = 404

/*
DeleteUserIDNotFound User not found

swagger:response deleteUserIdNotFound
*/
type DeleteUserIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDeleteUserIDNotFound creates DeleteUserIDNotFound with default headers values
func NewDeleteUserIDNotFound() *DeleteUserIDNotFound {

	return &DeleteUserIDNotFound{}
}

// WithPayload adds the payload to the delete user Id not found response
func (o *DeleteUserIDNotFound) WithPayload(payload *models.ErrorResponse) *DeleteUserIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete user Id not found response
func (o *DeleteUserIDNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteUserIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
