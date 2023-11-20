// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the runs meeting the filter criteria, where a ruleset is evaluated
// against a data source.
func (c *Client) ListDataQualityRulesetEvaluationRuns(ctx context.Context, params *ListDataQualityRulesetEvaluationRunsInput, optFns ...func(*Options)) (*ListDataQualityRulesetEvaluationRunsOutput, error) {
	if params == nil {
		params = &ListDataQualityRulesetEvaluationRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataQualityRulesetEvaluationRuns", params, optFns, c.addOperationListDataQualityRulesetEvaluationRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataQualityRulesetEvaluationRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataQualityRulesetEvaluationRunsInput struct {

	// The filter criteria.
	Filter *types.DataQualityRulesetEvaluationRunFilter

	// The maximum number of results to return.
	MaxResults *int32

	// A paginated token to offset the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataQualityRulesetEvaluationRunsOutput struct {

	// A pagination token, if more results are available.
	NextToken *string

	// A list of DataQualityRulesetEvaluationRunDescription objects representing data
	// quality ruleset runs.
	Runs []types.DataQualityRulesetEvaluationRunDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataQualityRulesetEvaluationRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDataQualityRulesetEvaluationRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDataQualityRulesetEvaluationRuns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataQualityRulesetEvaluationRuns"); err != nil {
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
	if err = addOpListDataQualityRulesetEvaluationRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataQualityRulesetEvaluationRuns(options.Region), middleware.Before); err != nil {
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

// ListDataQualityRulesetEvaluationRunsAPIClient is a client that implements the
// ListDataQualityRulesetEvaluationRuns operation.
type ListDataQualityRulesetEvaluationRunsAPIClient interface {
	ListDataQualityRulesetEvaluationRuns(context.Context, *ListDataQualityRulesetEvaluationRunsInput, ...func(*Options)) (*ListDataQualityRulesetEvaluationRunsOutput, error)
}

var _ ListDataQualityRulesetEvaluationRunsAPIClient = (*Client)(nil)

// ListDataQualityRulesetEvaluationRunsPaginatorOptions is the paginator options
// for ListDataQualityRulesetEvaluationRuns
type ListDataQualityRulesetEvaluationRunsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataQualityRulesetEvaluationRunsPaginator is a paginator for
// ListDataQualityRulesetEvaluationRuns
type ListDataQualityRulesetEvaluationRunsPaginator struct {
	options   ListDataQualityRulesetEvaluationRunsPaginatorOptions
	client    ListDataQualityRulesetEvaluationRunsAPIClient
	params    *ListDataQualityRulesetEvaluationRunsInput
	nextToken *string
	firstPage bool
}

// NewListDataQualityRulesetEvaluationRunsPaginator returns a new
// ListDataQualityRulesetEvaluationRunsPaginator
func NewListDataQualityRulesetEvaluationRunsPaginator(client ListDataQualityRulesetEvaluationRunsAPIClient, params *ListDataQualityRulesetEvaluationRunsInput, optFns ...func(*ListDataQualityRulesetEvaluationRunsPaginatorOptions)) *ListDataQualityRulesetEvaluationRunsPaginator {
	if params == nil {
		params = &ListDataQualityRulesetEvaluationRunsInput{}
	}

	options := ListDataQualityRulesetEvaluationRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataQualityRulesetEvaluationRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataQualityRulesetEvaluationRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataQualityRulesetEvaluationRuns page.
func (p *ListDataQualityRulesetEvaluationRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataQualityRulesetEvaluationRunsOutput, error) {
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

	result, err := p.client.ListDataQualityRulesetEvaluationRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataQualityRulesetEvaluationRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataQualityRulesetEvaluationRuns",
	}
}
