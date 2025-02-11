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

// PackerServiceCreateIterationReader is a Reader for the PackerServiceCreateIteration structure.
type PackerServiceCreateIterationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceCreateIterationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceCreateIterationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceCreateIterationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceCreateIterationOK creates a PackerServiceCreateIterationOK with default headers values
func NewPackerServiceCreateIterationOK() *PackerServiceCreateIterationOK {
	return &PackerServiceCreateIterationOK{}
}

/*PackerServiceCreateIterationOK handles this case with default header values.

A successful response.
*/
type PackerServiceCreateIterationOK struct {
	Payload *models.HashicorpCloudPackerCreateIterationResponse
}

func (o *PackerServiceCreateIterationOK) Error() string {
	return fmt.Sprintf("[POST /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] packerServiceCreateIterationOK  %+v", 200, o.Payload)
}

func (o *PackerServiceCreateIterationOK) GetPayload() *models.HashicorpCloudPackerCreateIterationResponse {
	return o.Payload
}

func (o *PackerServiceCreateIterationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerCreateIterationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceCreateIterationDefault creates a PackerServiceCreateIterationDefault with default headers values
func NewPackerServiceCreateIterationDefault(code int) *PackerServiceCreateIterationDefault {
	return &PackerServiceCreateIterationDefault{
		_statusCode: code,
	}
}

/*PackerServiceCreateIterationDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceCreateIterationDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service create iteration default response
func (o *PackerServiceCreateIterationDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceCreateIterationDefault) Error() string {
	return fmt.Sprintf("[POST /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/iterations][%d] PackerService_CreateIteration default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceCreateIterationDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceCreateIterationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
