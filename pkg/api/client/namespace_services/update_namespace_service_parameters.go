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

// NewUpdateNamespaceServiceParams creates a new UpdateNamespaceServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNamespaceServiceParams() *UpdateNamespaceServiceParams {
	return &UpdateNamespaceServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNamespaceServiceParamsWithTimeout creates a new UpdateNamespaceServiceParams object
// with the ability to set a timeout on a request.
func NewUpdateNamespaceServiceParamsWithTimeout(timeout time.Duration) *UpdateNamespaceServiceParams {
	return &UpdateNamespaceServiceParams{
		timeout: timeout,
	}
}

// NewUpdateNamespaceServiceParamsWithContext creates a new UpdateNamespaceServiceParams object
// with the ability to set a context for a request.
func NewUpdateNamespaceServiceParamsWithContext(ctx context.Context) *UpdateNamespaceServiceParams {
	return &UpdateNamespaceServiceParams{
		Context: ctx,
	}
}

// NewUpdateNamespaceServiceParamsWithHTTPClient creates a new UpdateNamespaceServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNamespaceServiceParamsWithHTTPClient(client *http.Client) *UpdateNamespaceServiceParams {
	return &UpdateNamespaceServiceParams{
		HTTPClient: client,
	}
}

/* UpdateNamespaceServiceParams contains all the parameters to send to the API endpoint
   for the update namespace service operation.

   Typically these are written to a http.Request.
*/
type UpdateNamespaceServiceParams struct {

	/* Service.

	   Payload that contains information on service revision
	*/
	Service UpdateNamespaceServiceBody

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

// WithDefaults hydrates default values in the update namespace service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNamespaceServiceParams) WithDefaults() *UpdateNamespaceServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update namespace service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNamespaceServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update namespace service params
func (o *UpdateNamespaceServiceParams) WithTimeout(timeout time.Duration) *UpdateNamespaceServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update namespace service params
func (o *UpdateNamespaceServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update namespace service params
func (o *UpdateNamespaceServiceParams) WithContext(ctx context.Context) *UpdateNamespaceServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update namespace service params
func (o *UpdateNamespaceServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update namespace service params
func (o *UpdateNamespaceServiceParams) WithHTTPClient(client *http.Client) *UpdateNamespaceServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update namespace service params
func (o *UpdateNamespaceServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithService adds the service to the update namespace service params
func (o *UpdateNamespaceServiceParams) WithService(service UpdateNamespaceServiceBody) *UpdateNamespaceServiceParams {
	o.SetService(service)
	return o
}

// SetService adds the service to the update namespace service params
func (o *UpdateNamespaceServiceParams) SetService(service UpdateNamespaceServiceBody) {
	o.Service = service
}

// WithNamespace adds the namespace to the update namespace service params
func (o *UpdateNamespaceServiceParams) WithNamespace(namespace string) *UpdateNamespaceServiceParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the update namespace service params
func (o *UpdateNamespaceServiceParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithServiceName adds the serviceName to the update namespace service params
func (o *UpdateNamespaceServiceParams) WithServiceName(serviceName string) *UpdateNamespaceServiceParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the update namespace service params
func (o *UpdateNamespaceServiceParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNamespaceServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Service); err != nil {
		return err
	}

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