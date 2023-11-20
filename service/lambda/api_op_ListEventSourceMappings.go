// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists event source mappings. Specify an EventSourceArn to show only event
// source mappings for a single event source.
func (c *Client) ListEventSourceMappings(ctx context.Context, params *ListEventSourceMappingsInput, optFns ...func(*Options)) (*ListEventSourceMappingsOutput, error) {
	if params == nil {
		params = &ListEventSourceMappingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventSourceMappings", params, optFns, c.addOperationListEventSourceMappingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventSourceMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventSourceMappingsInput struct {

	// The Amazon Resource Name (ARN) of the event source.
	//   - Amazon Kinesis – The ARN of the data stream or a stream consumer.
	//   - Amazon DynamoDB Streams – The ARN of the stream.
	//   - Amazon Simple Queue Service – The ARN of the queue.
	//   - Amazon Managed Streaming for Apache Kafka – The ARN of the cluster.
	//   - Amazon MQ – The ARN of the broker.
	//   - Amazon DocumentDB – The ARN of the DocumentDB change stream.
	EventSourceArn *string

	// The name of the Lambda function. Name formats
	//   - Function name – MyFunction .
	//   - Function ARN – arn:aws:lambda:us-west-2:123456789012:function:MyFunction .
	//   - Version or Alias ARN –
	//   arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD .
	//   - Partial ARN – 123456789012:function:MyFunction .
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it's limited to 64 characters in length.
	FunctionName *string

	// A pagination token returned by a previous call.
	Marker *string

	// The maximum number of event source mappings to return. Note that
	// ListEventSourceMappings returns a maximum of 100 items in each response, even if
	// you set the number higher.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListEventSourceMappingsOutput struct {

	// A list of event source mappings.
	EventSourceMappings []types.EventSourceMappingConfiguration

	// A pagination token that's returned when the response doesn't contain all event
	// source mappings.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventSourceMappingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEventSourceMappings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventSourceMappings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEventSourceMappings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventSourceMappings(options.Region), middleware.Before); err != nil {
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

// ListEventSourceMappingsAPIClient is a client that implements the
// ListEventSourceMappings operation.
type ListEventSourceMappingsAPIClient interface {
	ListEventSourceMappings(context.Context, *ListEventSourceMappingsInput, ...func(*Options)) (*ListEventSourceMappingsOutput, error)
}

var _ ListEventSourceMappingsAPIClient = (*Client)(nil)

// ListEventSourceMappingsPaginatorOptions is the paginator options for
// ListEventSourceMappings
type ListEventSourceMappingsPaginatorOptions struct {
	// The maximum number of event source mappings to return. Note that
	// ListEventSourceMappings returns a maximum of 100 items in each response, even if
	// you set the number higher.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventSourceMappingsPaginator is a paginator for ListEventSourceMappings
type ListEventSourceMappingsPaginator struct {
	options   ListEventSourceMappingsPaginatorOptions
	client    ListEventSourceMappingsAPIClient
	params    *ListEventSourceMappingsInput
	nextToken *string
	firstPage bool
}

// NewListEventSourceMappingsPaginator returns a new
// ListEventSourceMappingsPaginator
func NewListEventSourceMappingsPaginator(client ListEventSourceMappingsAPIClient, params *ListEventSourceMappingsInput, optFns ...func(*ListEventSourceMappingsPaginatorOptions)) *ListEventSourceMappingsPaginator {
	if params == nil {
		params = &ListEventSourceMappingsInput{}
	}

	options := ListEventSourceMappingsPaginatorOptions{}
	if params.MaxItems != nil {
		options.Limit = *params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventSourceMappingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventSourceMappingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventSourceMappings page.
func (p *ListEventSourceMappingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventSourceMappingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxItems = limit

	result, err := p.client.ListEventSourceMappings(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListEventSourceMappings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEventSourceMappings",
	}
}
