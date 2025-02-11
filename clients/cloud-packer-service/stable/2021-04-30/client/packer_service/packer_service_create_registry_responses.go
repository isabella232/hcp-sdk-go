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

// PackerServiceCreateRegistryReader is a Reader for the PackerServiceCreateRegistry structure.
type PackerServiceCreateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceCreateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceCreateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceCreateRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceCreateRegistryOK creates a PackerServiceCreateRegistryOK with default headers values
func NewPackerServiceCreateRegistryOK() *PackerServiceCreateRegistryOK {
	return &PackerServiceCreateRegistryOK{}
}

/*PackerServiceCreateRegistryOK handles this case with default header values.

A successful response.
*/
type PackerServiceCreateRegistryOK struct {
	Payload *models.HashicorpCloudPackerCreateRegistryResponse
}

func (o *PackerServiceCreateRegistryOK) Error() string {
	return fmt.Sprintf("[PUT /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/registry][%d] packerServiceCreateRegistryOK  %+v", 200, o.Payload)
}

func (o *PackerServiceCreateRegistryOK) GetPayload() *models.HashicorpCloudPackerCreateRegistryResponse {
	return o.Payload
}

func (o *PackerServiceCreateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerCreateRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceCreateRegistryDefault creates a PackerServiceCreateRegistryDefault with default headers values
func NewPackerServiceCreateRegistryDefault(code int) *PackerServiceCreateRegistryDefault {
	return &PackerServiceCreateRegistryDefault{
		_statusCode: code,
	}
}

/*PackerServiceCreateRegistryDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceCreateRegistryDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service create registry default response
func (o *PackerServiceCreateRegistryDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceCreateRegistryDefault) Error() string {
	return fmt.Sprintf("[PUT /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/registry][%d] PackerService_CreateRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceCreateRegistryDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceCreateRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
