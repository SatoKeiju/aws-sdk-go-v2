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

// Describes one or more Traffic Mirror sessions. By default, all Traffic Mirror
// sessions are described. Alternatively, you can filter the results.
func (c *Client) DescribeTrafficMirrorSessions(ctx context.Context, params *DescribeTrafficMirrorSessionsInput, optFns ...func(*Options)) (*DescribeTrafficMirrorSessionsOutput, error) {
	if params == nil {
		params = &DescribeTrafficMirrorSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTrafficMirrorSessions", params, optFns, c.addOperationDescribeTrafficMirrorSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTrafficMirrorSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrafficMirrorSessionsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// One or more filters. The possible values are:
	//   - description : The Traffic Mirror session description.
	//   - network-interface-id : The ID of the Traffic Mirror session network
	//   interface.
	//   - owner-id : The ID of the account that owns the Traffic Mirror session.
	//   - packet-length : The assigned number of packets to mirror.
	//   - session-number : The assigned session number.
	//   - traffic-mirror-filter-id : The ID of the Traffic Mirror filter.
	//   - traffic-mirror-session-id : The ID of the Traffic Mirror session.
	//   - traffic-mirror-target-id : The ID of the Traffic Mirror target.
	//   - virtual-network-id : The virtual network ID of the Traffic Mirror session.
	Filters []types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ID of the Traffic Mirror session.
	TrafficMirrorSessionIds []string

	noSmithyDocumentSerde
}

type DescribeTrafficMirrorSessionsOutput struct {

	// The token to use to retrieve the next page of results. The value is null when
	// there are no more results to return.
	NextToken *string

	// Describes one or more Traffic Mirror sessions. By default, all Traffic Mirror
	// sessions are described. Alternatively, you can filter the results.
	TrafficMirrorSessions []types.TrafficMirrorSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTrafficMirrorSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTrafficMirrorSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTrafficMirrorSessions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTrafficMirrorSessions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrafficMirrorSessions(options.Region), middleware.Before); err != nil {
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

// DescribeTrafficMirrorSessionsAPIClient is a client that implements the
// DescribeTrafficMirrorSessions operation.
type DescribeTrafficMirrorSessionsAPIClient interface {
	DescribeTrafficMirrorSessions(context.Context, *DescribeTrafficMirrorSessionsInput, ...func(*Options)) (*DescribeTrafficMirrorSessionsOutput, error)
}

var _ DescribeTrafficMirrorSessionsAPIClient = (*Client)(nil)

// DescribeTrafficMirrorSessionsPaginatorOptions is the paginator options for
// DescribeTrafficMirrorSessions
type DescribeTrafficMirrorSessionsPaginatorOptions struct {
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTrafficMirrorSessionsPaginator is a paginator for
// DescribeTrafficMirrorSessions
type DescribeTrafficMirrorSessionsPaginator struct {
	options   DescribeTrafficMirrorSessionsPaginatorOptions
	client    DescribeTrafficMirrorSessionsAPIClient
	params    *DescribeTrafficMirrorSessionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeTrafficMirrorSessionsPaginator returns a new
// DescribeTrafficMirrorSessionsPaginator
func NewDescribeTrafficMirrorSessionsPaginator(client DescribeTrafficMirrorSessionsAPIClient, params *DescribeTrafficMirrorSessionsInput, optFns ...func(*DescribeTrafficMirrorSessionsPaginatorOptions)) *DescribeTrafficMirrorSessionsPaginator {
	if params == nil {
		params = &DescribeTrafficMirrorSessionsInput{}
	}

	options := DescribeTrafficMirrorSessionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTrafficMirrorSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTrafficMirrorSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTrafficMirrorSessions page.
func (p *DescribeTrafficMirrorSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTrafficMirrorSessionsOutput, error) {
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

	result, err := p.client.DescribeTrafficMirrorSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTrafficMirrorSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTrafficMirrorSessions",
	}
}
