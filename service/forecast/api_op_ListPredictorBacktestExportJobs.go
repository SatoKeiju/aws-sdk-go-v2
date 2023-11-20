// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of predictor backtest export jobs created using the
// CreatePredictorBacktestExportJob operation. This operation returns a summary for
// each backtest export job. You can filter the list using an array of Filter
// objects. To retrieve the complete set of properties for a particular backtest
// export job, use the ARN with the DescribePredictorBacktestExportJob operation.
func (c *Client) ListPredictorBacktestExportJobs(ctx context.Context, params *ListPredictorBacktestExportJobsInput, optFns ...func(*Options)) (*ListPredictorBacktestExportJobsOutput, error) {
	if params == nil {
		params = &ListPredictorBacktestExportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPredictorBacktestExportJobs", params, optFns, c.addOperationListPredictorBacktestExportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPredictorBacktestExportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPredictorBacktestExportJobsInput struct {

	// An array of filters. For each filter, provide a condition and a match
	// statement. The condition is either IS or IS_NOT , which specifies whether to
	// include or exclude the predictor backtest export jobs that match the statement
	// from the list. The match statement consists of a key and a value. Filter
	// properties
	//   - Condition - The condition to apply. Valid values are IS and IS_NOT . To
	//   include the predictor backtest export jobs that match the statement, specify
	//   IS . To exclude matching predictor backtest export jobs, specify IS_NOT .
	//   - Key - The name of the parameter to filter on. Valid values are PredictorArn
	//   and Status .
	//   - Value - The value to match.
	Filters []types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPredictorBacktestExportJobsOutput struct {

	// Returns this token if the response is truncated. To retrieve the next set of
	// results, use the token in the next request.
	NextToken *string

	// An array of objects that summarize the properties of each predictor backtest
	// export job.
	PredictorBacktestExportJobs []types.PredictorBacktestExportJobSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPredictorBacktestExportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPredictorBacktestExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPredictorBacktestExportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPredictorBacktestExportJobs"); err != nil {
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
	if err = addOpListPredictorBacktestExportJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPredictorBacktestExportJobs(options.Region), middleware.Before); err != nil {
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

// ListPredictorBacktestExportJobsAPIClient is a client that implements the
// ListPredictorBacktestExportJobs operation.
type ListPredictorBacktestExportJobsAPIClient interface {
	ListPredictorBacktestExportJobs(context.Context, *ListPredictorBacktestExportJobsInput, ...func(*Options)) (*ListPredictorBacktestExportJobsOutput, error)
}

var _ ListPredictorBacktestExportJobsAPIClient = (*Client)(nil)

// ListPredictorBacktestExportJobsPaginatorOptions is the paginator options for
// ListPredictorBacktestExportJobs
type ListPredictorBacktestExportJobsPaginatorOptions struct {
	// The number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPredictorBacktestExportJobsPaginator is a paginator for
// ListPredictorBacktestExportJobs
type ListPredictorBacktestExportJobsPaginator struct {
	options   ListPredictorBacktestExportJobsPaginatorOptions
	client    ListPredictorBacktestExportJobsAPIClient
	params    *ListPredictorBacktestExportJobsInput
	nextToken *string
	firstPage bool
}

// NewListPredictorBacktestExportJobsPaginator returns a new
// ListPredictorBacktestExportJobsPaginator
func NewListPredictorBacktestExportJobsPaginator(client ListPredictorBacktestExportJobsAPIClient, params *ListPredictorBacktestExportJobsInput, optFns ...func(*ListPredictorBacktestExportJobsPaginatorOptions)) *ListPredictorBacktestExportJobsPaginator {
	if params == nil {
		params = &ListPredictorBacktestExportJobsInput{}
	}

	options := ListPredictorBacktestExportJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPredictorBacktestExportJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPredictorBacktestExportJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPredictorBacktestExportJobs page.
func (p *ListPredictorBacktestExportJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPredictorBacktestExportJobsOutput, error) {
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

	result, err := p.client.ListPredictorBacktestExportJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPredictorBacktestExportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPredictorBacktestExportJobs",
	}
}
