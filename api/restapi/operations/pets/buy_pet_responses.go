// Code generated by go-swagger; DO NOT EDIT.

package pets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// BuyPetOKCode is the HTTP code returned for type BuyPetOK
const BuyPetOKCode int = 200

/*
BuyPetOK 200

swagger:response buyPetOK
*/
type BuyPetOK struct {
}

// NewBuyPetOK creates BuyPetOK with default headers values
func NewBuyPetOK() *BuyPetOK {

	return &BuyPetOK{}
}

// WriteResponse to the client
func (o *BuyPetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// BuyPetBadRequestCode is the HTTP code returned for type BuyPetBadRequest
const BuyPetBadRequestCode int = 400

/*
BuyPetBadRequest Bad Request

swagger:response buyPetBadRequest
*/
type BuyPetBadRequest struct {
}

// NewBuyPetBadRequest creates BuyPetBadRequest with default headers values
func NewBuyPetBadRequest() *BuyPetBadRequest {

	return &BuyPetBadRequest{}
}

// WriteResponse to the client
func (o *BuyPetBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// BuyPetNotFoundCode is the HTTP code returned for type BuyPetNotFound
const BuyPetNotFoundCode int = 404

/*
BuyPetNotFound Pets Not Found

swagger:response buyPetNotFound
*/
type BuyPetNotFound struct {
}

// NewBuyPetNotFound creates BuyPetNotFound with default headers values
func NewBuyPetNotFound() *BuyPetNotFound {

	return &BuyPetNotFound{}
}

// WriteResponse to the client
func (o *BuyPetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// BuyPetInternalServerErrorCode is the HTTP code returned for type BuyPetInternalServerError
const BuyPetInternalServerErrorCode int = 500

/*
BuyPetInternalServerError Server Error

swagger:response buyPetInternalServerError
*/
type BuyPetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewBuyPetInternalServerError creates BuyPetInternalServerError with default headers values
func NewBuyPetInternalServerError() *BuyPetInternalServerError {

	return &BuyPetInternalServerError{}
}

// WithPayload adds the payload to the buy pet internal server error response
func (o *BuyPetInternalServerError) WithPayload(payload string) *BuyPetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the buy pet internal server error response
func (o *BuyPetInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BuyPetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
