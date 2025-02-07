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

package artifacts_stores_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetArtifactsStoreReader is a Reader for the GetArtifactsStore structure.
type GetArtifactsStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactsStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArtifactsStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetArtifactsStoreNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetArtifactsStoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArtifactsStoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetArtifactsStoreOK creates a GetArtifactsStoreOK with default headers values
func NewGetArtifactsStoreOK() *GetArtifactsStoreOK {
	return &GetArtifactsStoreOK{}
}

/*GetArtifactsStoreOK handles this case with default header values.

A successful response.
*/
type GetArtifactsStoreOK struct {
	Payload *service_model.V1ArtifactsStore
}

func (o *GetArtifactsStoreOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/artifacts_stores/{uuid}][%d] getArtifactsStoreOK  %+v", 200, o.Payload)
}

func (o *GetArtifactsStoreOK) GetPayload() *service_model.V1ArtifactsStore {
	return o.Payload
}

func (o *GetArtifactsStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ArtifactsStore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactsStoreNoContent creates a GetArtifactsStoreNoContent with default headers values
func NewGetArtifactsStoreNoContent() *GetArtifactsStoreNoContent {
	return &GetArtifactsStoreNoContent{}
}

/*GetArtifactsStoreNoContent handles this case with default header values.

No content.
*/
type GetArtifactsStoreNoContent struct {
	Payload interface{}
}

func (o *GetArtifactsStoreNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/artifacts_stores/{uuid}][%d] getArtifactsStoreNoContent  %+v", 204, o.Payload)
}

func (o *GetArtifactsStoreNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetArtifactsStoreNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactsStoreForbidden creates a GetArtifactsStoreForbidden with default headers values
func NewGetArtifactsStoreForbidden() *GetArtifactsStoreForbidden {
	return &GetArtifactsStoreForbidden{}
}

/*GetArtifactsStoreForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetArtifactsStoreForbidden struct {
	Payload interface{}
}

func (o *GetArtifactsStoreForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/artifacts_stores/{uuid}][%d] getArtifactsStoreForbidden  %+v", 403, o.Payload)
}

func (o *GetArtifactsStoreForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetArtifactsStoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArtifactsStoreNotFound creates a GetArtifactsStoreNotFound with default headers values
func NewGetArtifactsStoreNotFound() *GetArtifactsStoreNotFound {
	return &GetArtifactsStoreNotFound{}
}

/*GetArtifactsStoreNotFound handles this case with default header values.

Resource does not exist.
*/
type GetArtifactsStoreNotFound struct {
	Payload interface{}
}

func (o *GetArtifactsStoreNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/artifacts_stores/{uuid}][%d] getArtifactsStoreNotFound  %+v", 404, o.Payload)
}

func (o *GetArtifactsStoreNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetArtifactsStoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
