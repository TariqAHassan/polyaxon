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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetRunCodeRefsReader is a Reader for the GetRunCodeRefs structure.
type GetRunCodeRefsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunCodeRefsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunCodeRefsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetRunCodeRefsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRunCodeRefsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRunCodeRefsOK creates a GetRunCodeRefsOK with default headers values
func NewGetRunCodeRefsOK() *GetRunCodeRefsOK {
	return &GetRunCodeRefsOK{}
}

/*GetRunCodeRefsOK handles this case with default header values.

A successful response.
*/
type GetRunCodeRefsOK struct {
	Payload *service_model.V1ListCodeRefsResponse
}

func (o *GetRunCodeRefsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/coderef][%d] getRunCodeRefsOK  %+v", 200, o.Payload)
}

func (o *GetRunCodeRefsOK) GetPayload() *service_model.V1ListCodeRefsResponse {
	return o.Payload
}

func (o *GetRunCodeRefsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListCodeRefsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunCodeRefsForbidden creates a GetRunCodeRefsForbidden with default headers values
func NewGetRunCodeRefsForbidden() *GetRunCodeRefsForbidden {
	return &GetRunCodeRefsForbidden{}
}

/*GetRunCodeRefsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetRunCodeRefsForbidden struct {
	Payload interface{}
}

func (o *GetRunCodeRefsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/coderef][%d] getRunCodeRefsForbidden  %+v", 403, o.Payload)
}

func (o *GetRunCodeRefsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRunCodeRefsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunCodeRefsNotFound creates a GetRunCodeRefsNotFound with default headers values
func NewGetRunCodeRefsNotFound() *GetRunCodeRefsNotFound {
	return &GetRunCodeRefsNotFound{}
}

/*GetRunCodeRefsNotFound handles this case with default header values.

Resource does not exist.
*/
type GetRunCodeRefsNotFound struct {
	Payload string
}

func (o *GetRunCodeRefsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/runs/{uuid}/coderef][%d] getRunCodeRefsNotFound  %+v", 404, o.Payload)
}

func (o *GetRunCodeRefsNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetRunCodeRefsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
