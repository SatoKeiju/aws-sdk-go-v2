// Code generated by smithy-go-codegen DO NOT EDIT.

package entityresolution

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all the ProviderServices that are available in this Amazon
// Web Services Region.
func (c *Client) ListProviderServices(ctx context.Context, params *ListProviderServicesInput, optFns ...func(*Options)) (*ListProviderServicesOutput, error) {
	if params == nil {
		params = &ListProviderServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProviderServices", params, optFns, c.addOperationListProviderServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProviderServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProviderServicesInput struct {

	// The maximum number of objects returned per page.
	MaxResults *int32

	// The pagination token from the previous API call.
	NextToken *string

	// The name of the provider. This name is typically the company name.
	ProviderName *string

	noSmithyDocumentSerde
}

type ListProviderServicesOutput struct {

	// The pagination token from the previous API call.
	NextToken *string

	// A list of ProviderServices objects.
	ProviderServiceSummaries []types.ProviderServiceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProviderServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProviderServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProviderServices{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProviderServices"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProviderServices(options.Region), middleware.Before); err != nil {
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

// ListProviderServicesAPIClient is a client that implements the
// ListProviderServices operation.
type ListProviderServicesAPIClient interface {
	ListProviderServices(context.Context, *ListProviderServicesInput, ...func(*Options)) (*ListProviderServicesOutput, error)
}

var _ ListProviderServicesAPIClient = (*Client)(nil)

// ListProviderServicesPaginatorOptions is the paginator options for
// ListProviderServices
type ListProviderServicesPaginatorOptions struct {
	// The maximum number of objects returned per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProviderServicesPaginator is a paginator for ListProviderServices
type ListProviderServicesPaginator struct {
	options   ListProviderServicesPaginatorOptions
	client    ListProviderServicesAPIClient
	params    *ListProviderServicesInput
	nextToken *string
	firstPage bool
}

// NewListProviderServicesPaginator returns a new ListProviderServicesPaginator
func NewListProviderServicesPaginator(client ListProviderServicesAPIClient, params *ListProviderServicesInput, optFns ...func(*ListProviderServicesPaginatorOptions)) *ListProviderServicesPaginator {
	if params == nil {
		params = &ListProviderServicesInput{}
	}

	options := ListProviderServicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProviderServicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProviderServicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProviderServices page.
func (p *ListProviderServicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProviderServicesOutput, error) {
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

	result, err := p.client.ListProviderServices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProviderServices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProviderServices",
	}
}
