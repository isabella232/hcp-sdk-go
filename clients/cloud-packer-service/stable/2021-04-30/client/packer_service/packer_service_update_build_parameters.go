// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
)

// NewPackerServiceUpdateBuildParams creates a new PackerServiceUpdateBuildParams object
// with the default values initialized.
func NewPackerServiceUpdateBuildParams() *PackerServiceUpdateBuildParams {
	var ()
	return &PackerServiceUpdateBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceUpdateBuildParamsWithTimeout creates a new PackerServiceUpdateBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPackerServiceUpdateBuildParamsWithTimeout(timeout time.Duration) *PackerServiceUpdateBuildParams {
	var ()
	return &PackerServiceUpdateBuildParams{

		timeout: timeout,
	}
}

// NewPackerServiceUpdateBuildParamsWithContext creates a new PackerServiceUpdateBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewPackerServiceUpdateBuildParamsWithContext(ctx context.Context) *PackerServiceUpdateBuildParams {
	var ()
	return &PackerServiceUpdateBuildParams{

		Context: ctx,
	}
}

// NewPackerServiceUpdateBuildParamsWithHTTPClient creates a new PackerServiceUpdateBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPackerServiceUpdateBuildParamsWithHTTPClient(client *http.Client) *PackerServiceUpdateBuildParams {
	var ()
	return &PackerServiceUpdateBuildParams{
		HTTPClient: client,
	}
}

/*PackerServiceUpdateBuildParams contains all the parameters to send to the API endpoint
for the packer service update build operation typically these are written to a http.Request
*/
type PackerServiceUpdateBuildParams struct {

	/*Body*/
	Body *models.HashicorpCloudPackerUpdateBuildRequest
	/*BuildID
	  ULID of the build that should be updated.

	*/
	BuildID string
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithTimeout(timeout time.Duration) *PackerServiceUpdateBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithContext(ctx context.Context) *PackerServiceUpdateBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithHTTPClient(client *http.Client) *PackerServiceUpdateBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithBody(body *models.HashicorpCloudPackerUpdateBuildRequest) *PackerServiceUpdateBuildParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetBody(body *models.HashicorpCloudPackerUpdateBuildRequest) {
	o.Body = body
}

// WithBuildID adds the buildID to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithBuildID(buildID string) *PackerServiceUpdateBuildParams {
	o.SetBuildID(buildID)
	return o
}

// SetBuildID adds the buildId to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetBuildID(buildID string) {
	o.BuildID = buildID
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceUpdateBuildParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service update build params
func (o *PackerServiceUpdateBuildParams) WithLocationProjectID(locationProjectID string) *PackerServiceUpdateBuildParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service update build params
func (o *PackerServiceUpdateBuildParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceUpdateBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param build_id
	if err := r.SetPathParam("build_id", o.BuildID); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
