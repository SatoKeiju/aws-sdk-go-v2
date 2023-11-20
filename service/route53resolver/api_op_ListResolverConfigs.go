// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the Resolver configurations that you have defined. Route 53 Resolver
// uses the configurations to manage DNS resolution behavior for your VPCs.
func (c *Client) ListResolverConfigs(ctx context.Context, params *ListResolverConfigsInput, optFns ...func(*Options)) (*ListResolverConfigsOutput, error) {
	if params == nil {
		params = &ListResolverConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverConfigs", params, optFns, c.addOperationListResolverConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverConfigsInput struct {

	// The maximum number of Resolver configurations that you want to return in the
	// response to a ListResolverConfigs request. If you don't specify a value for
	// MaxResults , up to 100 Resolver configurations are returned.
	MaxResults *int32

	// (Optional) If the current Amazon Web Services account has more than MaxResults
	// Resolver configurations, use NextToken to get the second and subsequent pages
	// of results. For the first ListResolverConfigs request, omit this value. For the
	// second and subsequent requests, get the value of NextToken from the previous
	// response and specify that value for NextToken in the request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResolverConfigsOutput struct {

	// If a response includes the last of the Resolver configurations that are
	// associated with the current Amazon Web Services account, NextToken doesn't
	// appear in the response. If a response doesn't include the last of the
	// configurations, you can get more configurations by submitting another
	// ListResolverConfigs request. Get the value of NextToken that Amazon Route 53
	// returned in the previous response and include it in NextToken in the next
	// request.
	NextToken *string

	// An array that contains one ResolverConfigs element for each Resolver
	// configuration that is associated with the current Amazon Web Services account.
	ResolverConfigs []types.ResolverConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResolverConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResolverConfigs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverConfigs(options.Region), middleware.Before); err != nil {
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

// ListResolverConfigsAPIClient is a client that implements the
// ListResolverConfigs operation.
type ListResolverConfigsAPIClient interface {
	ListResolverConfigs(context.Context, *ListResolverConfigsInput, ...func(*Options)) (*ListResolverConfigsOutput, error)
}

var _ ListResolverConfigsAPIClient = (*Client)(nil)

// ListResolverConfigsPaginatorOptions is the paginator options for
// ListResolverConfigs
type ListResolverConfigsPaginatorOptions struct {
	// The maximum number of Resolver configurations that you want to return in the
	// response to a ListResolverConfigs request. If you don't specify a value for
	// MaxResults , up to 100 Resolver configurations are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResolverConfigsPaginator is a paginator for ListResolverConfigs
type ListResolverConfigsPaginator struct {
	options   ListResolverConfigsPaginatorOptions
	client    ListResolverConfigsAPIClient
	params    *ListResolverConfigsInput
	nextToken *string
	firstPage bool
}

// NewListResolverConfigsPaginator returns a new ListResolverConfigsPaginator
func NewListResolverConfigsPaginator(client ListResolverConfigsAPIClient, params *ListResolverConfigsInput, optFns ...func(*ListResolverConfigsPaginatorOptions)) *ListResolverConfigsPaginator {
	if params == nil {
		params = &ListResolverConfigsInput{}
	}

	options := ListResolverConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResolverConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResolverConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResolverConfigs page.
func (p *ListResolverConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResolverConfigsOutput, error) {
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

	result, err := p.client.ListResolverConfigs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResolverConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResolverConfigs",
	}
}
