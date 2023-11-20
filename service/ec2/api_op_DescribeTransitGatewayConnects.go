// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes one or more Connect attachments.
func (c *Client) DescribeTransitGatewayConnects(ctx context.Context, params *DescribeTransitGatewayConnectsInput, optFns ...func(*Options)) (*DescribeTransitGatewayConnectsOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayConnectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayConnects", params, optFns, c.addOperationDescribeTransitGatewayConnectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayConnectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayConnectsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//   - options.protocol - The tunnel protocol ( gre ).
	//   - state - The state of the attachment ( initiating | initiatingRequest |
	//   pendingAcceptance | rollingBack | pending | available | modifying | deleting |
	//   deleted | failed | rejected | rejecting | failing ).
	//   - transit-gateway-attachment-id - The ID of the Connect attachment.
	//   - transit-gateway-id - The ID of the transit gateway.
	//   - transport-transit-gateway-attachment-id - The ID of the transit gateway
	//   attachment from which the Connect attachment was created.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the attachments.
	TransitGatewayAttachmentIds []string

	noSmithyDocumentSerde
}

type DescribeTransitGatewayConnectsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the Connect attachments.
	TransitGatewayConnects []types.TransitGatewayConnect

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTransitGatewayConnectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayConnects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayConnects{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTransitGatewayConnects"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayConnects(options.Region), middleware.Before); err != nil {
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

// DescribeTransitGatewayConnectsAPIClient is a client that implements the
// DescribeTransitGatewayConnects operation.
type DescribeTransitGatewayConnectsAPIClient interface {
	DescribeTransitGatewayConnects(context.Context, *DescribeTransitGatewayConnectsInput, ...func(*Options)) (*DescribeTransitGatewayConnectsOutput, error)
}

var _ DescribeTransitGatewayConnectsAPIClient = (*Client)(nil)

// DescribeTransitGatewayConnectsPaginatorOptions is the paginator options for
// DescribeTransitGatewayConnects
type DescribeTransitGatewayConnectsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTransitGatewayConnectsPaginator is a paginator for
// DescribeTransitGatewayConnects
type DescribeTransitGatewayConnectsPaginator struct {
	options   DescribeTransitGatewayConnectsPaginatorOptions
	client    DescribeTransitGatewayConnectsAPIClient
	params    *DescribeTransitGatewayConnectsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTransitGatewayConnectsPaginator returns a new
// DescribeTransitGatewayConnectsPaginator
func NewDescribeTransitGatewayConnectsPaginator(client DescribeTransitGatewayConnectsAPIClient, params *DescribeTransitGatewayConnectsInput, optFns ...func(*DescribeTransitGatewayConnectsPaginatorOptions)) *DescribeTransitGatewayConnectsPaginator {
	if params == nil {
		params = &DescribeTransitGatewayConnectsInput{}
	}

	options := DescribeTransitGatewayConnectsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTransitGatewayConnectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTransitGatewayConnectsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTransitGatewayConnects page.
func (p *DescribeTransitGatewayConnectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTransitGatewayConnectsOutput, error) {
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

	result, err := p.client.DescribeTransitGatewayConnects(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTransitGatewayConnects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTransitGatewayConnects",
	}
}
