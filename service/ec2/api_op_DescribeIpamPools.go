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

// Get information about your IPAM pools.
func (c *Client) DescribeIpamPools(ctx context.Context, params *DescribeIpamPoolsInput, optFns ...func(*Options)) (*DescribeIpamPoolsOutput, error) {
	if params == nil {
		params = &DescribeIpamPoolsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeIpamPools", params, optFns, c.addOperationDescribeIpamPoolsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeIpamPoolsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeIpamPoolsInput struct {

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters for the request. For more information about filtering, see
	// Filtering CLI output (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-filter.html)
	// .
	Filters []types.Filter

	// The IDs of the IPAM pools you would like information on.
	IpamPoolIds []string

	// The maximum number of results to return in the request.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeIpamPoolsOutput struct {

	// Information about the IPAM pools.
	IpamPools []types.IpamPool

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeIpamPoolsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeIpamPools{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeIpamPools{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeIpamPools"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeIpamPools(options.Region), middleware.Before); err != nil {
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

// DescribeIpamPoolsAPIClient is a client that implements the DescribeIpamPools
// operation.
type DescribeIpamPoolsAPIClient interface {
	DescribeIpamPools(context.Context, *DescribeIpamPoolsInput, ...func(*Options)) (*DescribeIpamPoolsOutput, error)
}

var _ DescribeIpamPoolsAPIClient = (*Client)(nil)

// DescribeIpamPoolsPaginatorOptions is the paginator options for DescribeIpamPools
type DescribeIpamPoolsPaginatorOptions struct {
	// The maximum number of results to return in the request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeIpamPoolsPaginator is a paginator for DescribeIpamPools
type DescribeIpamPoolsPaginator struct {
	options   DescribeIpamPoolsPaginatorOptions
	client    DescribeIpamPoolsAPIClient
	params    *DescribeIpamPoolsInput
	nextToken *string
	firstPage bool
}

// NewDescribeIpamPoolsPaginator returns a new DescribeIpamPoolsPaginator
func NewDescribeIpamPoolsPaginator(client DescribeIpamPoolsAPIClient, params *DescribeIpamPoolsInput, optFns ...func(*DescribeIpamPoolsPaginatorOptions)) *DescribeIpamPoolsPaginator {
	if params == nil {
		params = &DescribeIpamPoolsInput{}
	}

	options := DescribeIpamPoolsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeIpamPoolsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeIpamPoolsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeIpamPools page.
func (p *DescribeIpamPoolsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeIpamPoolsOutput, error) {
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

	result, err := p.client.DescribeIpamPools(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeIpamPools(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeIpamPools",
	}
}
