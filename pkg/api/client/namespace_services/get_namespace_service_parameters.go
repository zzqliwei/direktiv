// Code generated by go-swagger; DO NOT EDIT.

package namespace_services

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

// NewGetNamespaceServiceParams creates a new GetNamespaceServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNamespaceServiceParams() *GetNamespaceServiceParams {
	return &GetNamespaceServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNamespaceServiceParamsWithTimeout creates a new GetNamespaceServiceParams object
// with the ability to set a timeout on a request.
func NewGetNamespaceServiceParamsWithTimeout(timeout time.Duration) *GetNamespaceServiceParams {
	return &GetNamespaceServiceParams{
		timeout: timeout,
	}
}

// NewGetNamespaceServiceParamsWithContext creates a new GetNamespaceServiceParams object
// with the ability to set a context for a request.
func NewGetNamespaceServiceParamsWithContext(ctx context.Context) *GetNamespaceServiceParams {
	return &GetNamespaceServiceParams{
		Context: ctx,
	}
}

// NewGetNamespaceServiceParamsWithHTTPClient creates a new GetNamespaceServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNamespaceServiceParamsWithHTTPClient(client *http.Client) *GetNamespaceServiceParams {
	return &GetNamespaceServiceParams{
		HTTPClient: client,
	}
}

/* GetNamespaceServiceParams contains all the parameters to send to the API endpoint
   for the get namespace service operation.

   Typically these are written to a http.Request.
*/
type GetNamespaceServiceParams struct {

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* ServiceName.

	   target service name
	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get namespace service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamespaceServiceParams) WithDefaults() *GetNamespaceServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get namespace service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNamespaceServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get namespace service params
func (o *GetNamespaceServiceParams) WithTimeout(timeout time.Duration) *GetNamespaceServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get namespace service params
func (o *GetNamespaceServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get namespace service params
func (o *GetNamespaceServiceParams) WithContext(ctx context.Context) *GetNamespaceServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get namespace service params
func (o *GetNamespaceServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get namespace service params
func (o *GetNamespaceServiceParams) WithHTTPClient(client *http.Client) *GetNamespaceServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get namespace service params
func (o *GetNamespaceServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the get namespace service params
func (o *GetNamespaceServiceParams) WithNamespace(namespace string) *GetNamespaceServiceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get namespace service params
func (o *GetNamespaceServiceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithServiceName adds the serviceName to the get namespace service params
func (o *GetNamespaceServiceParams) WithServiceName(serviceName string) *GetNamespaceServiceParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get namespace service params
func (o *GetNamespaceServiceParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetNamespaceServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
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
