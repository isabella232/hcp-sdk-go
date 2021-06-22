// Code generated by go-swagger; DO NOT EDIT.

package resource_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new resource service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for resource service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ResourceServiceList(params *ResourceServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceServiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ResourceServiceList lists lists the resources the caller has access to
*/
func (a *Client) ResourceServiceList(params *ResourceServiceListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ResourceServiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResourceServiceListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ResourceService_List",
		Method:             "GET",
		PathPattern:        "/resource-manager/2019-12-10/resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResourceServiceListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResourceServiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ResourceServiceListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
