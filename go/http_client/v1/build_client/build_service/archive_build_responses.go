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

package build_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ArchiveBuildReader is a Reader for the ArchiveBuild structure.
type ArchiveBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewArchiveBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewArchiveBuildOK creates a ArchiveBuildOK with default headers values
func NewArchiveBuildOK() *ArchiveBuildOK {
	return &ArchiveBuildOK{}
}

/*ArchiveBuildOK handles this case with default header values.

A successful response.
*/
type ArchiveBuildOK struct {
	Payload interface{}
}

func (o *ArchiveBuildOK) Error() string {
	return fmt.Sprintf("[POST /v1/{owner}/{project}/builds/{id}/archive][%d] archiveBuildOK  %+v", 200, o.Payload)
}

func (o *ArchiveBuildOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ArchiveBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveBuildNotFound creates a ArchiveBuildNotFound with default headers values
func NewArchiveBuildNotFound() *ArchiveBuildNotFound {
	return &ArchiveBuildNotFound{}
}

/*ArchiveBuildNotFound handles this case with default header values.

Resource does not exist.
*/
type ArchiveBuildNotFound struct {
	Payload string
}

func (o *ArchiveBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/{owner}/{project}/builds/{id}/archive][%d] archiveBuildNotFound  %+v", 404, o.Payload)
}

func (o *ArchiveBuildNotFound) GetPayload() string {
	return o.Payload
}

func (o *ArchiveBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
