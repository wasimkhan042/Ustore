// Code generated by go-swagger; DO NOT EDIT.

package item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ustore/gen/models"
)

// ItemsOKCode is the HTTP code returned for type ItemsOK
const ItemsOKCode int = 200

/*ItemsOK Success response when items are shown

swagger:response itemsOK
*/
type ItemsOK struct {

	/*
	  In: Body
	*/
	Payload models.Products `json:"body,omitempty"`
}

// NewItemsOK creates ItemsOK with default headers values
func NewItemsOK() *ItemsOK {

	return &ItemsOK{}
}

// WithPayload adds the payload to the items o k response
func (o *ItemsOK) WithPayload(payload models.Products) *ItemsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the items o k response
func (o *ItemsOK) SetPayload(payload models.Products) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ItemsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Products{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ItemsBadRequestCode is the HTTP code returned for type ItemsBadRequest
const ItemsBadRequestCode int = 400

/*ItemsBadRequest Bad Request

swagger:response itemsBadRequest
*/
type ItemsBadRequest struct {
}

// NewItemsBadRequest creates ItemsBadRequest with default headers values
func NewItemsBadRequest() *ItemsBadRequest {

	return &ItemsBadRequest{}
}

// WriteResponse to the client
func (o *ItemsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// ItemsNotFoundCode is the HTTP code returned for type ItemsNotFound
const ItemsNotFoundCode int = 404

/*ItemsNotFound items not found

swagger:response itemsNotFound
*/
type ItemsNotFound struct {
}

// NewItemsNotFound creates ItemsNotFound with default headers values
func NewItemsNotFound() *ItemsNotFound {

	return &ItemsNotFound{}
}

// WriteResponse to the client
func (o *ItemsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ItemsInternalServerErrorCode is the HTTP code returned for type ItemsInternalServerError
const ItemsInternalServerErrorCode int = 500

/*ItemsInternalServerError Server error

swagger:response itemsInternalServerError
*/
type ItemsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewItemsInternalServerError creates ItemsInternalServerError with default headers values
func NewItemsInternalServerError() *ItemsInternalServerError {

	return &ItemsInternalServerError{}
}

// WithPayload adds the payload to the items internal server error response
func (o *ItemsInternalServerError) WithPayload(payload string) *ItemsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the items internal server error response
func (o *ItemsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ItemsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
