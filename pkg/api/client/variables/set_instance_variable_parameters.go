// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewSetInstanceVariableParams creates a new SetInstanceVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetInstanceVariableParams() *SetInstanceVariableParams {
	return &SetInstanceVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetInstanceVariableParamsWithTimeout creates a new SetInstanceVariableParams object
// with the ability to set a timeout on a request.
func NewSetInstanceVariableParamsWithTimeout(timeout time.Duration) *SetInstanceVariableParams {
	return &SetInstanceVariableParams{
		timeout: timeout,
	}
}

// NewSetInstanceVariableParamsWithContext creates a new SetInstanceVariableParams object
// with the ability to set a context for a request.
func NewSetInstanceVariableParamsWithContext(ctx context.Context) *SetInstanceVariableParams {
	return &SetInstanceVariableParams{
		Context: ctx,
	}
}

// NewSetInstanceVariableParamsWithHTTPClient creates a new SetInstanceVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetInstanceVariableParamsWithHTTPClient(client *http.Client) *SetInstanceVariableParams {
	return &SetInstanceVariableParams{
		HTTPClient: client,
	}
}

/* SetInstanceVariableParams contains all the parameters to send to the API endpoint
   for the set instance variable operation.

   Typically these are written to a http.Request.
*/
type SetInstanceVariableParams struct {

	/* Data.

	   Payload that contains variable data.
	*/
	Data string

	/* Instance.

	   target instance
	*/
	Instance string

	/* Namespace.

	   target namespace
	*/
	Namespace string

	/* Variable.

	   target variable
	*/
	Variable string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set instance variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetInstanceVariableParams) WithDefaults() *SetInstanceVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set instance variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetInstanceVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set instance variable params
func (o *SetInstanceVariableParams) WithTimeout(timeout time.Duration) *SetInstanceVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set instance variable params
func (o *SetInstanceVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set instance variable params
func (o *SetInstanceVariableParams) WithContext(ctx context.Context) *SetInstanceVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set instance variable params
func (o *SetInstanceVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set instance variable params
func (o *SetInstanceVariableParams) WithHTTPClient(client *http.Client) *SetInstanceVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set instance variable params
func (o *SetInstanceVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the set instance variable params
func (o *SetInstanceVariableParams) WithData(data string) *SetInstanceVariableParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the set instance variable params
func (o *SetInstanceVariableParams) SetData(data string) {
	o.Data = data
}

// WithInstance adds the instance to the set instance variable params
func (o *SetInstanceVariableParams) WithInstance(instance string) *SetInstanceVariableParams {
	o.SetInstance(instance)
	return o
}

// SetInstance adds the instance to the set instance variable params
func (o *SetInstanceVariableParams) SetInstance(instance string) {
	o.Instance = instance
}

// WithNamespace adds the namespace to the set instance variable params
func (o *SetInstanceVariableParams) WithNamespace(namespace string) *SetInstanceVariableParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the set instance variable params
func (o *SetInstanceVariableParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithVariable adds the variable to the set instance variable params
func (o *SetInstanceVariableParams) WithVariable(variable string) *SetInstanceVariableParams {
	o.SetVariable(variable)
	return o
}

// SetVariable adds the variable to the set instance variable params
func (o *SetInstanceVariableParams) SetVariable(variable string) {
	o.Variable = variable
}

// WriteToRequest writes these params to a swagger request
func (o *SetInstanceVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param instance
	if err := r.SetPathParam("instance", o.Instance); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	// path param variable
	if err := r.SetPathParam("variable", o.Variable); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
