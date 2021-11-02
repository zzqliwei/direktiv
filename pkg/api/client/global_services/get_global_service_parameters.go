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

// NewGetGlobalServiceParams creates a new GetGlobalServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGlobalServiceParams() *GetGlobalServiceParams {
	return &GetGlobalServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalServiceParamsWithTimeout creates a new GetGlobalServiceParams object
// with the ability to set a timeout on a request.
func NewGetGlobalServiceParamsWithTimeout(timeout time.Duration) *GetGlobalServiceParams {
	return &GetGlobalServiceParams{
		timeout: timeout,
	}
}

// NewGetGlobalServiceParamsWithContext creates a new GetGlobalServiceParams object
// with the ability to set a context for a request.
func NewGetGlobalServiceParamsWithContext(ctx context.Context) *GetGlobalServiceParams {
	return &GetGlobalServiceParams{
		Context: ctx,
	}
}

// NewGetGlobalServiceParamsWithHTTPClient creates a new GetGlobalServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGlobalServiceParamsWithHTTPClient(client *http.Client) *GetGlobalServiceParams {
	return &GetGlobalServiceParams{
		HTTPClient: client,
	}
}

/* GetGlobalServiceParams contains all the parameters to send to the API endpoint
   for the get global service operation.

   Typically these are written to a http.Request.
*/
type GetGlobalServiceParams struct {

	/* ServiceName.

	   target service name
	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get global service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalServiceParams) WithDefaults() *GetGlobalServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get global service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get global service params
func (o *GetGlobalServiceParams) WithTimeout(timeout time.Duration) *GetGlobalServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global service params
func (o *GetGlobalServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global service params
func (o *GetGlobalServiceParams) WithContext(ctx context.Context) *GetGlobalServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global service params
func (o *GetGlobalServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global service params
func (o *GetGlobalServiceParams) WithHTTPClient(client *http.Client) *GetGlobalServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global service params
func (o *GetGlobalServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the get global service params
func (o *GetGlobalServiceParams) WithServiceName(serviceName string) *GetGlobalServiceParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the get global service params
func (o *GetGlobalServiceParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serviceName
	if err := r.SetPathParam("serviceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
