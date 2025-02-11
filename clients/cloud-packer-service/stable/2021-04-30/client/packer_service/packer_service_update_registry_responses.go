// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceUpdateRegistryReader is a Reader for the PackerServiceUpdateRegistry structure.
type PackerServiceUpdateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceUpdateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceUpdateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceUpdateRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceUpdateRegistryOK creates a PackerServiceUpdateRegistryOK with default headers values
func NewPackerServiceUpdateRegistryOK() *PackerServiceUpdateRegistryOK {
	return &PackerServiceUpdateRegistryOK{}
}

/*PackerServiceUpdateRegistryOK handles this case with default header values.

A successful response.
*/
type PackerServiceUpdateRegistryOK struct {
	Payload *models.HashicorpCloudPackerUpdateRegistryResponse
}

func (o *PackerServiceUpdateRegistryOK) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/registry][%d] packerServiceUpdateRegistryOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateRegistryOK) GetPayload() *models.HashicorpCloudPackerUpdateRegistryResponse {
	return o.Payload
}

func (o *PackerServiceUpdateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerUpdateRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceUpdateRegistryDefault creates a PackerServiceUpdateRegistryDefault with default headers values
func NewPackerServiceUpdateRegistryDefault(code int) *PackerServiceUpdateRegistryDefault {
	return &PackerServiceUpdateRegistryDefault{
		_statusCode: code,
	}
}

/*PackerServiceUpdateRegistryDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceUpdateRegistryDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service update registry default response
func (o *PackerServiceUpdateRegistryDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceUpdateRegistryDefault) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/registry][%d] PackerService_UpdateRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateRegistryDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceUpdateRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
