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

// Returns a list of what-if analyses created using the CreateWhatIfAnalysis
// operation. For each what-if analysis, this operation returns a summary of its
// properties, including its Amazon Resource Name (ARN). You can retrieve the
// complete set of properties by using the what-if analysis ARN with the
// DescribeWhatIfAnalysis operation.
func (c *Client) ListWhatIfAnalyses(ctx context.Context, params *ListWhatIfAnalysesInput, optFns ...func(*Options)) (*ListWhatIfAnalysesOutput, error) {
	if params == nil {
		params = &ListWhatIfAnalysesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWhatIfAnalyses", params, optFns, c.addOperationListWhatIfAnalysesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWhatIfAnalysesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWhatIfAnalysesInput struct {

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether to
	// include or exclude the what-if analysis jobs that match the statement from the
	// list, respectively. The match statement consists of a key and a value. Filter
	// properties
	//
	// * Condition - The condition to apply. Valid values are IS and
	// IS_NOT. To include the what-if analysis jobs that match the statement, specify
	// IS. To exclude matching what-if analysis jobs, specify IS_NOT.
	//
	// * Key - The name
	// of the parameter to filter on. Valid values are WhatIfAnalysisArn and Status.
	//
	// *
	// Value - The value to match.
	//
	// For example, to list all jobs that export a
	// forecast named electricityWhatIf, specify the following filter: "Filters": [ {
	// "Condition": "IS", "Key": "WhatIfAnalysisArn", "Value":
	// "arn:aws:forecast:us-west-2::forecast/electricityWhatIf" } ]
	Filters []types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWhatIfAnalysesOutput struct {

	// If the response is truncated, Forecast returns this token. To retrieve the next
	// set of results, use the token in the next request.
	NextToken *string

	// An array of WhatIfAnalysisSummary objects that describe the matched analyses.
	WhatIfAnalyses []types.WhatIfAnalysisSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWhatIfAnalysesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWhatIfAnalyses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWhatIfAnalyses{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListWhatIfAnalysesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWhatIfAnalyses(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListWhatIfAnalysesAPIClient is a client that implements the ListWhatIfAnalyses
// operation.
type ListWhatIfAnalysesAPIClient interface {
	ListWhatIfAnalyses(context.Context, *ListWhatIfAnalysesInput, ...func(*Options)) (*ListWhatIfAnalysesOutput, error)
}

var _ ListWhatIfAnalysesAPIClient = (*Client)(nil)

// ListWhatIfAnalysesPaginatorOptions is the paginator options for
// ListWhatIfAnalyses
type ListWhatIfAnalysesPaginatorOptions struct {
	// The number of items to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWhatIfAnalysesPaginator is a paginator for ListWhatIfAnalyses
type ListWhatIfAnalysesPaginator struct {
	options   ListWhatIfAnalysesPaginatorOptions
	client    ListWhatIfAnalysesAPIClient
	params    *ListWhatIfAnalysesInput
	nextToken *string
	firstPage bool
}

// NewListWhatIfAnalysesPaginator returns a new ListWhatIfAnalysesPaginator
func NewListWhatIfAnalysesPaginator(client ListWhatIfAnalysesAPIClient, params *ListWhatIfAnalysesInput, optFns ...func(*ListWhatIfAnalysesPaginatorOptions)) *ListWhatIfAnalysesPaginator {
	if params == nil {
		params = &ListWhatIfAnalysesInput{}
	}

	options := ListWhatIfAnalysesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWhatIfAnalysesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWhatIfAnalysesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWhatIfAnalyses page.
func (p *ListWhatIfAnalysesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWhatIfAnalysesOutput, error) {
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

	result, err := p.client.ListWhatIfAnalyses(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWhatIfAnalyses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListWhatIfAnalyses",
	}
}
