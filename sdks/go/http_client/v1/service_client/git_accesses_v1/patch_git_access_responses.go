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

package git_accesses_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchGitAccessReader is a Reader for the PatchGitAccess structure.
type PatchGitAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchGitAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchGitAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchGitAccessNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchGitAccessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchGitAccessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchGitAccessOK creates a PatchGitAccessOK with default headers values
func NewPatchGitAccessOK() *PatchGitAccessOK {
	return &PatchGitAccessOK{}
}

/*PatchGitAccessOK handles this case with default header values.

A successful response.
*/
type PatchGitAccessOK struct {
	Payload *service_model.V1HostAccess
}

func (o *PatchGitAccessOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/git_accesses/{host_access.uuid}][%d] patchGitAccessOK  %+v", 200, o.Payload)
}

func (o *PatchGitAccessOK) GetPayload() *service_model.V1HostAccess {
	return o.Payload
}

func (o *PatchGitAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1HostAccess)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGitAccessNoContent creates a PatchGitAccessNoContent with default headers values
func NewPatchGitAccessNoContent() *PatchGitAccessNoContent {
	return &PatchGitAccessNoContent{}
}

/*PatchGitAccessNoContent handles this case with default header values.

No content.
*/
type PatchGitAccessNoContent struct {
	Payload interface{}
}

func (o *PatchGitAccessNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/git_accesses/{host_access.uuid}][%d] patchGitAccessNoContent  %+v", 204, o.Payload)
}

func (o *PatchGitAccessNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchGitAccessNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGitAccessForbidden creates a PatchGitAccessForbidden with default headers values
func NewPatchGitAccessForbidden() *PatchGitAccessForbidden {
	return &PatchGitAccessForbidden{}
}

/*PatchGitAccessForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchGitAccessForbidden struct {
	Payload interface{}
}

func (o *PatchGitAccessForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/git_accesses/{host_access.uuid}][%d] patchGitAccessForbidden  %+v", 403, o.Payload)
}

func (o *PatchGitAccessForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchGitAccessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchGitAccessNotFound creates a PatchGitAccessNotFound with default headers values
func NewPatchGitAccessNotFound() *PatchGitAccessNotFound {
	return &PatchGitAccessNotFound{}
}

/*PatchGitAccessNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchGitAccessNotFound struct {
	Payload interface{}
}

func (o *PatchGitAccessNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/git_accesses/{host_access.uuid}][%d] patchGitAccessNotFound  %+v", 404, o.Payload)
}

func (o *PatchGitAccessNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchGitAccessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
