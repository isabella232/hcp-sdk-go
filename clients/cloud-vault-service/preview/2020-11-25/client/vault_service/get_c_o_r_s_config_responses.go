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
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/preview/2020-11-25/models"
)

// GetCORSConfigReader is a Reader for the GetCORSConfig structure.
type GetCORSConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCORSConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCORSConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCORSConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCORSConfigOK creates a GetCORSConfigOK with default headers values
func NewGetCORSConfigOK() *GetCORSConfigOK {
	return &GetCORSConfigOK{}
}

/* GetCORSConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetCORSConfigOK struct {
	Payload *models.HashicorpCloudVault20201125GetCORSConfigResponse
}

func (o *GetCORSConfigOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/get-cors-config][%d] getCORSConfigOK  %+v", 200, o.Payload)
}
func (o *GetCORSConfigOK) GetPayload() *models.HashicorpCloudVault20201125GetCORSConfigResponse {
	return o.Payload
}

func (o *GetCORSConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125GetCORSConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCORSConfigDefault creates a GetCORSConfigDefault with default headers values
func NewGetCORSConfigDefault(code int) *GetCORSConfigDefault {
	return &GetCORSConfigDefault{
		_statusCode: code,
	}
}

/* GetCORSConfigDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetCORSConfigDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get c o r s config default response
func (o *GetCORSConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetCORSConfigDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/get-cors-config][%d] GetCORSConfig default  %+v", o._statusCode, o.Payload)
}
func (o *GetCORSConfigDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetCORSConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
