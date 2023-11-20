// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a JSON array of up to twenty of your queues. This will return the
// queues themselves, not just a list of them. To retrieve the next twenty queues,
// use the nextToken string returned with the array.
func (c *Client) ListQueues(ctx context.Context, params *ListQueuesInput, optFns ...func(*Options)) (*ListQueuesOutput, error) {
	if params == nil {
		params = &ListQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueues", params, optFns, c.addOperationListQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueuesInput struct {

	// Optional. When you request a list of queues, you can choose to list them
	// alphabetically by NAME or chronologically by CREATION_DATE. If you don't
	// specify, the service will list them by creation date.
	ListBy types.QueueListBy

	// Optional. Number of queues, up to twenty, that will be returned at one time.
	MaxResults *int32

	// Use this string, provided with the response to a previous request, to request
	// the next batch of queues.
	NextToken *string

	// Optional. When you request lists of resources, you can specify whether they are
	// sorted in ASCENDING or DESCENDING order. Default varies by resource.
	Order types.Order

	noSmithyDocumentSerde
}

type ListQueuesOutput struct {

	// Use this string to request the next batch of queues.
	NextToken *string

	// List of queues.
	Queues []types.Queue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListQueues{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListQueues"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueues(options.Region), middleware.Before); err != nil {
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

// ListQueuesAPIClient is a client that implements the ListQueues operation.
type ListQueuesAPIClient interface {
	ListQueues(context.Context, *ListQueuesInput, ...func(*Options)) (*ListQueuesOutput, error)
}

var _ ListQueuesAPIClient = (*Client)(nil)

// ListQueuesPaginatorOptions is the paginator options for ListQueues
type ListQueuesPaginatorOptions struct {
	// Optional. Number of queues, up to twenty, that will be returned at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueuesPaginator is a paginator for ListQueues
type ListQueuesPaginator struct {
	options   ListQueuesPaginatorOptions
	client    ListQueuesAPIClient
	params    *ListQueuesInput
	nextToken *string
	firstPage bool
}

// NewListQueuesPaginator returns a new ListQueuesPaginator
func NewListQueuesPaginator(client ListQueuesAPIClient, params *ListQueuesInput, optFns ...func(*ListQueuesPaginatorOptions)) *ListQueuesPaginator {
	if params == nil {
		params = &ListQueuesInput{}
	}

	options := ListQueuesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueuesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueuesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueues page.
func (p *ListQueuesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueuesOutput, error) {
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

	result, err := p.client.ListQueues(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQueues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListQueues",
	}
}
