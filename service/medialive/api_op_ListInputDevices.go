// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List input devices
func (c *Client) ListInputDevices(ctx context.Context, params *ListInputDevicesInput, optFns ...func(*Options)) (*ListInputDevicesOutput, error) {
	if params == nil {
		params = &ListInputDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInputDevices", params, optFns, c.addOperationListInputDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInputDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListInputDevicesRequest
type ListInputDevicesInput struct {

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// Placeholder documentation for __string
	NextToken *string

	noSmithyDocumentSerde
}

// Placeholder documentation for ListInputDevicesResponse
type ListInputDevicesOutput struct {

	// The list of input devices.
	InputDevices []types.InputDeviceSummary

	// A token to get additional list results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInputDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInputDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInputDevices{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInputDevices"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addV4DetectSkewMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInputDevices(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListInputDevicesAPIClient is a client that implements the ListInputDevices
// operation.
type ListInputDevicesAPIClient interface {
	ListInputDevices(context.Context, *ListInputDevicesInput, ...func(*Options)) (*ListInputDevicesOutput, error)
}

var _ ListInputDevicesAPIClient = (*Client)(nil)

// ListInputDevicesPaginatorOptions is the paginator options for ListInputDevices
type ListInputDevicesPaginatorOptions struct {
	// Placeholder documentation for MaxResults
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInputDevicesPaginator is a paginator for ListInputDevices
type ListInputDevicesPaginator struct {
	options   ListInputDevicesPaginatorOptions
	client    ListInputDevicesAPIClient
	params    *ListInputDevicesInput
	nextToken *string
	firstPage bool
}

// NewListInputDevicesPaginator returns a new ListInputDevicesPaginator
func NewListInputDevicesPaginator(client ListInputDevicesAPIClient, params *ListInputDevicesInput, optFns ...func(*ListInputDevicesPaginatorOptions)) *ListInputDevicesPaginator {
	if params == nil {
		params = &ListInputDevicesInput{}
	}

	options := ListInputDevicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInputDevicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInputDevicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInputDevices page.
func (p *ListInputDevicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInputDevicesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListInputDevices(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListInputDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInputDevices",
	}
}
