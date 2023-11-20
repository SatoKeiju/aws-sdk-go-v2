// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of available query execution IDs for the queries in the
// specified workgroup. Athena keeps a query history for 45 days. If a workgroup is
// not specified, returns a list of query execution IDs for the primary workgroup.
// Requires you to have access to the workgroup in which the queries ran.
func (c *Client) ListQueryExecutions(ctx context.Context, params *ListQueryExecutionsInput, optFns ...func(*Options)) (*ListQueryExecutionsOutput, error) {
	if params == nil {
		params = &ListQueryExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueryExecutions", params, optFns, c.addOperationListQueryExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueryExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueryExecutionsInput struct {

	// The maximum number of query executions to return in this request.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// The name of the workgroup from which queries are being returned. If a workgroup
	// is not specified, a list of available query execution IDs for the queries in the
	// primary workgroup is returned.
	WorkGroup *string

	noSmithyDocumentSerde
}

type ListQueryExecutionsOutput struct {

	// A token to be used by the next request if this request is truncated.
	NextToken *string

	// The unique IDs of each query execution as an array of strings.
	QueryExecutionIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueryExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListQueryExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListQueryExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListQueryExecutions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueryExecutions(options.Region), middleware.Before); err != nil {
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

// ListQueryExecutionsAPIClient is a client that implements the
// ListQueryExecutions operation.
type ListQueryExecutionsAPIClient interface {
	ListQueryExecutions(context.Context, *ListQueryExecutionsInput, ...func(*Options)) (*ListQueryExecutionsOutput, error)
}

var _ ListQueryExecutionsAPIClient = (*Client)(nil)

// ListQueryExecutionsPaginatorOptions is the paginator options for
// ListQueryExecutions
type ListQueryExecutionsPaginatorOptions struct {
	// The maximum number of query executions to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueryExecutionsPaginator is a paginator for ListQueryExecutions
type ListQueryExecutionsPaginator struct {
	options   ListQueryExecutionsPaginatorOptions
	client    ListQueryExecutionsAPIClient
	params    *ListQueryExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListQueryExecutionsPaginator returns a new ListQueryExecutionsPaginator
func NewListQueryExecutionsPaginator(client ListQueryExecutionsAPIClient, params *ListQueryExecutionsInput, optFns ...func(*ListQueryExecutionsPaginatorOptions)) *ListQueryExecutionsPaginator {
	if params == nil {
		params = &ListQueryExecutionsInput{}
	}

	options := ListQueryExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueryExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueryExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueryExecutions page.
func (p *ListQueryExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueryExecutionsOutput, error) {
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

	result, err := p.client.ListQueryExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQueryExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListQueryExecutions",
	}
}
