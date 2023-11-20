// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all long-term pricing types.
func (c *Client) ListLongTermPricing(ctx context.Context, params *ListLongTermPricingInput, optFns ...func(*Options)) (*ListLongTermPricingOutput, error) {
	if params == nil {
		params = &ListLongTermPricingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLongTermPricing", params, optFns, c.addOperationListLongTermPricingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLongTermPricingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLongTermPricingInput struct {

	// The maximum number of ListLongTermPricing objects to return.
	MaxResults *int32

	// Because HTTP requests are stateless, this is the starting point for your next
	// list of ListLongTermPricing to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLongTermPricingOutput struct {

	// Each LongTermPricingEntry object contains a status, ID, and other information
	// about the LongTermPricing type.
	LongTermPricingEntries []types.LongTermPricingListEntry

	// Because HTTP requests are stateless, this is the starting point for your next
	// list of returned ListLongTermPricing list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLongTermPricingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListLongTermPricing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLongTermPricing{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListLongTermPricing"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLongTermPricing(options.Region), middleware.Before); err != nil {
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

// ListLongTermPricingAPIClient is a client that implements the
// ListLongTermPricing operation.
type ListLongTermPricingAPIClient interface {
	ListLongTermPricing(context.Context, *ListLongTermPricingInput, ...func(*Options)) (*ListLongTermPricingOutput, error)
}

var _ ListLongTermPricingAPIClient = (*Client)(nil)

// ListLongTermPricingPaginatorOptions is the paginator options for
// ListLongTermPricing
type ListLongTermPricingPaginatorOptions struct {
	// The maximum number of ListLongTermPricing objects to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLongTermPricingPaginator is a paginator for ListLongTermPricing
type ListLongTermPricingPaginator struct {
	options   ListLongTermPricingPaginatorOptions
	client    ListLongTermPricingAPIClient
	params    *ListLongTermPricingInput
	nextToken *string
	firstPage bool
}

// NewListLongTermPricingPaginator returns a new ListLongTermPricingPaginator
func NewListLongTermPricingPaginator(client ListLongTermPricingAPIClient, params *ListLongTermPricingInput, optFns ...func(*ListLongTermPricingPaginatorOptions)) *ListLongTermPricingPaginator {
	if params == nil {
		params = &ListLongTermPricingInput{}
	}

	options := ListLongTermPricingPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLongTermPricingPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLongTermPricingPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLongTermPricing page.
func (p *ListLongTermPricingPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLongTermPricingOutput, error) {
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

	result, err := p.client.ListLongTermPricing(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLongTermPricing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListLongTermPricing",
	}
}
