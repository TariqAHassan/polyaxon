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

package regsitry_accesses_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRegsitryAccessParams creates a new GetRegsitryAccessParams object
// with the default values initialized.
func NewGetRegsitryAccessParams() *GetRegsitryAccessParams {
	var ()
	return &GetRegsitryAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegsitryAccessParamsWithTimeout creates a new GetRegsitryAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegsitryAccessParamsWithTimeout(timeout time.Duration) *GetRegsitryAccessParams {
	var ()
	return &GetRegsitryAccessParams{

		timeout: timeout,
	}
}

// NewGetRegsitryAccessParamsWithContext creates a new GetRegsitryAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegsitryAccessParamsWithContext(ctx context.Context) *GetRegsitryAccessParams {
	var ()
	return &GetRegsitryAccessParams{

		Context: ctx,
	}
}

// NewGetRegsitryAccessParamsWithHTTPClient creates a new GetRegsitryAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegsitryAccessParamsWithHTTPClient(client *http.Client) *GetRegsitryAccessParams {
	var ()
	return &GetRegsitryAccessParams{
		HTTPClient: client,
	}
}

/*GetRegsitryAccessParams contains all the parameters to send to the API endpoint
for the get regsitry access operation typically these are written to a http.Request
*/
type GetRegsitryAccessParams struct {

	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*UUID
	  Unique integer identifier of the entity

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get regsitry access params
func (o *GetRegsitryAccessParams) WithTimeout(timeout time.Duration) *GetRegsitryAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get regsitry access params
func (o *GetRegsitryAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get regsitry access params
func (o *GetRegsitryAccessParams) WithContext(ctx context.Context) *GetRegsitryAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get regsitry access params
func (o *GetRegsitryAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get regsitry access params
func (o *GetRegsitryAccessParams) WithHTTPClient(client *http.Client) *GetRegsitryAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get regsitry access params
func (o *GetRegsitryAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get regsitry access params
func (o *GetRegsitryAccessParams) WithOwner(owner string) *GetRegsitryAccessParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get regsitry access params
func (o *GetRegsitryAccessParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get regsitry access params
func (o *GetRegsitryAccessParams) WithUUID(uuid string) *GetRegsitryAccessParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get regsitry access params
func (o *GetRegsitryAccessParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegsitryAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
