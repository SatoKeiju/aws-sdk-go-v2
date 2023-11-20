// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List all version of an adapter that meet the specified filtration criteria.
func (c *Client) ListAdapterVersions(ctx context.Context, params *ListAdapterVersionsInput, optFns ...func(*Options)) (*ListAdapterVersionsOutput, error) {
	if params == nil {
		params = &ListAdapterVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAdapterVersions", params, optFns, c.addOperationListAdapterVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAdapterVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAdapterVersionsInput struct {

	// A string containing a unique ID for the adapter to match for when listing
	// adapter versions.
	AdapterId *string

	// Specifies the lower bound for the ListAdapterVersions operation. Ensures
	// ListAdapterVersions returns only adapter versions created after the specified
	// creation time.
	AfterCreationTime *time.Time

	// Specifies the upper bound for the ListAdapterVersions operation. Ensures
	// ListAdapterVersions returns only adapter versions created after the specified
	// creation time.
	BeforeCreationTime *time.Time

	// The maximum number of results to return when listing adapter versions.
	MaxResults *int32

	// Identifies the next page of results to return when listing adapter versions.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAdapterVersionsOutput struct {

	// Adapter versions that match the filtering criteria specified when calling
	// ListAdapters.
	AdapterVersions []types.AdapterVersionOverview

	// Identifies the next page of results to return when listing adapter versions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAdapterVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAdapterVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAdapterVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAdapterVersions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAdapterVersions(options.Region), middleware.Before); err != nil {
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

// ListAdapterVersionsAPIClient is a client that implements the
// ListAdapterVersions operation.
type ListAdapterVersionsAPIClient interface {
	ListAdapterVersions(context.Context, *ListAdapterVersionsInput, ...func(*Options)) (*ListAdapterVersionsOutput, error)
}

var _ ListAdapterVersionsAPIClient = (*Client)(nil)

// ListAdapterVersionsPaginatorOptions is the paginator options for
// ListAdapterVersions
type ListAdapterVersionsPaginatorOptions struct {
	// The maximum number of results to return when listing adapter versions.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAdapterVersionsPaginator is a paginator for ListAdapterVersions
type ListAdapterVersionsPaginator struct {
	options   ListAdapterVersionsPaginatorOptions
	client    ListAdapterVersionsAPIClient
	params    *ListAdapterVersionsInput
	nextToken *string
	firstPage bool
}

// NewListAdapterVersionsPaginator returns a new ListAdapterVersionsPaginator
func NewListAdapterVersionsPaginator(client ListAdapterVersionsAPIClient, params *ListAdapterVersionsInput, optFns ...func(*ListAdapterVersionsPaginatorOptions)) *ListAdapterVersionsPaginator {
	if params == nil {
		params = &ListAdapterVersionsInput{}
	}

	options := ListAdapterVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAdapterVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAdapterVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAdapterVersions page.
func (p *ListAdapterVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAdapterVersionsOutput, error) {
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

	result, err := p.client.ListAdapterVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAdapterVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAdapterVersions",
	}
}
