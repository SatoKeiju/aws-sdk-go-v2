// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of simulation jobs. You can optionally provide filters to
// retrieve specific simulation jobs.
func (c *Client) ListSimulationJobs(ctx context.Context, params *ListSimulationJobsInput, optFns ...func(*Options)) (*ListSimulationJobsOutput, error) {
	if params == nil {
		params = &ListSimulationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSimulationJobs", params, optFns, c.addOperationListSimulationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSimulationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSimulationJobsInput struct {

	// Optional filters to limit results. The filter names status and
	// simulationApplicationName and robotApplicationName are supported. When
	// filtering, you must use the complete value of the filtered item. You can use up
	// to three filters, but they must be for the same named item. For example, if you
	// are looking for items with the status Preparing or the status Running .
	Filters []types.Filter

	// When this parameter is used, ListSimulationJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListSimulationJobs
	// request with the returned nextToken value. This value can be between 1 and
	// 1000. If this parameter is not used, then ListSimulationJobs returns up to 1000
	// results and a nextToken value if applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListSimulationJobs again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSimulationJobsOutput struct {

	// A list of simulation job summaries that meet the criteria of the request.
	//
	// This member is required.
	SimulationJobSummaries []types.SimulationJobSummary

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListSimulationJobs again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSimulationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSimulationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSimulationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSimulationJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSimulationJobs(options.Region), middleware.Before); err != nil {
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

// ListSimulationJobsAPIClient is a client that implements the ListSimulationJobs
// operation.
type ListSimulationJobsAPIClient interface {
	ListSimulationJobs(context.Context, *ListSimulationJobsInput, ...func(*Options)) (*ListSimulationJobsOutput, error)
}

var _ ListSimulationJobsAPIClient = (*Client)(nil)

// ListSimulationJobsPaginatorOptions is the paginator options for
// ListSimulationJobs
type ListSimulationJobsPaginatorOptions struct {
	// When this parameter is used, ListSimulationJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListSimulationJobs
	// request with the returned nextToken value. This value can be between 1 and
	// 1000. If this parameter is not used, then ListSimulationJobs returns up to 1000
	// results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSimulationJobsPaginator is a paginator for ListSimulationJobs
type ListSimulationJobsPaginator struct {
	options   ListSimulationJobsPaginatorOptions
	client    ListSimulationJobsAPIClient
	params    *ListSimulationJobsInput
	nextToken *string
	firstPage bool
}

// NewListSimulationJobsPaginator returns a new ListSimulationJobsPaginator
func NewListSimulationJobsPaginator(client ListSimulationJobsAPIClient, params *ListSimulationJobsInput, optFns ...func(*ListSimulationJobsPaginatorOptions)) *ListSimulationJobsPaginator {
	if params == nil {
		params = &ListSimulationJobsInput{}
	}

	options := ListSimulationJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSimulationJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSimulationJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSimulationJobs page.
func (p *ListSimulationJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSimulationJobsOutput, error) {
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

	result, err := p.client.ListSimulationJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSimulationJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSimulationJobs",
	}
}
