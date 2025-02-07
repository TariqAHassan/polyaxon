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

// ListArtifactsStoreNamesReader is a Reader for the ListArtifactsStoreNames structure.
type ListArtifactsStoreNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArtifactsStoreNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArtifactsStoreNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListArtifactsStoreNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListArtifactsStoreNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListArtifactsStoreNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListArtifactsStoreNamesOK creates a ListArtifactsStoreNamesOK with default headers values
func NewListArtifactsStoreNamesOK() *ListArtifactsStoreNamesOK {
	return &ListArtifactsStoreNamesOK{}
}

/*ListArtifactsStoreNamesOK handles this case with default header values.

A successful response.
*/
type ListArtifactsStoreNamesOK struct {
	Payload *service_model.V1ListArtifactsStoresResponse
}

func (o *ListArtifactsStoreNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/artifacts_stores/names][%d] listArtifactsStoreNamesOK  %+v", 200, o.Payload)
}

func (o *ListArtifactsStoreNamesOK) GetPayload() *service_model.V1ListArtifactsStoresResponse {
	return o.Payload
}

func (o *ListArtifactsStoreNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListArtifactsStoresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArtifactsStoreNamesNoContent creates a ListArtifactsStoreNamesNoContent with default headers values
func NewListArtifactsStoreNamesNoContent() *ListArtifactsStoreNamesNoContent {
	return &ListArtifactsStoreNamesNoContent{}
}

/*ListArtifactsStoreNamesNoContent handles this case with default header values.

No content.
*/
type ListArtifactsStoreNamesNoContent struct {
	Payload interface{}
}

func (o *ListArtifactsStoreNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/artifacts_stores/names][%d] listArtifactsStoreNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListArtifactsStoreNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArtifactsStoreNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArtifactsStoreNamesForbidden creates a ListArtifactsStoreNamesForbidden with default headers values
func NewListArtifactsStoreNamesForbidden() *ListArtifactsStoreNamesForbidden {
	return &ListArtifactsStoreNamesForbidden{}
}

/*ListArtifactsStoreNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListArtifactsStoreNamesForbidden struct {
	Payload interface{}
}

func (o *ListArtifactsStoreNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/artifacts_stores/names][%d] listArtifactsStoreNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListArtifactsStoreNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArtifactsStoreNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArtifactsStoreNamesNotFound creates a ListArtifactsStoreNamesNotFound with default headers values
func NewListArtifactsStoreNamesNotFound() *ListArtifactsStoreNamesNotFound {
	return &ListArtifactsStoreNamesNotFound{}
}

/*ListArtifactsStoreNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListArtifactsStoreNamesNotFound struct {
	Payload interface{}
}

func (o *ListArtifactsStoreNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/artifacts_stores/names][%d] listArtifactsStoreNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListArtifactsStoreNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListArtifactsStoreNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
