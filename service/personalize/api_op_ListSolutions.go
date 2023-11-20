// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of solutions that use the given dataset group. When a dataset
// group is not specified, all the solutions associated with the account are
// listed. The response provides the properties for each solution, including the
// Amazon Resource Name (ARN). For more information on solutions, see
// CreateSolution (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateSolution.html)
// .
func (c *Client) ListSolutions(ctx context.Context, params *ListSolutionsInput, optFns ...func(*Options)) (*ListSolutionsOutput, error) {
	if params == nil {
		params = &ListSolutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSolutions", params, optFns, c.addOperationListSolutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSolutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSolutionsInput struct {

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string

	// The maximum number of solutions to return.
	MaxResults *int32

	// A token returned from the previous call to ListSolutions for getting the next
	// set of solutions (if they exist).
	NextToken *string

	noSmithyDocumentSerde
}

type ListSolutionsOutput struct {

	// A token for getting the next set of solutions (if they exist).
	NextToken *string

	// A list of the current solutions.
	Solutions []types.SolutionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSolutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSolutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSolutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSolutions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSolutions(options.Region), middleware.Before); err != nil {
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

// ListSolutionsAPIClient is a client that implements the ListSolutions operation.
type ListSolutionsAPIClient interface {
	ListSolutions(context.Context, *ListSolutionsInput, ...func(*Options)) (*ListSolutionsOutput, error)
}

var _ ListSolutionsAPIClient = (*Client)(nil)

// ListSolutionsPaginatorOptions is the paginator options for ListSolutions
type ListSolutionsPaginatorOptions struct {
	// The maximum number of solutions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSolutionsPaginator is a paginator for ListSolutions
type ListSolutionsPaginator struct {
	options   ListSolutionsPaginatorOptions
	client    ListSolutionsAPIClient
	params    *ListSolutionsInput
	nextToken *string
	firstPage bool
}

// NewListSolutionsPaginator returns a new ListSolutionsPaginator
func NewListSolutionsPaginator(client ListSolutionsAPIClient, params *ListSolutionsInput, optFns ...func(*ListSolutionsPaginatorOptions)) *ListSolutionsPaginator {
	if params == nil {
		params = &ListSolutionsInput{}
	}

	options := ListSolutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSolutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSolutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSolutions page.
func (p *ListSolutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSolutionsOutput, error) {
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

	result, err := p.client.ListSolutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSolutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSolutions",
	}
}
