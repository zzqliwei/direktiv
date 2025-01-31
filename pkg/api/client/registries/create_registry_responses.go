// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateRegistryReader is a Reader for the CreateRegistry structure.
type CreateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRegistryOK creates a CreateRegistryOK with default headers values
func NewCreateRegistryOK() *CreateRegistryOK {
	return &CreateRegistryOK{}
}

/* CreateRegistryOK describes a response with status code 200, with default header values.

successfully created namespace registry
*/
type CreateRegistryOK struct {
}

func (o *CreateRegistryOK) Error() string {
	return fmt.Sprintf("[POST /api/functions/registries/namespaces/{namespace}][%d] createRegistryOK ", 200)
}

func (o *CreateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateRegistryBody create registry body
// Example: {"data":"admin:8QwFLg%D$qg*","reg":"https://prod.customreg.io"}
swagger:model CreateRegistryBody
*/
type CreateRegistryBody struct {

	// Target registry connection data containing the user and token.
	// Required: true
	Data *string `json:"data"`

	// Target registry URL
	// Required: true
	Reg *string `json:"reg"`
}

// Validate validates this create registry body
func (o *CreateRegistryBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateRegistryBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("Registry Payload"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

func (o *CreateRegistryBody) validateReg(formats strfmt.Registry) error {

	if err := validate.Required("Registry Payload"+"."+"reg", "body", o.Reg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create registry body based on context it is used
func (o *CreateRegistryBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateRegistryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRegistryBody) UnmarshalBinary(b []byte) error {
	var res CreateRegistryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
