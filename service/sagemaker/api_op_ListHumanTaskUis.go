// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the human task user interfaces in your account.
func (c *Client) ListHumanTaskUis(ctx context.Context, params *ListHumanTaskUisInput, optFns ...func(*Options)) (*ListHumanTaskUisOutput, error) {
	if params == nil {
		params = &ListHumanTaskUisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHumanTaskUis", params, optFns, c.addOperationListHumanTaskUisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHumanTaskUisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHumanTaskUisInput struct {

	// A filter that returns only human task user interfaces with a creation time
	// greater than or equal to the specified timestamp.
	CreationTimeAfter *time.Time

	// A filter that returns only human task user interfaces that were created before
	// the specified timestamp.
	CreationTimeBefore *time.Time

	// The total number of items to return. If the total number of available items is
	// more than the value specified in MaxResults , then a NextToken will be provided
	// in the output that you can use to resume pagination.
	MaxResults *int32

	// A token to resume pagination.
	NextToken *string

	// An optional value that specifies whether you want the results sorted in
	// Ascending or Descending order.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListHumanTaskUisOutput struct {

	// An array of objects describing the human task user interfaces.
	//
	// This member is required.
	HumanTaskUiSummaries []types.HumanTaskUiSummary

	// A token to resume pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHumanTaskUisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListHumanTaskUis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListHumanTaskUis{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListHumanTaskUis"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHumanTaskUis(options.Region), middleware.Before); err != nil {
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

// ListHumanTaskUisAPIClient is a client that implements the ListHumanTaskUis
// operation.
type ListHumanTaskUisAPIClient interface {
	ListHumanTaskUis(context.Context, *ListHumanTaskUisInput, ...func(*Options)) (*ListHumanTaskUisOutput, error)
}

var _ ListHumanTaskUisAPIClient = (*Client)(nil)

// ListHumanTaskUisPaginatorOptions is the paginator options for ListHumanTaskUis
type ListHumanTaskUisPaginatorOptions struct {
	// The total number of items to return. If the total number of available items is
	// more than the value specified in MaxResults , then a NextToken will be provided
	// in the output that you can use to resume pagination.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHumanTaskUisPaginator is a paginator for ListHumanTaskUis
type ListHumanTaskUisPaginator struct {
	options   ListHumanTaskUisPaginatorOptions
	client    ListHumanTaskUisAPIClient
	params    *ListHumanTaskUisInput
	nextToken *string
	firstPage bool
}

// NewListHumanTaskUisPaginator returns a new ListHumanTaskUisPaginator
func NewListHumanTaskUisPaginator(client ListHumanTaskUisAPIClient, params *ListHumanTaskUisInput, optFns ...func(*ListHumanTaskUisPaginatorOptions)) *ListHumanTaskUisPaginator {
	if params == nil {
		params = &ListHumanTaskUisInput{}
	}

	options := ListHumanTaskUisPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHumanTaskUisPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHumanTaskUisPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHumanTaskUis page.
func (p *ListHumanTaskUisPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHumanTaskUisOutput, error) {
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

	result, err := p.client.ListHumanTaskUis(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHumanTaskUis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListHumanTaskUis",
	}
}
