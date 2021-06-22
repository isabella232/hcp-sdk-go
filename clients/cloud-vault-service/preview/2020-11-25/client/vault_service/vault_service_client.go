// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vault service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vault service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOK, error)

	CreateSnapshot(params *CreateSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSnapshotOK, error)

	Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, error)

	DeleteSnapshot(params *DeleteSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSnapshotOK, error)

	DisableCORS(params *DisableCORSParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableCORSOK, error)

	FetchAuditLog(params *FetchAuditLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchAuditLogOK, error)

	Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error)

	GetAdminToken(params *GetAdminTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdminTokenOK, error)

	GetAuditLogStatus(params *GetAuditLogStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuditLogStatusOK, error)

	GetCORSConfig(params *GetCORSConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCORSConfigOK, error)

	GetClientCounts(params *GetClientCountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClientCountsOK, error)

	GetSnapshot(params *GetSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSnapshotOK, error)

	List(params *ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOK, error)

	ListSnapshots(params *ListSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSnapshotsOK, error)

	RestoreSnapshot(params *RestoreSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreSnapshotOK, error)

	Seal(params *SealParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SealOK, error)

	Unseal(params *UnsealParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnsealOK, error)

	UpdateCORSConfig(params *UpdateCORSConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCORSConfigOK, error)

	UpdatePublicIps(params *UpdatePublicIpsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePublicIpsOK, error)

	UpdateSnapshot(params *UpdateSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSnapshotOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Create create API
*/
func (a *Client) Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Create",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{cluster.location.organization_id}/projects/{cluster.location.project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
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
	success, ok := result.(*CreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CreateSnapshot create snapshot API
*/
func (a *Client) CreateSnapshot(params *CreateSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateSnapshot",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSnapshotReader{formats: a.formats},
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
	success, ok := result.(*CreateSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Delete delete API
*/
func (a *Client) Delete(params *DeleteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Delete",
		Method:             "DELETE",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteReader{formats: a.formats},
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
	success, ok := result.(*DeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSnapshot delete snapshot API
*/
func (a *Client) DeleteSnapshot(params *DeleteSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSnapshot",
		Method:             "DELETE",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/snapshots/{snapshot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSnapshotReader{formats: a.formats},
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
	success, ok := result.(*DeleteSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DisableCORS disable c o r s API
*/
func (a *Client) DisableCORS(params *DisableCORSParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableCORSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableCORSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DisableCORS",
		Method:             "DELETE",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/disable-cors",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableCORSReader{formats: a.formats},
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
	success, ok := result.(*DisableCORSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DisableCORSDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FetchAuditLog fetch audit log API
*/
func (a *Client) FetchAuditLog(params *FetchAuditLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FetchAuditLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFetchAuditLogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "FetchAuditLog",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FetchAuditLogReader{formats: a.formats},
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
	success, ok := result.(*FetchAuditLogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FetchAuditLogDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Get get API
*/
func (a *Client) Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
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
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAdminToken get admin token API
*/
func (a *Client) GetAdminToken(params *GetAdminTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAdminTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdminTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAdminToken",
		Method:             "GET",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/admintoken",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAdminTokenReader{formats: a.formats},
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
	success, ok := result.(*GetAdminTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAdminTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAuditLogStatus get audit log status API
*/
func (a *Client) GetAuditLogStatus(params *GetAuditLogStatusParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuditLogStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuditLogStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAuditLogStatus",
		Method:             "GET",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog/{log_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAuditLogStatusReader{formats: a.formats},
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
	success, ok := result.(*GetAuditLogStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAuditLogStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetCORSConfig get c o r s config API
*/
func (a *Client) GetCORSConfig(params *GetCORSConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCORSConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCORSConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetCORSConfig",
		Method:             "GET",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/get-cors-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCORSConfigReader{formats: a.formats},
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
	success, ok := result.(*GetCORSConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCORSConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetClientCounts get client counts API
*/
func (a *Client) GetClientCounts(params *GetClientCountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetClientCountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClientCountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetClientCounts",
		Method:             "GET",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClientCountsReader{formats: a.formats},
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
	success, ok := result.(*GetClientCountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetClientCountsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSnapshot get snapshot API
*/
func (a *Client) GetSnapshot(params *GetSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSnapshot",
		Method:             "GET",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/snapshots/{snapshot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSnapshotReader{formats: a.formats},
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
	success, ok := result.(*GetSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  List list API
*/
func (a *Client) List(params *ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "List",
		Method:             "GET",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
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
	success, ok := result.(*ListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListSnapshots list snapshots API
*/
func (a *Client) ListSnapshots(params *ListSnapshotsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSnapshotsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSnapshotsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListSnapshots",
		Method:             "GET",
		PathPattern:        "/vault/2020-11-25/organizations/{resource.location.organization_id}/projects/{resource.location.project_id}/snapshots",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSnapshotsReader{formats: a.formats},
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
	success, ok := result.(*ListSnapshotsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSnapshotsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RestoreSnapshot restore snapshot API
*/
func (a *Client) RestoreSnapshot(params *RestoreSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestoreSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestoreSnapshot",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RestoreSnapshotReader{formats: a.formats},
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
	success, ok := result.(*RestoreSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestoreSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Seal seal API
*/
func (a *Client) Seal(params *SealParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SealOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSealParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Seal",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/seal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SealReader{formats: a.formats},
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
	success, ok := result.(*SealOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SealDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Unseal unseal API
*/
func (a *Client) Unseal(params *UnsealParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UnsealOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnsealParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Unseal",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/unseal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UnsealReader{formats: a.formats},
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
	success, ok := result.(*UnsealOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnsealDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateCORSConfig update c o r s config API
*/
func (a *Client) UpdateCORSConfig(params *UpdateCORSConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateCORSConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCORSConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateCORSConfig",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/update-cors-config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCORSConfigReader{formats: a.formats},
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
	success, ok := result.(*UpdateCORSConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateCORSConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdatePublicIps update public ips API
*/
func (a *Client) UpdatePublicIps(params *UpdatePublicIpsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePublicIpsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePublicIpsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdatePublicIps",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/public-ips",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePublicIpsReader{formats: a.formats},
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
	success, ok := result.(*UpdatePublicIpsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdatePublicIpsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateSnapshot update snapshot API
*/
func (a *Client) UpdateSnapshot(params *UpdateSnapshotParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSnapshotOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSnapshotParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateSnapshot",
		Method:             "POST",
		PathPattern:        "/vault/2020-11-25/organizations/{snapshot.location.organization_id}/projects/{snapshot.location.project_id}/snapshots/{snapshot.snapshot_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSnapshotReader{formats: a.formats},
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
	success, ok := result.(*UpdateSnapshotOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSnapshotDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
