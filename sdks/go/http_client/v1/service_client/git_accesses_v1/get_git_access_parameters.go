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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetGitAccessParams creates a new GetGitAccessParams object
// with the default values initialized.
func NewGetGitAccessParams() *GetGitAccessParams {
	var ()
	return &GetGitAccessParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGitAccessParamsWithTimeout creates a new GetGitAccessParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGitAccessParamsWithTimeout(timeout time.Duration) *GetGitAccessParams {
	var ()
	return &GetGitAccessParams{

		timeout: timeout,
	}
}

// NewGetGitAccessParamsWithContext creates a new GetGitAccessParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGitAccessParamsWithContext(ctx context.Context) *GetGitAccessParams {
	var ()
	return &GetGitAccessParams{

		Context: ctx,
	}
}

// NewGetGitAccessParamsWithHTTPClient creates a new GetGitAccessParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGitAccessParamsWithHTTPClient(client *http.Client) *GetGitAccessParams {
	var ()
	return &GetGitAccessParams{
		HTTPClient: client,
	}
}

/*GetGitAccessParams contains all the parameters to send to the API endpoint
for the get git access operation typically these are written to a http.Request
*/
type GetGitAccessParams struct {

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

// WithTimeout adds the timeout to the get git access params
func (o *GetGitAccessParams) WithTimeout(timeout time.Duration) *GetGitAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get git access params
func (o *GetGitAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get git access params
func (o *GetGitAccessParams) WithContext(ctx context.Context) *GetGitAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get git access params
func (o *GetGitAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get git access params
func (o *GetGitAccessParams) WithHTTPClient(client *http.Client) *GetGitAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get git access params
func (o *GetGitAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get git access params
func (o *GetGitAccessParams) WithOwner(owner string) *GetGitAccessParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get git access params
func (o *GetGitAccessParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithUUID adds the uuid to the get git access params
func (o *GetGitAccessParams) WithUUID(uuid string) *GetGitAccessParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get git access params
func (o *GetGitAccessParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetGitAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
