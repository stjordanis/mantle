// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetGovProposalsProposalIDProposerParams creates a new GetGovProposalsProposalIDProposerParams object
// with the default values initialized.
func NewGetGovProposalsProposalIDProposerParams() *GetGovProposalsProposalIDProposerParams {
	var ()
	return &GetGovProposalsProposalIDProposerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGovProposalsProposalIDProposerParamsWithTimeout creates a new GetGovProposalsProposalIDProposerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGovProposalsProposalIDProposerParamsWithTimeout(timeout time.Duration) *GetGovProposalsProposalIDProposerParams {
	var ()
	return &GetGovProposalsProposalIDProposerParams{

		timeout: timeout,
	}
}

// NewGetGovProposalsProposalIDProposerParamsWithContext creates a new GetGovProposalsProposalIDProposerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGovProposalsProposalIDProposerParamsWithContext(ctx context.Context) *GetGovProposalsProposalIDProposerParams {
	var ()
	return &GetGovProposalsProposalIDProposerParams{

		Context: ctx,
	}
}

// NewGetGovProposalsProposalIDProposerParamsWithHTTPClient creates a new GetGovProposalsProposalIDProposerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGovProposalsProposalIDProposerParamsWithHTTPClient(client *http.Client) *GetGovProposalsProposalIDProposerParams {
	var ()
	return &GetGovProposalsProposalIDProposerParams{
		HTTPClient: client,
	}
}

/*GetGovProposalsProposalIDProposerParams contains all the parameters to send to the API endpoint
for the get gov proposals proposal ID proposer operation typically these are written to a http.Request
*/
type GetGovProposalsProposalIDProposerParams struct {

	/*ProposalID*/
	ProposalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gov proposals proposal ID proposer params
func (o *GetGovProposalsProposalIDProposerParams) WithTimeout(timeout time.Duration) *GetGovProposalsProposalIDProposerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gov proposals proposal ID proposer params
func (o *GetGovProposalsProposalIDProposerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gov proposals proposal ID proposer params
func (o *GetGovProposalsProposalIDProposerParams) WithContext(ctx context.Context) *GetGovProposalsProposalIDProposerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gov proposals proposal ID proposer params
func (o *GetGovProposalsProposalIDProposerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gov proposals proposal ID proposer params
func (o *GetGovProposalsProposalIDProposerParams) WithHTTPClient(client *http.Client) *GetGovProposalsProposalIDProposerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gov proposals proposal ID proposer params
func (o *GetGovProposalsProposalIDProposerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProposalID adds the proposalID to the get gov proposals proposal ID proposer params
func (o *GetGovProposalsProposalIDProposerParams) WithProposalID(proposalID string) *GetGovProposalsProposalIDProposerParams {
	o.SetProposalID(proposalID)
	return o
}

// SetProposalID adds the proposalId to the get gov proposals proposal ID proposer params
func (o *GetGovProposalsProposalIDProposerParams) SetProposalID(proposalID string) {
	o.ProposalID = proposalID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGovProposalsProposalIDProposerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param proposalId
	if err := r.SetPathParam("proposalId", o.ProposalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}