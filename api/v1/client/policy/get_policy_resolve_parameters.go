// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// NewGetPolicyResolveParams creates a new GetPolicyResolveParams object
// with the default values initialized.
func NewGetPolicyResolveParams() *GetPolicyResolveParams {
	var ()
	return &GetPolicyResolveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyResolveParamsWithTimeout creates a new GetPolicyResolveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyResolveParamsWithTimeout(timeout time.Duration) *GetPolicyResolveParams {
	var ()
	return &GetPolicyResolveParams{

		timeout: timeout,
	}
}

// NewGetPolicyResolveParamsWithContext creates a new GetPolicyResolveParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyResolveParamsWithContext(ctx context.Context) *GetPolicyResolveParams {
	var ()
	return &GetPolicyResolveParams{

		Context: ctx,
	}
}

// NewGetPolicyResolveParamsWithHTTPClient creates a new GetPolicyResolveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyResolveParamsWithHTTPClient(client *http.Client) *GetPolicyResolveParams {
	var ()
	return &GetPolicyResolveParams{
		HTTPClient: client,
	}
}

/*GetPolicyResolveParams contains all the parameters to send to the API endpoint
for the get policy resolve operation typically these are written to a http.Request
*/
type GetPolicyResolveParams struct {

	/*IdentityContext
	  Context to provide policy evaluation on

	*/
	IdentityContext *models.IdentityContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy resolve params
func (o *GetPolicyResolveParams) WithTimeout(timeout time.Duration) *GetPolicyResolveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy resolve params
func (o *GetPolicyResolveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy resolve params
func (o *GetPolicyResolveParams) WithContext(ctx context.Context) *GetPolicyResolveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy resolve params
func (o *GetPolicyResolveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy resolve params
func (o *GetPolicyResolveParams) WithHTTPClient(client *http.Client) *GetPolicyResolveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy resolve params
func (o *GetPolicyResolveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentityContext adds the identityContext to the get policy resolve params
func (o *GetPolicyResolveParams) WithIdentityContext(identityContext *models.IdentityContext) *GetPolicyResolveParams {
	o.SetIdentityContext(identityContext)
	return o
}

// SetIdentityContext adds the identityContext to the get policy resolve params
func (o *GetPolicyResolveParams) SetIdentityContext(identityContext *models.IdentityContext) {
	o.IdentityContext = identityContext
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyResolveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IdentityContext != nil {
		if err := r.SetBodyParam(o.IdentityContext); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
