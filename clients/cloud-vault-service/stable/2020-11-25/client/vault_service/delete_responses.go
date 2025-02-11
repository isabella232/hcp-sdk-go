// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// DeleteReader is a Reader for the Delete structure.
type DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOK creates a DeleteOK with default headers values
func NewDeleteOK() *DeleteOK {
	return &DeleteOK{}
}

/*DeleteOK handles this case with default header values.

A successful response.
*/
type DeleteOK struct {
	Payload *models.HashicorpCloudVault20201125DeleteResponse
}

func (o *DeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}][%d] deleteOK  %+v", 200, o.Payload)
}

func (o *DeleteOK) GetPayload() *models.HashicorpCloudVault20201125DeleteResponse {
	return o.Payload
}

func (o *DeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125DeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDefault creates a DeleteDefault with default headers values
func NewDeleteDefault(code int) *DeleteDefault {
	return &DeleteDefault{
		_statusCode: code,
	}
}

/*DeleteDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the delete default response
func (o *DeleteDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}][%d] Delete default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *DeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
