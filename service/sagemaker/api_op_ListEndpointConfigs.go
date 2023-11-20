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

// Lists endpoint configurations.
func (c *Client) ListEndpointConfigs(ctx context.Context, params *ListEndpointConfigsInput, optFns ...func(*Options)) (*ListEndpointConfigsOutput, error) {
	if params == nil {
		params = &ListEndpointConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEndpointConfigs", params, optFns, c.addOperationListEndpointConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEndpointConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointConfigsInput struct {

	// A filter that returns only endpoint configurations with a creation time greater
	// than or equal to the specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only endpoint configurations created before the specified
	// time (timestamp).
	CreationTimeBefore *time.Time

	// The maximum number of training jobs to return in the response.
	MaxResults *int32

	// A string in the endpoint configuration name. This filter returns only endpoint
	// configurations whose name contains the specified string.
	NameContains *string

	// If the result of the previous ListEndpointConfig request was truncated, the
	// response includes a NextToken . To retrieve the next set of endpoint
	// configurations, use the token in the next request.
	NextToken *string

	// The field to sort results by. The default is CreationTime .
	SortBy types.EndpointConfigSortKey

	// The sort order for results. The default is Descending .
	SortOrder types.OrderKey

	noSmithyDocumentSerde
}

type ListEndpointConfigsOutput struct {

	// An array of endpoint configurations.
	//
	// This member is required.
	EndpointConfigs []types.EndpointConfigSummary

	// If the response is truncated, SageMaker returns this token. To retrieve the
	// next set of endpoint configurations, use it in the subsequent request
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEndpointConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpointConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpointConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEndpointConfigs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpointConfigs(options.Region), middleware.Before); err != nil {
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

// ListEndpointConfigsAPIClient is a client that implements the
// ListEndpointConfigs operation.
type ListEndpointConfigsAPIClient interface {
	ListEndpointConfigs(context.Context, *ListEndpointConfigsInput, ...func(*Options)) (*ListEndpointConfigsOutput, error)
}

var _ ListEndpointConfigsAPIClient = (*Client)(nil)

// ListEndpointConfigsPaginatorOptions is the paginator options for
// ListEndpointConfigs
type ListEndpointConfigsPaginatorOptions struct {
	// The maximum number of training jobs to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEndpointConfigsPaginator is a paginator for ListEndpointConfigs
type ListEndpointConfigsPaginator struct {
	options   ListEndpointConfigsPaginatorOptions
	client    ListEndpointConfigsAPIClient
	params    *ListEndpointConfigsInput
	nextToken *string
	firstPage bool
}

// NewListEndpointConfigsPaginator returns a new ListEndpointConfigsPaginator
func NewListEndpointConfigsPaginator(client ListEndpointConfigsAPIClient, params *ListEndpointConfigsInput, optFns ...func(*ListEndpointConfigsPaginatorOptions)) *ListEndpointConfigsPaginator {
	if params == nil {
		params = &ListEndpointConfigsInput{}
	}

	options := ListEndpointConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEndpointConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEndpointConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEndpointConfigs page.
func (p *ListEndpointConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEndpointConfigsOutput, error) {
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

	result, err := p.client.ListEndpointConfigs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEndpointConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEndpointConfigs",
	}
}
