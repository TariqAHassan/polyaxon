// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package regsitry_accesses_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateRegsitryAccessReader is a Reader for the CreateRegsitryAccess structure.
type CreateRegsitryAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRegsitryAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRegsitryAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateRegsitryAccessNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateRegsitryAccessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRegsitryAccessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRegsitryAccessOK creates a CreateRegsitryAccessOK with default headers values
func NewCreateRegsitryAccessOK() *CreateRegsitryAccessOK {
	return &CreateRegsitryAccessOK{}
}

/*CreateRegsitryAccessOK handles this case with default header values.

A successful response.
*/
type CreateRegsitryAccessOK struct {
	Payload *service_model.V1HostAccess
}

func (o *CreateRegsitryAccessOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry_accesses][%d] createRegsitryAccessOK  %+v", 200, o.Payload)
}

func (o *CreateRegsitryAccessOK) GetPayload() *service_model.V1HostAccess {
	return o.Payload
}

func (o *CreateRegsitryAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1HostAccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegsitryAccessNoContent creates a CreateRegsitryAccessNoContent with default headers values
func NewCreateRegsitryAccessNoContent() *CreateRegsitryAccessNoContent {
	return &CreateRegsitryAccessNoContent{}
}

/*CreateRegsitryAccessNoContent handles this case with default header values.

No content.
*/
type CreateRegsitryAccessNoContent struct {
	Payload interface{}
}

func (o *CreateRegsitryAccessNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry_accesses][%d] createRegsitryAccessNoContent  %+v", 204, o.Payload)
}

func (o *CreateRegsitryAccessNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRegsitryAccessNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegsitryAccessForbidden creates a CreateRegsitryAccessForbidden with default headers values
func NewCreateRegsitryAccessForbidden() *CreateRegsitryAccessForbidden {
	return &CreateRegsitryAccessForbidden{}
}

/*CreateRegsitryAccessForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateRegsitryAccessForbidden struct {
	Payload interface{}
}

func (o *CreateRegsitryAccessForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry_accesses][%d] createRegsitryAccessForbidden  %+v", 403, o.Payload)
}

func (o *CreateRegsitryAccessForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRegsitryAccessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRegsitryAccessNotFound creates a CreateRegsitryAccessNotFound with default headers values
func NewCreateRegsitryAccessNotFound() *CreateRegsitryAccessNotFound {
	return &CreateRegsitryAccessNotFound{}
}

/*CreateRegsitryAccessNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateRegsitryAccessNotFound struct {
	Payload interface{}
}

func (o *CreateRegsitryAccessNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/registry_accesses][%d] createRegsitryAccessNotFound  %+v", 404, o.Payload)
}

func (o *CreateRegsitryAccessNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRegsitryAccessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
