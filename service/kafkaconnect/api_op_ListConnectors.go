// Code generated by smithy-go-codegen DO NOT EDIT.

package kafkaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafkaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the connectors in this account and Region. The list is
// limited to connectors whose name starts with the specified prefix. The response
// also includes a description of each of the listed connectors.
func (c *Client) ListConnectors(ctx context.Context, params *ListConnectorsInput, optFns ...func(*Options)) (*ListConnectorsOutput, error) {
	if params == nil {
		params = &ListConnectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConnectors", params, optFns, c.addOperationListConnectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConnectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectorsInput struct {

	// The name prefix that you want to use to search for and list connectors.
	ConnectorNamePrefix *string

	// The maximum number of connectors to list in one response.
	MaxResults int32

	// If the response of a ListConnectors operation is truncated, it will include a
	// NextToken. Send this NextToken in a subsequent request to continue listing from
	// where the previous operation left off.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConnectorsOutput struct {

	// An array of connector descriptions.
	Connectors []types.ConnectorSummary

	// If the response of a ListConnectors operation is truncated, it will include a
	// NextToken. Send this NextToken in a subsequent request to continue listing from
	// where it left off.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConnectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListConnectors"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConnectors(options.Region), middleware.Before); err != nil {
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

// ListConnectorsAPIClient is a client that implements the ListConnectors
// operation.
type ListConnectorsAPIClient interface {
	ListConnectors(context.Context, *ListConnectorsInput, ...func(*Options)) (*ListConnectorsOutput, error)
}

var _ ListConnectorsAPIClient = (*Client)(nil)

// ListConnectorsPaginatorOptions is the paginator options for ListConnectors
type ListConnectorsPaginatorOptions struct {
	// The maximum number of connectors to list in one response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConnectorsPaginator is a paginator for ListConnectors
type ListConnectorsPaginator struct {
	options   ListConnectorsPaginatorOptions
	client    ListConnectorsAPIClient
	params    *ListConnectorsInput
	nextToken *string
	firstPage bool
}

// NewListConnectorsPaginator returns a new ListConnectorsPaginator
func NewListConnectorsPaginator(client ListConnectorsAPIClient, params *ListConnectorsInput, optFns ...func(*ListConnectorsPaginatorOptions)) *ListConnectorsPaginator {
	if params == nil {
		params = &ListConnectorsInput{}
	}

	options := ListConnectorsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConnectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConnectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConnectors page.
func (p *ListConnectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConnectorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListConnectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConnectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListConnectors",
	}
}
