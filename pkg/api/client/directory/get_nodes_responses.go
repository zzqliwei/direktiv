// Code generated by go-swagger; DO NOT EDIT.

package directory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/direktiv/direktiv/pkg/api/models"
)

// GetNodesReader is a Reader for the GetNodes structure.
type GetNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodesOK creates a GetNodesOK with default headers values
func NewGetNodesOK() *GetNodesOK {
	return &GetNodesOK{}
}

/* GetNodesOK describes a response with status code 200, with default header values.

successfully got namespace nodes
*/
type GetNodesOK struct {
	Payload models.OkBody
}

func (o *GetNodesOK) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/tree/{nodePath}][%d] getNodesOK  %+v", 200, o.Payload)
}
func (o *GetNodesOK) GetPayload() models.OkBody {
	return o.Payload
}

func (o *GetNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesDefault creates a GetNodesDefault with default headers values
func NewGetNodesDefault(code int) *GetNodesDefault {
	return &GetNodesDefault{
		_statusCode: code,
	}
}

/* GetNodesDefault describes a response with status code -1, with default header values.

an error has occurred
*/
type GetNodesDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get nodes default response
func (o *GetNodesDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesDefault) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/tree/{nodePath}][%d] getNodes default  %+v", o._statusCode, o.Payload)
}
func (o *GetNodesDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}