// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of anomalous metrics for a measure in an anomaly group.
func (c *Client) ListAnomalyGroupTimeSeries(ctx context.Context, params *ListAnomalyGroupTimeSeriesInput, optFns ...func(*Options)) (*ListAnomalyGroupTimeSeriesOutput, error) {
	if params == nil {
		params = &ListAnomalyGroupTimeSeriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnomalyGroupTimeSeries", params, optFns, c.addOperationListAnomalyGroupTimeSeriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnomalyGroupTimeSeriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAnomalyGroupTimeSeriesInput struct {

	// The Amazon Resource Name (ARN) of the anomaly detector.
	//
	// This member is required.
	AnomalyDetectorArn *string

	// The ID of the anomaly group.
	//
	// This member is required.
	AnomalyGroupId *string

	// The name of the measure field.
	//
	// This member is required.
	MetricName *string

	// The maximum number of results to return.
	MaxResults *int32

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAnomalyGroupTimeSeriesOutput struct {

	// The ID of the anomaly group.
	AnomalyGroupId *string

	// The name of the measure field.
	MetricName *string

	// The pagination token that's included if more results are available.
	NextToken *string

	// A list of anomalous metrics.
	TimeSeriesList []types.TimeSeries

	// Timestamps for the anomalous metrics.
	TimestampList []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAnomalyGroupTimeSeriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnomalyGroupTimeSeries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnomalyGroupTimeSeries{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAnomalyGroupTimeSeries"); err != nil {
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
	if err = addOpListAnomalyGroupTimeSeriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAnomalyGroupTimeSeries(options.Region), middleware.Before); err != nil {
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

// ListAnomalyGroupTimeSeriesAPIClient is a client that implements the
// ListAnomalyGroupTimeSeries operation.
type ListAnomalyGroupTimeSeriesAPIClient interface {
	ListAnomalyGroupTimeSeries(context.Context, *ListAnomalyGroupTimeSeriesInput, ...func(*Options)) (*ListAnomalyGroupTimeSeriesOutput, error)
}

var _ ListAnomalyGroupTimeSeriesAPIClient = (*Client)(nil)

// ListAnomalyGroupTimeSeriesPaginatorOptions is the paginator options for
// ListAnomalyGroupTimeSeries
type ListAnomalyGroupTimeSeriesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnomalyGroupTimeSeriesPaginator is a paginator for
// ListAnomalyGroupTimeSeries
type ListAnomalyGroupTimeSeriesPaginator struct {
	options   ListAnomalyGroupTimeSeriesPaginatorOptions
	client    ListAnomalyGroupTimeSeriesAPIClient
	params    *ListAnomalyGroupTimeSeriesInput
	nextToken *string
	firstPage bool
}

// NewListAnomalyGroupTimeSeriesPaginator returns a new
// ListAnomalyGroupTimeSeriesPaginator
func NewListAnomalyGroupTimeSeriesPaginator(client ListAnomalyGroupTimeSeriesAPIClient, params *ListAnomalyGroupTimeSeriesInput, optFns ...func(*ListAnomalyGroupTimeSeriesPaginatorOptions)) *ListAnomalyGroupTimeSeriesPaginator {
	if params == nil {
		params = &ListAnomalyGroupTimeSeriesInput{}
	}

	options := ListAnomalyGroupTimeSeriesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAnomalyGroupTimeSeriesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnomalyGroupTimeSeriesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAnomalyGroupTimeSeries page.
func (p *ListAnomalyGroupTimeSeriesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnomalyGroupTimeSeriesOutput, error) {
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

	result, err := p.client.ListAnomalyGroupTimeSeries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAnomalyGroupTimeSeries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAnomalyGroupTimeSeries",
	}
}
