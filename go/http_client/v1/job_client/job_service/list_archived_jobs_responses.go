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

	job_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/job_model"
)

// ListArchivedJobsReader is a Reader for the ListArchivedJobs structure.
type ListArchivedJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListArchivedJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListArchivedJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListArchivedJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListArchivedJobsOK creates a ListArchivedJobsOK with default headers values
func NewListArchivedJobsOK() *ListArchivedJobsOK {
	return &ListArchivedJobsOK{}
}

/*ListArchivedJobsOK handles this case with default header values.

A successful response.
*/
type ListArchivedJobsOK struct {
	Payload *job_model.V1ListJobsResponse
}

func (o *ListArchivedJobsOK) Error() string {
	return fmt.Sprintf("[GET /v1/archives/{owner}/jobs][%d] listArchivedJobsOK  %+v", 200, o.Payload)
}

func (o *ListArchivedJobsOK) GetPayload() *job_model.V1ListJobsResponse {
	return o.Payload
}

func (o *ListArchivedJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(job_model.V1ListJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListArchivedJobsNotFound creates a ListArchivedJobsNotFound with default headers values
func NewListArchivedJobsNotFound() *ListArchivedJobsNotFound {
	return &ListArchivedJobsNotFound{}
}

/*ListArchivedJobsNotFound handles this case with default header values.

Resource does not exist.
*/
type ListArchivedJobsNotFound struct {
	Payload string
}

func (o *ListArchivedJobsNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/archives/{owner}/jobs][%d] listArchivedJobsNotFound  %+v", 404, o.Payload)
}

func (o *ListArchivedJobsNotFound) GetPayload() string {
	return o.Payload
}

func (o *ListArchivedJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
