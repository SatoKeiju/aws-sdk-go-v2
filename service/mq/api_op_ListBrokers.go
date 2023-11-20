// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all brokers.
func (c *Client) ListBrokers(ctx context.Context, params *ListBrokersInput, optFns ...func(*Options)) (*ListBrokersOutput, error) {
	if params == nil {
		params = &ListBrokersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBrokers", params, optFns, c.addOperationListBrokersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBrokersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBrokersInput struct {

	// The maximum number of brokers that Amazon MQ can return per page (20 by
	// default). This value must be an integer from 5 to 100.
	MaxResults *int32

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBrokersOutput struct {

	// A list of information about all brokers.
	BrokerSummaries []types.BrokerSummary

	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBrokersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBrokers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBrokers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBrokers"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBrokers(options.Region), middleware.Before); err != nil {
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

// ListBrokersAPIClient is a client that implements the ListBrokers operation.
type ListBrokersAPIClient interface {
	ListBrokers(context.Context, *ListBrokersInput, ...func(*Options)) (*ListBrokersOutput, error)
}

var _ ListBrokersAPIClient = (*Client)(nil)

// ListBrokersPaginatorOptions is the paginator options for ListBrokers
type ListBrokersPaginatorOptions struct {
	// The maximum number of brokers that Amazon MQ can return per page (20 by
	// default). This value must be an integer from 5 to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBrokersPaginator is a paginator for ListBrokers
type ListBrokersPaginator struct {
	options   ListBrokersPaginatorOptions
	client    ListBrokersAPIClient
	params    *ListBrokersInput
	nextToken *string
	firstPage bool
}

// NewListBrokersPaginator returns a new ListBrokersPaginator
func NewListBrokersPaginator(client ListBrokersAPIClient, params *ListBrokersInput, optFns ...func(*ListBrokersPaginatorOptions)) *ListBrokersPaginator {
	if params == nil {
		params = &ListBrokersInput{}
	}

	options := ListBrokersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBrokersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBrokersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBrokers page.
func (p *ListBrokersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBrokersOutput, error) {
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

	result, err := p.client.ListBrokers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBrokers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBrokers",
	}
}
