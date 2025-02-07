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

package k8_s_config_maps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PatchK8SConfigMapReader is a Reader for the PatchK8SConfigMap structure.
type PatchK8SConfigMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchK8SConfigMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchK8SConfigMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPatchK8SConfigMapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPatchK8SConfigMapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchK8SConfigMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchK8SConfigMapOK creates a PatchK8SConfigMapOK with default headers values
func NewPatchK8SConfigMapOK() *PatchK8SConfigMapOK {
	return &PatchK8SConfigMapOK{}
}

/*PatchK8SConfigMapOK handles this case with default header values.

A successful response.
*/
type PatchK8SConfigMapOK struct {
	Payload *service_model.V1K8SResource
}

func (o *PatchK8SConfigMapOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] patchK8SConfigMapOK  %+v", 200, o.Payload)
}

func (o *PatchK8SConfigMapOK) GetPayload() *service_model.V1K8SResource {
	return o.Payload
}

func (o *PatchK8SConfigMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1K8SResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchK8SConfigMapNoContent creates a PatchK8SConfigMapNoContent with default headers values
func NewPatchK8SConfigMapNoContent() *PatchK8SConfigMapNoContent {
	return &PatchK8SConfigMapNoContent{}
}

/*PatchK8SConfigMapNoContent handles this case with default header values.

No content.
*/
type PatchK8SConfigMapNoContent struct {
	Payload interface{}
}

func (o *PatchK8SConfigMapNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] patchK8SConfigMapNoContent  %+v", 204, o.Payload)
}

func (o *PatchK8SConfigMapNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchK8SConfigMapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchK8SConfigMapForbidden creates a PatchK8SConfigMapForbidden with default headers values
func NewPatchK8SConfigMapForbidden() *PatchK8SConfigMapForbidden {
	return &PatchK8SConfigMapForbidden{}
}

/*PatchK8SConfigMapForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PatchK8SConfigMapForbidden struct {
	Payload interface{}
}

func (o *PatchK8SConfigMapForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] patchK8SConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *PatchK8SConfigMapForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchK8SConfigMapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchK8SConfigMapNotFound creates a PatchK8SConfigMapNotFound with default headers values
func NewPatchK8SConfigMapNotFound() *PatchK8SConfigMapNotFound {
	return &PatchK8SConfigMapNotFound{}
}

/*PatchK8SConfigMapNotFound handles this case with default header values.

Resource does not exist.
*/
type PatchK8SConfigMapNotFound struct {
	Payload interface{}
}

func (o *PatchK8SConfigMapNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] patchK8SConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *PatchK8SConfigMapNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchK8SConfigMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
