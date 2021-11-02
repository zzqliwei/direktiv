// Code generated by go-swagger; DO NOT EDIT.

package metrics

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

// NewNamespaceMetricsSuccessfulParams creates a new NamespaceMetricsSuccessfulParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNamespaceMetricsSuccessfulParams() *NamespaceMetricsSuccessfulParams {
	return &NamespaceMetricsSuccessfulParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNamespaceMetricsSuccessfulParamsWithTimeout creates a new NamespaceMetricsSuccessfulParams object
// with the ability to set a timeout on a request.
func NewNamespaceMetricsSuccessfulParamsWithTimeout(timeout time.Duration) *NamespaceMetricsSuccessfulParams {
	return &NamespaceMetricsSuccessfulParams{
		timeout: timeout,
	}
}

// NewNamespaceMetricsSuccessfulParamsWithContext creates a new NamespaceMetricsSuccessfulParams object
// with the ability to set a context for a request.
func NewNamespaceMetricsSuccessfulParamsWithContext(ctx context.Context) *NamespaceMetricsSuccessfulParams {
	return &NamespaceMetricsSuccessfulParams{
		Context: ctx,
	}
}

// NewNamespaceMetricsSuccessfulParamsWithHTTPClient creates a new NamespaceMetricsSuccessfulParams object
// with the ability to set a custom HTTPClient for a request.
func NewNamespaceMetricsSuccessfulParamsWithHTTPClient(client *http.Client) *NamespaceMetricsSuccessfulParams {
	return &NamespaceMetricsSuccessfulParams{
		HTTPClient: client,
	}
}

/* NamespaceMetricsSuccessfulParams contains all the parameters to send to the API endpoint
   for the namespace metrics successful operation.

   Typically these are written to a http.Request.
*/
type NamespaceMetricsSuccessfulParams struct {

	/* Namespace.

	   target namespace
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the namespace metrics successful params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamespaceMetricsSuccessfulParams) WithDefaults() *NamespaceMetricsSuccessfulParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the namespace metrics successful params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NamespaceMetricsSuccessfulParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the namespace metrics successful params
func (o *NamespaceMetricsSuccessfulParams) WithTimeout(timeout time.Duration) *NamespaceMetricsSuccessfulParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the namespace metrics successful params
func (o *NamespaceMetricsSuccessfulParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the namespace metrics successful params
func (o *NamespaceMetricsSuccessfulParams) WithContext(ctx context.Context) *NamespaceMetricsSuccessfulParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the namespace metrics successful params
func (o *NamespaceMetricsSuccessfulParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the namespace metrics successful params
func (o *NamespaceMetricsSuccessfulParams) WithHTTPClient(client *http.Client) *NamespaceMetricsSuccessfulParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the namespace metrics successful params
func (o *NamespaceMetricsSuccessfulParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the namespace metrics successful params
func (o *NamespaceMetricsSuccessfulParams) WithNamespace(namespace string) *NamespaceMetricsSuccessfulParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the namespace metrics successful params
func (o *NamespaceMetricsSuccessfulParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *NamespaceMetricsSuccessfulParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
