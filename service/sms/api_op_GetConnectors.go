// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the connectors registered with the Server Migration Service.
func (c *Client) GetConnectors(ctx context.Context, params *GetConnectorsInput, optFns ...func(*Options)) (*GetConnectorsOutput, error) {
	if params == nil {
		params = &GetConnectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConnectors", params, optFns, c.addOperationGetConnectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConnectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConnectorsInput struct {

	// The maximum number of results to return in a single call. The default value is
	// 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetConnectorsOutput struct {

	// Information about the registered connectors.
	ConnectorList []types.Connector

	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConnectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetConnectors"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConnectors(options.Region), middleware.Before); err != nil {
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

// GetConnectorsAPIClient is a client that implements the GetConnectors operation.
type GetConnectorsAPIClient interface {
	GetConnectors(context.Context, *GetConnectorsInput, ...func(*Options)) (*GetConnectorsOutput, error)
}

var _ GetConnectorsAPIClient = (*Client)(nil)

// GetConnectorsPaginatorOptions is the paginator options for GetConnectors
type GetConnectorsPaginatorOptions struct {
	// The maximum number of results to return in a single call. The default value is
	// 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetConnectorsPaginator is a paginator for GetConnectors
type GetConnectorsPaginator struct {
	options   GetConnectorsPaginatorOptions
	client    GetConnectorsAPIClient
	params    *GetConnectorsInput
	nextToken *string
	firstPage bool
}

// NewGetConnectorsPaginator returns a new GetConnectorsPaginator
func NewGetConnectorsPaginator(client GetConnectorsAPIClient, params *GetConnectorsInput, optFns ...func(*GetConnectorsPaginatorOptions)) *GetConnectorsPaginator {
	if params == nil {
		params = &GetConnectorsInput{}
	}

	options := GetConnectorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetConnectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetConnectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetConnectors page.
func (p *GetConnectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetConnectorsOutput, error) {
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

	result, err := p.client.GetConnectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetConnectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetConnectors",
	}
}
