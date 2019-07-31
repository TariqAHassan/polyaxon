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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	build_model "github.com/polyaxon/polyaxon-sdks/go/http_client/v1/build_model"
)

// NewDeleteBuildsParams creates a new DeleteBuildsParams object
// with the default values initialized.
func NewDeleteBuildsParams() *DeleteBuildsParams {
	var ()
	return &DeleteBuildsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBuildsParamsWithTimeout creates a new DeleteBuildsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBuildsParamsWithTimeout(timeout time.Duration) *DeleteBuildsParams {
	var ()
	return &DeleteBuildsParams{

		timeout: timeout,
	}
}

// NewDeleteBuildsParamsWithContext creates a new DeleteBuildsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBuildsParamsWithContext(ctx context.Context) *DeleteBuildsParams {
	var ()
	return &DeleteBuildsParams{

		Context: ctx,
	}
}

// NewDeleteBuildsParamsWithHTTPClient creates a new DeleteBuildsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBuildsParamsWithHTTPClient(client *http.Client) *DeleteBuildsParams {
	var ()
	return &DeleteBuildsParams{
		HTTPClient: client,
	}
}

/*DeleteBuildsParams contains all the parameters to send to the API endpoint
for the delete builds operation typically these are written to a http.Request
*/
type DeleteBuildsParams struct {

	/*Body*/
	Body *build_model.V1OwnedEntityIDRequest
	/*Owner
	  Owner of the namespace

	*/
	Owner string
	/*Project
	  Project where the experiement will be assigned

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete builds params
func (o *DeleteBuildsParams) WithTimeout(timeout time.Duration) *DeleteBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete builds params
func (o *DeleteBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete builds params
func (o *DeleteBuildsParams) WithContext(ctx context.Context) *DeleteBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete builds params
func (o *DeleteBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete builds params
func (o *DeleteBuildsParams) WithHTTPClient(client *http.Client) *DeleteBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete builds params
func (o *DeleteBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete builds params
func (o *DeleteBuildsParams) WithBody(body *build_model.V1OwnedEntityIDRequest) *DeleteBuildsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete builds params
func (o *DeleteBuildsParams) SetBody(body *build_model.V1OwnedEntityIDRequest) {
	o.Body = body
}

// WithOwner adds the owner to the delete builds params
func (o *DeleteBuildsParams) WithOwner(owner string) *DeleteBuildsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the delete builds params
func (o *DeleteBuildsParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the delete builds params
func (o *DeleteBuildsParams) WithProject(project string) *DeleteBuildsParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete builds params
func (o *DeleteBuildsParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
