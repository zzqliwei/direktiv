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

// NewUpdateGlobalServiceTrafficParams creates a new UpdateGlobalServiceTrafficParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateGlobalServiceTrafficParams() *UpdateGlobalServiceTrafficParams {
	return &UpdateGlobalServiceTrafficParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGlobalServiceTrafficParamsWithTimeout creates a new UpdateGlobalServiceTrafficParams object
// with the ability to set a timeout on a request.
func NewUpdateGlobalServiceTrafficParamsWithTimeout(timeout time.Duration) *UpdateGlobalServiceTrafficParams {
	return &UpdateGlobalServiceTrafficParams{
		timeout: timeout,
	}
}

// NewUpdateGlobalServiceTrafficParamsWithContext creates a new UpdateGlobalServiceTrafficParams object
// with the ability to set a context for a request.
func NewUpdateGlobalServiceTrafficParamsWithContext(ctx context.Context) *UpdateGlobalServiceTrafficParams {
	return &UpdateGlobalServiceTrafficParams{
		Context: ctx,
	}
}

// NewUpdateGlobalServiceTrafficParamsWithHTTPClient creates a new UpdateGlobalServiceTrafficParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateGlobalServiceTrafficParamsWithHTTPClient(client *http.Client) *UpdateGlobalServiceTrafficParams {
	return &UpdateGlobalServiceTrafficParams{
		HTTPClient: client,
	}
}

/* UpdateGlobalServiceTrafficParams contains all the parameters to send to the API endpoint
   for the update global service traffic operation.

   Typically these are written to a http.Request.
*/
type UpdateGlobalServiceTrafficParams struct {

	/* ServiceTraffic.

	   Payload that contains information on service traffic
	*/
	ServiceTraffic UpdateGlobalServiceTrafficBody

	/* ServiceName.

	   target service name
	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update global service traffic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGlobalServiceTrafficParams) WithDefaults() *UpdateGlobalServiceTrafficParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update global service traffic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateGlobalServiceTrafficParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) WithTimeout(timeout time.Duration) *UpdateGlobalServiceTrafficParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) WithContext(ctx context.Context) *UpdateGlobalServiceTrafficParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) WithHTTPClient(client *http.Client) *UpdateGlobalServiceTrafficParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceTraffic adds the serviceTraffic to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) WithServiceTraffic(serviceTraffic UpdateGlobalServiceTrafficBody) *UpdateGlobalServiceTrafficParams {
	o.SetServiceTraffic(serviceTraffic)
	return o
}

// SetServiceTraffic adds the serviceTraffic to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) SetServiceTraffic(serviceTraffic UpdateGlobalServiceTrafficBody) {
	o.ServiceTraffic = serviceTraffic
}

// WithServiceName adds the serviceName to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) WithServiceName(serviceName string) *UpdateGlobalServiceTrafficParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the update global service traffic params
func (o *UpdateGlobalServiceTrafficParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGlobalServiceTrafficParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.ServiceTraffic); err != nil {
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
