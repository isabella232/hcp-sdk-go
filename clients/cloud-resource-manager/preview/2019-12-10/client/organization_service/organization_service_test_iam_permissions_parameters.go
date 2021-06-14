// Code generated by go-swagger; DO NOT EDIT.

package organization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/models"
)

// NewOrganizationServiceTestIamPermissionsParams creates a new OrganizationServiceTestIamPermissionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOrganizationServiceTestIamPermissionsParams() *OrganizationServiceTestIamPermissionsParams {
	return &OrganizationServiceTestIamPermissionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOrganizationServiceTestIamPermissionsParamsWithTimeout creates a new OrganizationServiceTestIamPermissionsParams object
// with the ability to set a timeout on a request.
func NewOrganizationServiceTestIamPermissionsParamsWithTimeout(timeout time.Duration) *OrganizationServiceTestIamPermissionsParams {
	return &OrganizationServiceTestIamPermissionsParams{
		timeout: timeout,
	}
}

// NewOrganizationServiceTestIamPermissionsParamsWithContext creates a new OrganizationServiceTestIamPermissionsParams object
// with the ability to set a context for a request.
func NewOrganizationServiceTestIamPermissionsParamsWithContext(ctx context.Context) *OrganizationServiceTestIamPermissionsParams {
	return &OrganizationServiceTestIamPermissionsParams{
		Context: ctx,
	}
}

// NewOrganizationServiceTestIamPermissionsParamsWithHTTPClient creates a new OrganizationServiceTestIamPermissionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewOrganizationServiceTestIamPermissionsParamsWithHTTPClient(client *http.Client) *OrganizationServiceTestIamPermissionsParams {
	return &OrganizationServiceTestIamPermissionsParams{
		HTTPClient: client,
	}
}

/* OrganizationServiceTestIamPermissionsParams contains all the parameters to send to the API endpoint
   for the organization service test iam permissions operation.

   Typically these are written to a http.Request.
*/
type OrganizationServiceTestIamPermissionsParams struct {

	// Body.
	Body *models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsRequest

	/* ID.

	   ID is the identifier of the organization.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the organization service test iam permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationServiceTestIamPermissionsParams) WithDefaults() *OrganizationServiceTestIamPermissionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the organization service test iam permissions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OrganizationServiceTestIamPermissionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) WithTimeout(timeout time.Duration) *OrganizationServiceTestIamPermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) WithContext(ctx context.Context) *OrganizationServiceTestIamPermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) WithHTTPClient(client *http.Client) *OrganizationServiceTestIamPermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) WithBody(body *models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsRequest) *OrganizationServiceTestIamPermissionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) SetBody(body *models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsRequest) {
	o.Body = body
}

// WithID adds the id to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) WithID(id string) *OrganizationServiceTestIamPermissionsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the organization service test iam permissions params
func (o *OrganizationServiceTestIamPermissionsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *OrganizationServiceTestIamPermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
