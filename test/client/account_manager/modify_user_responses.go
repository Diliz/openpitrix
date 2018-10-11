// Code generated by go-swagger; DO NOT EDIT.

package account_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// ModifyUserReader is a Reader for the ModifyUser structure.
type ModifyUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyUserOK creates a ModifyUserOK with default headers values
func NewModifyUserOK() *ModifyUserOK {
	return &ModifyUserOK{}
}

/*ModifyUserOK handles this case with default header values.

A successful response.
*/
type ModifyUserOK struct {
	Payload *models.OpenpitrixModifyUserResponse
}

func (o *ModifyUserOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/users][%d] modifyUserOK  %+v", 200, o.Payload)
}

func (o *ModifyUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixModifyUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}