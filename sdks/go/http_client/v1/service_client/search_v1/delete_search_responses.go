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

package search_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSearchReader is a Reader for the DeleteSearch structure.
type DeleteSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteSearchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSearchOK creates a DeleteSearchOK with default headers values
func NewDeleteSearchOK() *DeleteSearchOK {
	return &DeleteSearchOK{}
}

/*DeleteSearchOK handles this case with default header values.

A successful response.
*/
type DeleteSearchOK struct {
}

func (o *DeleteSearchOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] deleteSearchOK ", 200)
}

func (o *DeleteSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSearchNoContent creates a DeleteSearchNoContent with default headers values
func NewDeleteSearchNoContent() *DeleteSearchNoContent {
	return &DeleteSearchNoContent{}
}

/*DeleteSearchNoContent handles this case with default header values.

No content.
*/
type DeleteSearchNoContent struct {
	Payload interface{}
}

func (o *DeleteSearchNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] deleteSearchNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSearchNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteSearchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSearchForbidden creates a DeleteSearchForbidden with default headers values
func NewDeleteSearchForbidden() *DeleteSearchForbidden {
	return &DeleteSearchForbidden{}
}

/*DeleteSearchForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteSearchForbidden struct {
	Payload interface{}
}

func (o *DeleteSearchForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] deleteSearchForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSearchForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSearchNotFound creates a DeleteSearchNotFound with default headers values
func NewDeleteSearchNotFound() *DeleteSearchNotFound {
	return &DeleteSearchNotFound{}
}

/*DeleteSearchNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteSearchNotFound struct {
	Payload interface{}
}

func (o *DeleteSearchNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/{project}/searches/{uuid}][%d] deleteSearchNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSearchNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
