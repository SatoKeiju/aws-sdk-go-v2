// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all usage limits within Amazon Redshift Serverless.
func (c *Client) ListUsageLimits(ctx context.Context, params *ListUsageLimitsInput, optFns ...func(*Options)) (*ListUsageLimitsOutput, error) {
	if params == nil {
		params = &ListUsageLimitsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUsageLimits", params, optFns, c.addOperationListUsageLimitsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUsageLimitsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUsageLimitsInput struct {

	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results. The default is 100.
	MaxResults *int32

	// If your initial ListUsageLimits operation returns a nextToken , you can include
	// the returned nextToken in following ListUsageLimits operations, which returns
	// results in the next page.
	NextToken *string

	// The Amazon Resource Name (ARN) associated with the resource whose usage limits
	// you want to list.
	ResourceArn *string

	// The Amazon Redshift Serverless feature whose limits you want to see.
	UsageType types.UsageLimitUsageType

	noSmithyDocumentSerde
}

type ListUsageLimitsOutput struct {

	// When nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again using
	// the returned token to retrieve the next page.
	NextToken *string

	// An array of returned usage limit objects.
	UsageLimits []types.UsageLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUsageLimitsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListUsageLimits{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUsageLimits{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListUsageLimits"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUsageLimits(options.Region), middleware.Before); err != nil {
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

// ListUsageLimitsAPIClient is a client that implements the ListUsageLimits
// operation.
type ListUsageLimitsAPIClient interface {
	ListUsageLimits(context.Context, *ListUsageLimitsInput, ...func(*Options)) (*ListUsageLimitsOutput, error)
}

var _ ListUsageLimitsAPIClient = (*Client)(nil)

// ListUsageLimitsPaginatorOptions is the paginator options for ListUsageLimits
type ListUsageLimitsPaginatorOptions struct {
	// An optional parameter that specifies the maximum number of results to return.
	// You can use nextToken to get the next page of results. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUsageLimitsPaginator is a paginator for ListUsageLimits
type ListUsageLimitsPaginator struct {
	options   ListUsageLimitsPaginatorOptions
	client    ListUsageLimitsAPIClient
	params    *ListUsageLimitsInput
	nextToken *string
	firstPage bool
}

// NewListUsageLimitsPaginator returns a new ListUsageLimitsPaginator
func NewListUsageLimitsPaginator(client ListUsageLimitsAPIClient, params *ListUsageLimitsInput, optFns ...func(*ListUsageLimitsPaginatorOptions)) *ListUsageLimitsPaginator {
	if params == nil {
		params = &ListUsageLimitsInput{}
	}

	options := ListUsageLimitsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUsageLimitsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUsageLimitsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUsageLimits page.
func (p *ListUsageLimitsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUsageLimitsOutput, error) {
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

	result, err := p.client.ListUsageLimits(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUsageLimits(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListUsageLimits",
	}
}
