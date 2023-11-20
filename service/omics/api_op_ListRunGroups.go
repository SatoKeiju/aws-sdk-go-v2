// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of run groups.
func (c *Client) ListRunGroups(ctx context.Context, params *ListRunGroupsInput, optFns ...func(*Options)) (*ListRunGroupsOutput, error) {
	if params == nil {
		params = &ListRunGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRunGroups", params, optFns, c.addOperationListRunGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRunGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRunGroupsInput struct {

	// The maximum number of run groups to return in one page of results.
	MaxResults *int32

	// The run groups' name.
	Name *string

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	StartingToken *string

	noSmithyDocumentSerde
}

type ListRunGroupsOutput struct {

	// A list of groups.
	Items []types.RunGroupListItem

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRunGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRunGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRunGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRunGroups"); err != nil {
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
	if err = addEndpointPrefix_opListRunGroupsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRunGroups(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListRunGroupsMiddleware struct {
}

func (*endpointPrefix_opListRunGroupsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListRunGroupsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "workflows-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opListRunGroupsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opListRunGroupsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// ListRunGroupsAPIClient is a client that implements the ListRunGroups operation.
type ListRunGroupsAPIClient interface {
	ListRunGroups(context.Context, *ListRunGroupsInput, ...func(*Options)) (*ListRunGroupsOutput, error)
}

var _ ListRunGroupsAPIClient = (*Client)(nil)

// ListRunGroupsPaginatorOptions is the paginator options for ListRunGroups
type ListRunGroupsPaginatorOptions struct {
	// The maximum number of run groups to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRunGroupsPaginator is a paginator for ListRunGroups
type ListRunGroupsPaginator struct {
	options   ListRunGroupsPaginatorOptions
	client    ListRunGroupsAPIClient
	params    *ListRunGroupsInput
	nextToken *string
	firstPage bool
}

// NewListRunGroupsPaginator returns a new ListRunGroupsPaginator
func NewListRunGroupsPaginator(client ListRunGroupsAPIClient, params *ListRunGroupsInput, optFns ...func(*ListRunGroupsPaginatorOptions)) *ListRunGroupsPaginator {
	if params == nil {
		params = &ListRunGroupsInput{}
	}

	options := ListRunGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRunGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.StartingToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRunGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRunGroups page.
func (p *ListRunGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRunGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.StartingToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListRunGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRunGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRunGroups",
	}
}
