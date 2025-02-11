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

// UpdateVersionReader is a Reader for the UpdateVersion structure.
type UpdateVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateVersionOK creates a UpdateVersionOK with default headers values
func NewUpdateVersionOK() *UpdateVersionOK {
	return &UpdateVersionOK{}
}

/*UpdateVersionOK handles this case with default header values.

A successful response.
*/
type UpdateVersionOK struct {
	Payload *models.HashicorpCloudVault20201125UpdateVersionResponse
}

func (o *UpdateVersionOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/{version}][%d] updateVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateVersionOK) GetPayload() *models.HashicorpCloudVault20201125UpdateVersionResponse {
	return o.Payload
}

func (o *UpdateVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125UpdateVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVersionDefault creates a UpdateVersionDefault with default headers values
func NewUpdateVersionDefault(code int) *UpdateVersionDefault {
	return &UpdateVersionDefault{
		_statusCode: code,
	}
}

/*UpdateVersionDefault handles this case with default header values.

An unexpected error response.
*/
type UpdateVersionDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the update version default response
func (o *UpdateVersionDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVersionDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/{version}][%d] UpdateVersion default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateVersionDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *UpdateVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
