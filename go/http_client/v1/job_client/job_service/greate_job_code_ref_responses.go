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

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GreateJobCodeRefReader is a Reader for the GreateJobCodeRef structure.
type GreateJobCodeRefReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GreateJobCodeRefReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGreateJobCodeRefOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGreateJobCodeRefNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGreateJobCodeRefOK creates a GreateJobCodeRefOK with default headers values
func NewGreateJobCodeRefOK() *GreateJobCodeRefOK {
	return &GreateJobCodeRefOK{}
}

/*GreateJobCodeRefOK handles this case with default header values.

A successful response.
*/
type GreateJobCodeRefOK struct {
	Payload interface{}
}

func (o *GreateJobCodeRefOK) Error() string {
	return fmt.Sprintf("[POST /v1/{owner}/{project}/jobs/{id}/coderef][%d] greateJobCodeRefOK  %+v", 200, o.Payload)
}

func (o *GreateJobCodeRefOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GreateJobCodeRefOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGreateJobCodeRefNotFound creates a GreateJobCodeRefNotFound with default headers values
func NewGreateJobCodeRefNotFound() *GreateJobCodeRefNotFound {
	return &GreateJobCodeRefNotFound{}
}

/*GreateJobCodeRefNotFound handles this case with default header values.

Resource does not exist.
*/
type GreateJobCodeRefNotFound struct {
	Payload string
}

func (o *GreateJobCodeRefNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/{owner}/{project}/jobs/{id}/coderef][%d] greateJobCodeRefNotFound  %+v", 404, o.Payload)
}

func (o *GreateJobCodeRefNotFound) GetPayload() string {
	return o.Payload
}

func (o *GreateJobCodeRefNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
