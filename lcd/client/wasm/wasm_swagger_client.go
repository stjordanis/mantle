// Code generated by go-swagger; DO NOT EDIT.

package wasm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new wasm API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wasm API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetWasmCodesCodeID(params *GetWasmCodesCodeIDParams) (*GetWasmCodesCodeIDOK, error)

	GetWasmContractsContractAddress(params *GetWasmContractsContractAddressParams) (*GetWasmContractsContractAddressOK, error)

	GetWasmContractsContractAddressStore(params *GetWasmContractsContractAddressStoreParams) (*GetWasmContractsContractAddressStoreOK, error)

	GetWasmContractsContractAddressStoreRaw(params *GetWasmContractsContractAddressStoreRawParams) (*GetWasmContractsContractAddressStoreRawOK, error)

	GetWasmParameters(params *GetWasmParametersParams) (*GetWasmParametersOK, error)

	PostWasmCodes(params *PostWasmCodesParams) (*PostWasmCodesOK, error)

	PostWasmCodesCodeID(params *PostWasmCodesCodeIDParams) (*PostWasmCodesCodeIDOK, error)

	PostWasmContractsContractAddress(params *PostWasmContractsContractAddressParams) (*PostWasmContractsContractAddressOK, error)

	PostWasmContractsContractAddressMigrate(params *PostWasmContractsContractAddressMigrateParams) (*PostWasmContractsContractAddressMigrateOK, error)

	PostWasmContractsContractAddressOwner(params *PostWasmContractsContractAddressOwnerParams) (*PostWasmContractsContractAddressOwnerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetWasmCodesCodeID gets code info of the code ID
*/
func (a *Client) GetWasmCodesCodeID(params *GetWasmCodesCodeIDParams) (*GetWasmCodesCodeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWasmCodesCodeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWasmCodesCodeID",
		Method:             "GET",
		PathPattern:        "/wasm/codes/{codeID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWasmCodesCodeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWasmCodesCodeIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWasmCodesCodeID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWasmContractsContractAddress gets contract info of the contract address
*/
func (a *Client) GetWasmContractsContractAddress(params *GetWasmContractsContractAddressParams) (*GetWasmContractsContractAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWasmContractsContractAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWasmContractsContractAddress",
		Method:             "GET",
		PathPattern:        "/wasm/contracts/{contractAddress}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWasmContractsContractAddressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWasmContractsContractAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWasmContractsContractAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWasmContractsContractAddressStore gets stored information with query msg
*/
func (a *Client) GetWasmContractsContractAddressStore(params *GetWasmContractsContractAddressStoreParams) (*GetWasmContractsContractAddressStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWasmContractsContractAddressStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWasmContractsContractAddressStore",
		Method:             "GET",
		PathPattern:        "/wasm/contracts/{contractAddress}/store",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWasmContractsContractAddressStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWasmContractsContractAddressStoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWasmContractsContractAddressStore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWasmContractsContractAddressStoreRaw gets stored information with store key
*/
func (a *Client) GetWasmContractsContractAddressStoreRaw(params *GetWasmContractsContractAddressStoreRawParams) (*GetWasmContractsContractAddressStoreRawOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWasmContractsContractAddressStoreRawParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWasmContractsContractAddressStoreRaw",
		Method:             "GET",
		PathPattern:        "/wasm/contracts/{contractAddress}/store/raw",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWasmContractsContractAddressStoreRawReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWasmContractsContractAddressStoreRawOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWasmContractsContractAddressStoreRaw: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetWasmParameters gets wasm module params
*/
func (a *Client) GetWasmParameters(params *GetWasmParametersParams) (*GetWasmParametersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWasmParametersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWasmParameters",
		Method:             "GET",
		PathPattern:        "/wasm/parameters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWasmParametersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWasmParametersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWasmParameters: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostWasmCodes generates wasm store code message
*/
func (a *Client) PostWasmCodes(params *PostWasmCodesParams) (*PostWasmCodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWasmCodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWasmCodes",
		Method:             "POST",
		PathPattern:        "/wasm/codes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWasmCodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWasmCodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWasmCodes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostWasmCodesCodeID instantiates wasm contract message
*/
func (a *Client) PostWasmCodesCodeID(params *PostWasmCodesCodeIDParams) (*PostWasmCodesCodeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWasmCodesCodeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWasmCodesCodeID",
		Method:             "POST",
		PathPattern:        "/wasm/codes/{codeID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWasmCodesCodeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWasmCodesCodeIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWasmCodesCodeID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostWasmContractsContractAddress executes wasm contract message
*/
func (a *Client) PostWasmContractsContractAddress(params *PostWasmContractsContractAddressParams) (*PostWasmContractsContractAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWasmContractsContractAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWasmContractsContractAddress",
		Method:             "POST",
		PathPattern:        "/wasm/contracts/{contractAddress}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWasmContractsContractAddressReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWasmContractsContractAddressOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWasmContractsContractAddress: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostWasmContractsContractAddressMigrate migrates wasm contract to new code base
*/
func (a *Client) PostWasmContractsContractAddressMigrate(params *PostWasmContractsContractAddressMigrateParams) (*PostWasmContractsContractAddressMigrateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWasmContractsContractAddressMigrateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWasmContractsContractAddressMigrate",
		Method:             "POST",
		PathPattern:        "/wasm/contracts/{contractAddress}/migrate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWasmContractsContractAddressMigrateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWasmContractsContractAddressMigrateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWasmContractsContractAddressMigrate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostWasmContractsContractAddressOwner updates wasm contract owner to new address
*/
func (a *Client) PostWasmContractsContractAddressOwner(params *PostWasmContractsContractAddressOwnerParams) (*PostWasmContractsContractAddressOwnerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWasmContractsContractAddressOwnerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWasmContractsContractAddressOwner",
		Method:             "POST",
		PathPattern:        "/wasm/contracts/{contractAddress}/owner",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWasmContractsContractAddressOwnerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWasmContractsContractAddressOwnerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWasmContractsContractAddressOwner: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}