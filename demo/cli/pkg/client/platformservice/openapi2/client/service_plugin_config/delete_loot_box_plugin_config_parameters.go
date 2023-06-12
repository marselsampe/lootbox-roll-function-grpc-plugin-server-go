// Code generated by go-swagger; DO NOT EDIT.

package service_plugin_config

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
)

// NewDeleteLootBoxPluginConfigParams creates a new DeleteLootBoxPluginConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLootBoxPluginConfigParams() *DeleteLootBoxPluginConfigParams {
	return &DeleteLootBoxPluginConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLootBoxPluginConfigParamsWithTimeout creates a new DeleteLootBoxPluginConfigParams object
// with the ability to set a timeout on a request.
func NewDeleteLootBoxPluginConfigParamsWithTimeout(timeout time.Duration) *DeleteLootBoxPluginConfigParams {
	return &DeleteLootBoxPluginConfigParams{
		timeout: timeout,
	}
}

// NewDeleteLootBoxPluginConfigParamsWithContext creates a new DeleteLootBoxPluginConfigParams object
// with the ability to set a context for a request.
func NewDeleteLootBoxPluginConfigParamsWithContext(ctx context.Context) *DeleteLootBoxPluginConfigParams {
	return &DeleteLootBoxPluginConfigParams{
		Context: ctx,
	}
}

// NewDeleteLootBoxPluginConfigParamsWithHTTPClient creates a new DeleteLootBoxPluginConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLootBoxPluginConfigParamsWithHTTPClient(client *http.Client) *DeleteLootBoxPluginConfigParams {
	return &DeleteLootBoxPluginConfigParams{
		HTTPClient: client,
	}
}

/*
DeleteLootBoxPluginConfigParams contains all the parameters to send to the API endpoint

	for the delete loot box plugin config operation.

	Typically these are written to a http.Request.
*/
type DeleteLootBoxPluginConfigParams struct {

	// Namespace.
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete loot box plugin config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLootBoxPluginConfigParams) WithDefaults() *DeleteLootBoxPluginConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete loot box plugin config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLootBoxPluginConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete loot box plugin config params
func (o *DeleteLootBoxPluginConfigParams) WithTimeout(timeout time.Duration) *DeleteLootBoxPluginConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete loot box plugin config params
func (o *DeleteLootBoxPluginConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete loot box plugin config params
func (o *DeleteLootBoxPluginConfigParams) WithContext(ctx context.Context) *DeleteLootBoxPluginConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete loot box plugin config params
func (o *DeleteLootBoxPluginConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete loot box plugin config params
func (o *DeleteLootBoxPluginConfigParams) WithHTTPClient(client *http.Client) *DeleteLootBoxPluginConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete loot box plugin config params
func (o *DeleteLootBoxPluginConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete loot box plugin config params
func (o *DeleteLootBoxPluginConfigParams) WithNamespace(namespace string) *DeleteLootBoxPluginConfigParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete loot box plugin config params
func (o *DeleteLootBoxPluginConfigParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLootBoxPluginConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}