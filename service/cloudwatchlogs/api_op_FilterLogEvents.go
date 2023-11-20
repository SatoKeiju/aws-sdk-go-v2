// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists log events from the specified log group. You can list all the log events
// or filter the results using a filter pattern, a time range, and the name of the
// log stream. You must have the logs:FilterLogEvents permission to perform this
// operation. You can specify the log group to search by using either
// logGroupIdentifier or logGroupName . You must include one of these two
// parameters, but you can't include both. By default, this operation returns as
// many log events as can fit in 1 MB (up to 10,000 log events) or all the events
// found within the specified time range. If the results include a token, that
// means there are more log events available. You can get additional results by
// specifying the token in a subsequent call. This operation can return empty
// results while there are more log events available through the token. The
// returned log events are sorted by event timestamp, the timestamp when the event
// was ingested by CloudWatch Logs, and the ID of the PutLogEvents request. If you
// are using CloudWatch cross-account observability, you can use this operation in
// a monitoring account and view data from the linked source accounts. For more
// information, see CloudWatch cross-account observability (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html)
// .
func (c *Client) FilterLogEvents(ctx context.Context, params *FilterLogEventsInput, optFns ...func(*Options)) (*FilterLogEventsOutput, error) {
	if params == nil {
		params = &FilterLogEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FilterLogEvents", params, optFns, c.addOperationFilterLogEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FilterLogEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FilterLogEventsInput struct {

	// The end of the time range, expressed as the number of milliseconds after Jan 1,
	// 1970 00:00:00 UTC . Events with a timestamp later than this time are not
	// returned.
	EndTime *int64

	// The filter pattern to use. For more information, see Filter and Pattern Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html)
	// . If not provided, all the events are matched.
	FilterPattern *string

	// If the value is true, the operation attempts to provide responses that contain
	// events from multiple log streams within the log group, interleaved in a single
	// response. If the value is false, all the matched log events in the first log
	// stream are searched first, then those in the next log stream, and so on.
	// Important As of June 17, 2019, this parameter is ignored and the value is
	// assumed to be true. The response from this operation always interleaves events
	// from multiple log streams within a log group.
	//
	// Deprecated: Starting on June 17, 2019, this parameter will be ignored and the
	// value will be assumed to be true. The response from this operation will always
	// interleave events from multiple log streams within a log group.
	Interleaved *bool

	// The maximum number of events to return. The default is 10,000 events.
	Limit *int32

	// Specify either the name or ARN of the log group to view log events from. If the
	// log group is in a source account and you are using a monitoring account, you
	// must use the log group ARN. You must include either logGroupIdentifier or
	// logGroupName , but not both.
	LogGroupIdentifier *string

	// The name of the log group to search. You must include either logGroupIdentifier
	// or logGroupName , but not both.
	LogGroupName *string

	// Filters the results to include only events from log streams that have names
	// starting with this prefix. If you specify a value for both logStreamNamePrefix
	// and logStreamNames , but the value for logStreamNamePrefix does not match any
	// log stream names specified in logStreamNames , the action returns an
	// InvalidParameterException error.
	LogStreamNamePrefix *string

	// Filters the results to only logs from the log streams in this list. If you
	// specify a value for both logStreamNamePrefix and logStreamNames , the action
	// returns an InvalidParameterException error.
	LogStreamNames []string

	// The token for the next set of events to return. (You received this token from a
	// previous call.)
	NextToken *string

	// The start of the time range, expressed as the number of milliseconds after Jan
	// 1, 1970 00:00:00 UTC . Events with a timestamp before this time are not returned.
	StartTime *int64

	// Specify true to display the log event fields with all sensitive data unmasked
	// and visible. The default is false . To use this operation with this parameter,
	// you must be signed into an account with the logs:Unmask permission.
	Unmask bool

	noSmithyDocumentSerde
}

type FilterLogEventsOutput struct {

	// The matched events.
	Events []types.FilteredLogEvent

	// The token to use when requesting the next set of items. The token expires after
	// 24 hours.
	NextToken *string

	// Important As of May 15, 2020, this parameter is no longer supported. This
	// parameter returns an empty list. Indicates which log streams have been searched
	// and whether each has been searched completely.
	SearchedLogStreams []types.SearchedLogStream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationFilterLogEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpFilterLogEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpFilterLogEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "FilterLogEvents"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFilterLogEvents(options.Region), middleware.Before); err != nil {
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

// FilterLogEventsAPIClient is a client that implements the FilterLogEvents
// operation.
type FilterLogEventsAPIClient interface {
	FilterLogEvents(context.Context, *FilterLogEventsInput, ...func(*Options)) (*FilterLogEventsOutput, error)
}

var _ FilterLogEventsAPIClient = (*Client)(nil)

// FilterLogEventsPaginatorOptions is the paginator options for FilterLogEvents
type FilterLogEventsPaginatorOptions struct {
	// The maximum number of events to return. The default is 10,000 events.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// FilterLogEventsPaginator is a paginator for FilterLogEvents
type FilterLogEventsPaginator struct {
	options   FilterLogEventsPaginatorOptions
	client    FilterLogEventsAPIClient
	params    *FilterLogEventsInput
	nextToken *string
	firstPage bool
}

// NewFilterLogEventsPaginator returns a new FilterLogEventsPaginator
func NewFilterLogEventsPaginator(client FilterLogEventsAPIClient, params *FilterLogEventsInput, optFns ...func(*FilterLogEventsPaginatorOptions)) *FilterLogEventsPaginator {
	if params == nil {
		params = &FilterLogEventsInput{}
	}

	options := FilterLogEventsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &FilterLogEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *FilterLogEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next FilterLogEvents page.
func (p *FilterLogEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*FilterLogEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.FilterLogEvents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opFilterLogEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "FilterLogEvents",
	}
}
