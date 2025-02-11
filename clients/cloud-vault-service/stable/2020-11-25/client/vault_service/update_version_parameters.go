// Code generated by go-swagger; DO NOT EDIT.

package vault_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// NewUpdateVersionParams creates a new UpdateVersionParams object
// with the default values initialized.
func NewUpdateVersionParams() *UpdateVersionParams {
	var ()
	return &UpdateVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVersionParamsWithTimeout creates a new UpdateVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVersionParamsWithTimeout(timeout time.Duration) *UpdateVersionParams {
	var ()
	return &UpdateVersionParams{

		timeout: timeout,
	}
}

// NewUpdateVersionParamsWithContext creates a new UpdateVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVersionParamsWithContext(ctx context.Context) *UpdateVersionParams {
	var ()
	return &UpdateVersionParams{

		Context: ctx,
	}
}

// NewUpdateVersionParamsWithHTTPClient creates a new UpdateVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVersionParamsWithHTTPClient(client *http.Client) *UpdateVersionParams {
	var ()
	return &UpdateVersionParams{
		HTTPClient: client,
	}
}

/*UpdateVersionParams contains all the parameters to send to the API endpoint
for the update version operation typically these are written to a http.Request
*/
type UpdateVersionParams struct {

	/*Body*/
	Body *models.HashicorpCloudVault20201125UpdateVersionRequest
	/*ClusterID*/
	ClusterID string
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string
	/*Version*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update version params
func (o *UpdateVersionParams) WithTimeout(timeout time.Duration) *UpdateVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update version params
func (o *UpdateVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update version params
func (o *UpdateVersionParams) WithContext(ctx context.Context) *UpdateVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update version params
func (o *UpdateVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update version params
func (o *UpdateVersionParams) WithHTTPClient(client *http.Client) *UpdateVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update version params
func (o *UpdateVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update version params
func (o *UpdateVersionParams) WithBody(body *models.HashicorpCloudVault20201125UpdateVersionRequest) *UpdateVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update version params
func (o *UpdateVersionParams) SetBody(body *models.HashicorpCloudVault20201125UpdateVersionRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update version params
func (o *UpdateVersionParams) WithClusterID(clusterID string) *UpdateVersionParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update version params
func (o *UpdateVersionParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the update version params
func (o *UpdateVersionParams) WithLocationOrganizationID(locationOrganizationID string) *UpdateVersionParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the update version params
func (o *UpdateVersionParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the update version params
func (o *UpdateVersionParams) WithLocationProjectID(locationProjectID string) *UpdateVersionParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update version params
func (o *UpdateVersionParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithVersion adds the version to the update version params
func (o *UpdateVersionParams) WithVersion(version string) *UpdateVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update version params
func (o *UpdateVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
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

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
