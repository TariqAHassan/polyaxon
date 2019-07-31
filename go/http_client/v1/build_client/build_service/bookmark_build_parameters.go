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
)

// NewBookmarkBuildParams creates a new BookmarkBuildParams object
// with the default values initialized.
func NewBookmarkBuildParams() *BookmarkBuildParams {
	var ()
	return &BookmarkBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBookmarkBuildParamsWithTimeout creates a new BookmarkBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBookmarkBuildParamsWithTimeout(timeout time.Duration) *BookmarkBuildParams {
	var ()
	return &BookmarkBuildParams{

		timeout: timeout,
	}
}

// NewBookmarkBuildParamsWithContext creates a new BookmarkBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewBookmarkBuildParamsWithContext(ctx context.Context) *BookmarkBuildParams {
	var ()
	return &BookmarkBuildParams{

		Context: ctx,
	}
}

// NewBookmarkBuildParamsWithHTTPClient creates a new BookmarkBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBookmarkBuildParamsWithHTTPClient(client *http.Client) *BookmarkBuildParams {
	var ()
	return &BookmarkBuildParams{
		HTTPClient: client,
	}
}

/*BookmarkBuildParams contains all the parameters to send to the API endpoint
for the bookmark build operation typically these are written to a http.Request
*/
type BookmarkBuildParams struct {

	/*ID
	  Unique integer identifier of the entity

	*/
	ID string
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

// WithTimeout adds the timeout to the bookmark build params
func (o *BookmarkBuildParams) WithTimeout(timeout time.Duration) *BookmarkBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bookmark build params
func (o *BookmarkBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bookmark build params
func (o *BookmarkBuildParams) WithContext(ctx context.Context) *BookmarkBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bookmark build params
func (o *BookmarkBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bookmark build params
func (o *BookmarkBuildParams) WithHTTPClient(client *http.Client) *BookmarkBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bookmark build params
func (o *BookmarkBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the bookmark build params
func (o *BookmarkBuildParams) WithID(id string) *BookmarkBuildParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the bookmark build params
func (o *BookmarkBuildParams) SetID(id string) {
	o.ID = id
}

// WithOwner adds the owner to the bookmark build params
func (o *BookmarkBuildParams) WithOwner(owner string) *BookmarkBuildParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the bookmark build params
func (o *BookmarkBuildParams) SetOwner(owner string) {
	o.Owner = owner
}

// WithProject adds the project to the bookmark build params
func (o *BookmarkBuildParams) WithProject(project string) *BookmarkBuildParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the bookmark build params
func (o *BookmarkBuildParams) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *BookmarkBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
