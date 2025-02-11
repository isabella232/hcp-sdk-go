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

// PackerServiceGetRegistryTFCRunTaskAPIReader is a Reader for the PackerServiceGetRegistryTFCRunTaskAPI structure.
type PackerServiceGetRegistryTFCRunTaskAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceGetRegistryTFCRunTaskAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceGetRegistryTFCRunTaskAPIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceGetRegistryTFCRunTaskAPIDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceGetRegistryTFCRunTaskAPIOK creates a PackerServiceGetRegistryTFCRunTaskAPIOK with default headers values
func NewPackerServiceGetRegistryTFCRunTaskAPIOK() *PackerServiceGetRegistryTFCRunTaskAPIOK {
	return &PackerServiceGetRegistryTFCRunTaskAPIOK{}
}

/*PackerServiceGetRegistryTFCRunTaskAPIOK handles this case with default header values.

A successful response.
*/
type PackerServiceGetRegistryTFCRunTaskAPIOK struct {
	Payload *models.HashicorpCloudPackerGetRegistryTFCRunTaskAPIResponse
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/{task_type}][%d] packerServiceGetRegistryTFCRunTaskApiOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) GetPayload() *models.HashicorpCloudPackerGetRegistryTFCRunTaskAPIResponse {
	return o.Payload
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerGetRegistryTFCRunTaskAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceGetRegistryTFCRunTaskAPIDefault creates a PackerServiceGetRegistryTFCRunTaskAPIDefault with default headers values
func NewPackerServiceGetRegistryTFCRunTaskAPIDefault(code int) *PackerServiceGetRegistryTFCRunTaskAPIDefault {
	return &PackerServiceGetRegistryTFCRunTaskAPIDefault{
		_statusCode: code,
	}
}

/*PackerServiceGetRegistryTFCRunTaskAPIDefault handles this case with default header values.

An unexpected error response.
*/
type PackerServiceGetRegistryTFCRunTaskAPIDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the packer service get registry t f c run task API default response
func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) Code() int {
	return o._statusCode
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/runtasks/{task_type}][%d] PackerService_GetRegistryTFCRunTaskAPI default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *PackerServiceGetRegistryTFCRunTaskAPIDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
