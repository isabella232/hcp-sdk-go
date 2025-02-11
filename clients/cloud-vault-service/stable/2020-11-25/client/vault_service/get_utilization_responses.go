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

// GetUtilizationReader is a Reader for the GetUtilization structure.
type GetUtilizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUtilizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUtilizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetUtilizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUtilizationOK creates a GetUtilizationOK with default headers values
func NewGetUtilizationOK() *GetUtilizationOK {
	return &GetUtilizationOK{}
}

/*GetUtilizationOK handles this case with default header values.

A successful response.
*/
type GetUtilizationOK struct {
	Payload *models.HashicorpCloudVault20201125GetUtilizationResponse
}

func (o *GetUtilizationOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/utilization][%d] getUtilizationOK  %+v", 200, o.Payload)
}

func (o *GetUtilizationOK) GetPayload() *models.HashicorpCloudVault20201125GetUtilizationResponse {
	return o.Payload
}

func (o *GetUtilizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125GetUtilizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUtilizationDefault creates a GetUtilizationDefault with default headers values
func NewGetUtilizationDefault(code int) *GetUtilizationDefault {
	return &GetUtilizationDefault{
		_statusCode: code,
	}
}

/*GetUtilizationDefault handles this case with default header values.

An unexpected error response.
*/
type GetUtilizationDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get utilization default response
func (o *GetUtilizationDefault) Code() int {
	return o._statusCode
}

func (o *GetUtilizationDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/utilization][%d] GetUtilization default  %+v", o._statusCode, o.Payload)
}

func (o *GetUtilizationDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetUtilizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
