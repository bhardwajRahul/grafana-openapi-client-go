// Code generated by go-swagger; DO NOT EDIT.

package group_attribute_sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// New creates a new group attribute sync API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group attribute sync API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateGroupMappings(groupID string, body *models.GroupAttributes, opts ...ClientOption) (*CreateGroupMappingsCreated, error)
	CreateGroupMappingsWithParams(params *CreateGroupMappingsParams, opts ...ClientOption) (*CreateGroupMappingsCreated, error)

	DeleteGroupMappings(groupID string, opts ...ClientOption) (*DeleteGroupMappingsNoContent, error)
	DeleteGroupMappingsWithParams(params *DeleteGroupMappingsParams, opts ...ClientOption) (*DeleteGroupMappingsNoContent, error)

	GetGroupRoles(groupID string, opts ...ClientOption) (*GetGroupRolesOK, error)
	GetGroupRolesWithParams(params *GetGroupRolesParams, opts ...ClientOption) (*GetGroupRolesOK, error)

	GetMappedGroups(opts ...ClientOption) (*GetMappedGroupsOK, error)
	GetMappedGroupsWithParams(params *GetMappedGroupsParams, opts ...ClientOption) (*GetMappedGroupsOK, error)

	UpdateGroupMappings(groupID string, body *models.GroupAttributes, opts ...ClientOption) (*UpdateGroupMappingsCreated, error)
	UpdateGroupMappingsWithParams(params *UpdateGroupMappingsParams, opts ...ClientOption) (*UpdateGroupMappingsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateGroupMappings creates mappings for a group this endpoint is behind the feature flag group attribute sync and is considered experimental
*/
func (a *Client) CreateGroupMappings(groupID string, body *models.GroupAttributes, opts ...ClientOption) (*CreateGroupMappingsCreated, error) {
	params := NewCreateGroupMappingsParams().WithBody(body).WithGroupID(groupID)
	return a.CreateGroupMappingsWithParams(params, opts...)
}

func (a *Client) CreateGroupMappingsWithParams(params *CreateGroupMappingsParams, opts ...ClientOption) (*CreateGroupMappingsCreated, error) {
	if params == nil {
		params = NewCreateGroupMappingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createGroupMappings",
		Method:             "POST",
		PathPattern:        "/groupsync/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateGroupMappingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateGroupMappingsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createGroupMappings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteGroupMappings deletes mappings for a group this endpoint is behind the feature flag group attribute sync and is considered experimental
*/
func (a *Client) DeleteGroupMappings(groupID string, opts ...ClientOption) (*DeleteGroupMappingsNoContent, error) {
	params := NewDeleteGroupMappingsParams().WithGroupID(groupID)
	return a.DeleteGroupMappingsWithParams(params, opts...)
}

func (a *Client) DeleteGroupMappingsWithParams(params *DeleteGroupMappingsParams, opts ...ClientOption) (*DeleteGroupMappingsNoContent, error) {
	if params == nil {
		params = NewDeleteGroupMappingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteGroupMappings",
		Method:             "DELETE",
		PathPattern:        "/groupsync/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteGroupMappingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteGroupMappingsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteGroupMappings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGroupRoles gets roles mapped to a group this endpoint is behind the feature flag group attribute sync and is considered experimental
*/
func (a *Client) GetGroupRoles(groupID string, opts ...ClientOption) (*GetGroupRolesOK, error) {
	params := NewGetGroupRolesParams().WithGroupID(groupID)
	return a.GetGroupRolesWithParams(params, opts...)
}

func (a *Client) GetGroupRolesWithParams(params *GetGroupRolesParams, opts ...ClientOption) (*GetGroupRolesOK, error) {
	if params == nil {
		params = NewGetGroupRolesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getGroupRoles",
		Method:             "GET",
		PathPattern:        "/groupsync/groups/{group_id}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGroupRolesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGroupRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGroupRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMappedGroups lists groups that have mappings set this endpoint is behind the feature flag group attribute sync and is considered experimental
*/
func (a *Client) GetMappedGroups(opts ...ClientOption) (*GetMappedGroupsOK, error) {
	params := NewGetMappedGroupsParams()
	return a.GetMappedGroupsWithParams(params, opts...)
}

func (a *Client) GetMappedGroupsWithParams(params *GetMappedGroupsParams, opts ...ClientOption) (*GetMappedGroupsOK, error) {
	if params == nil {
		params = NewGetMappedGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getMappedGroups",
		Method:             "GET",
		PathPattern:        "/groupsync/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMappedGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMappedGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMappedGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateGroupMappings updates mappings for a group this endpoint is behind the feature flag group attribute sync and is considered experimental
*/
func (a *Client) UpdateGroupMappings(groupID string, body *models.GroupAttributes, opts ...ClientOption) (*UpdateGroupMappingsCreated, error) {
	params := NewUpdateGroupMappingsParams().WithBody(body).WithGroupID(groupID)
	return a.UpdateGroupMappingsWithParams(params, opts...)
}

func (a *Client) UpdateGroupMappingsWithParams(params *UpdateGroupMappingsParams, opts ...ClientOption) (*UpdateGroupMappingsCreated, error) {
	if params == nil {
		params = NewUpdateGroupMappingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateGroupMappings",
		Method:             "PUT",
		PathPattern:        "/groupsync/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateGroupMappingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(op)
		}
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateGroupMappingsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateGroupMappings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// WithAuthInfo changes the transport on the client
func WithAuthInfo(authInfo runtime.ClientAuthInfoWriter) ClientOption {
	return func(op *runtime.ClientOperation) {
		op.AuthInfo = authInfo
	}
}
