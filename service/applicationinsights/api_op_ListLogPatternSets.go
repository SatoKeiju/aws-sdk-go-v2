// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationinsights

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the log pattern sets in the specific application.
func (c *Client) ListLogPatternSets(ctx context.Context, params *ListLogPatternSetsInput, optFns ...func(*Options)) (*ListLogPatternSetsOutput, error) {
	if params == nil {
		params = &ListLogPatternSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLogPatternSets", params, optFns, c.addOperationListLogPatternSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLogPatternSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLogPatternSetsInput struct {

	// The name of the resource group.
	//
	// This member is required.
	ResourceGroupName *string

	// The AWS account ID for the resource group owner.
	AccountId *string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLogPatternSetsOutput struct {

	// The AWS account ID for the resource group owner.
	AccountId *string

	// The list of log pattern sets.
	LogPatternSets []string

	// The token used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The name of the resource group.
	ResourceGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLogPatternSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLogPatternSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLogPatternSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLogPatternSets"); err != nil {
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
	if err = addOpListLogPatternSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLogPatternSets(options.Region), middleware.Before); err != nil {
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

// ListLogPatternSetsAPIClient is a client that implements the ListLogPatternSets
// operation.
type ListLogPatternSetsAPIClient interface {
	ListLogPatternSets(context.Context, *ListLogPatternSetsInput, ...func(*Options)) (*ListLogPatternSetsOutput, error)
}

var _ ListLogPatternSetsAPIClient = (*Client)(nil)

// ListLogPatternSetsPaginatorOptions is the paginator options for
// ListLogPatternSets
type ListLogPatternSetsPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLogPatternSetsPaginator is a paginator for ListLogPatternSets
type ListLogPatternSetsPaginator struct {
	options   ListLogPatternSetsPaginatorOptions
	client    ListLogPatternSetsAPIClient
	params    *ListLogPatternSetsInput
	nextToken *string
	firstPage bool
}

// NewListLogPatternSetsPaginator returns a new ListLogPatternSetsPaginator
func NewListLogPatternSetsPaginator(client ListLogPatternSetsAPIClient, params *ListLogPatternSetsInput, optFns ...func(*ListLogPatternSetsPaginatorOptions)) *ListLogPatternSetsPaginator {
	if params == nil {
		params = &ListLogPatternSetsInput{}
	}

	options := ListLogPatternSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLogPatternSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLogPatternSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLogPatternSets page.
func (p *ListLogPatternSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLogPatternSetsOutput, error) {
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

	result, err := p.client.ListLogPatternSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLogPatternSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLogPatternSets",
	}
}
