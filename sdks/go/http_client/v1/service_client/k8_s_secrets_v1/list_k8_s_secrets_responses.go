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

package k8_s_secrets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListK8SSecretsReader is a Reader for the ListK8SSecrets structure.
type ListK8SSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListK8SSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListK8SSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListK8SSecretsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListK8SSecretsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListK8SSecretsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListK8SSecretsOK creates a ListK8SSecretsOK with default headers values
func NewListK8SSecretsOK() *ListK8SSecretsOK {
	return &ListK8SSecretsOK{}
}

/*ListK8SSecretsOK handles this case with default header values.

A successful response.
*/
type ListK8SSecretsOK struct {
	Payload *service_model.V1ListK8SResourcesResponse
}

func (o *ListK8SSecretsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets][%d] listK8SSecretsOK  %+v", 200, o.Payload)
}

func (o *ListK8SSecretsOK) GetPayload() *service_model.V1ListK8SResourcesResponse {
	return o.Payload
}

func (o *ListK8SSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListK8SResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8SSecretsNoContent creates a ListK8SSecretsNoContent with default headers values
func NewListK8SSecretsNoContent() *ListK8SSecretsNoContent {
	return &ListK8SSecretsNoContent{}
}

/*ListK8SSecretsNoContent handles this case with default header values.

No content.
*/
type ListK8SSecretsNoContent struct {
	Payload interface{}
}

func (o *ListK8SSecretsNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets][%d] listK8SSecretsNoContent  %+v", 204, o.Payload)
}

func (o *ListK8SSecretsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8SSecretsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8SSecretsForbidden creates a ListK8SSecretsForbidden with default headers values
func NewListK8SSecretsForbidden() *ListK8SSecretsForbidden {
	return &ListK8SSecretsForbidden{}
}

/*ListK8SSecretsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListK8SSecretsForbidden struct {
	Payload interface{}
}

func (o *ListK8SSecretsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets][%d] listK8SSecretsForbidden  %+v", 403, o.Payload)
}

func (o *ListK8SSecretsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8SSecretsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListK8SSecretsNotFound creates a ListK8SSecretsNotFound with default headers values
func NewListK8SSecretsNotFound() *ListK8SSecretsNotFound {
	return &ListK8SSecretsNotFound{}
}

/*ListK8SSecretsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListK8SSecretsNotFound struct {
	Payload interface{}
}

func (o *ListK8SSecretsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_secrets][%d] listK8SSecretsNotFound  %+v", 404, o.Payload)
}

func (o *ListK8SSecretsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListK8SSecretsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
