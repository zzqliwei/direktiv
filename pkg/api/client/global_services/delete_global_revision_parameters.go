// Code generated by go-swagger; DO NOT EDIT.

package global_services

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

// NewDeleteGlobalRevisionParams creates a new DeleteGlobalRevisionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGlobalRevisionParams() *DeleteGlobalRevisionParams {
	return &DeleteGlobalRevisionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGlobalRevisionParamsWithTimeout creates a new DeleteGlobalRevisionParams object
// with the ability to set a timeout on a request.
func NewDeleteGlobalRevisionParamsWithTimeout(timeout time.Duration) *DeleteGlobalRevisionParams {
	return &DeleteGlobalRevisionParams{
		timeout: timeout,
	}
}

// NewDeleteGlobalRevisionParamsWithContext creates a new DeleteGlobalRevisionParams object
// with the ability to set a context for a request.
func NewDeleteGlobalRevisionParamsWithContext(ctx context.Context) *DeleteGlobalRevisionParams {
	return &DeleteGlobalRevisionParams{
		Context: ctx,
	}
}

// NewDeleteGlobalRevisionParamsWithHTTPClient creates a new DeleteGlobalRevisionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGlobalRevisionParamsWithHTTPClient(client *http.Client) *DeleteGlobalRevisionParams {
	return &DeleteGlobalRevisionParams{
		HTTPClient: client,
	}
}

/* DeleteGlobalRevisionParams contains all the parameters to send to the API endpoint
   for the delete global revision operation.

   Typically these are written to a http.Request.
*/
type DeleteGlobalRevisionParams struct {

	/* RevisionGeneration.

	   target revision generation
	*/
	RevisionGeneration string

	/* ServiceName.

	   target service name
	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete global revision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGlobalRevisionParams) WithDefaults() *DeleteGlobalRevisionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete global revision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGlobalRevisionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete global revision params
func (o *DeleteGlobalRevisionParams) WithTimeout(timeout time.Duration) *DeleteGlobalRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete global revision params
func (o *DeleteGlobalRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete global revision params
func (o *DeleteGlobalRevisionParams) WithContext(ctx context.Context) *DeleteGlobalRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete global revision params
func (o *DeleteGlobalRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete global revision params
func (o *DeleteGlobalRevisionParams) WithHTTPClient(client *http.Client) *DeleteGlobalRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete global revision params
func (o *DeleteGlobalRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRevisionGeneration adds the revisionGeneration to the delete global revision params
func (o *DeleteGlobalRevisionParams) WithRevisionGeneration(revisionGeneration string) *DeleteGlobalRevisionParams {
	o.SetRevisionGeneration(revisionGeneration)
	return o
}

// SetRevisionGeneration adds the revisionGeneration to the delete global revision params
func (o *DeleteGlobalRevisionParams) SetRevisionGeneration(revisionGeneration string) {
	o.RevisionGeneration = revisionGeneration
}

// WithServiceName adds the serviceName to the delete global revision params
func (o *DeleteGlobalRevisionParams) WithServiceName(serviceName string) *DeleteGlobalRevisionParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the delete global revision params
func (o *DeleteGlobalRevisionParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGlobalRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param revisionGeneration
	if err := r.SetPathParam("revisionGeneration", o.RevisionGeneration); err != nil {
		return err
	}

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
