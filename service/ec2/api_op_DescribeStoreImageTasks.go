// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the progress of the AMI store tasks. You can describe the store tasks
// for specified AMIs. If you don't specify the AMIs, you get a paginated list of
// store tasks from the last 31 days. For each AMI task, the response indicates if
// the task is InProgress , Completed , or Failed . For tasks InProgress , the
// response shows the estimated progress as a percentage. Tasks are listed in
// reverse chronological order. Currently, only tasks from the past 31 days can be
// viewed. To use this API, you must have the required permissions. For more
// information, see Permissions for storing and restoring AMIs using Amazon S3 (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html#ami-s3-permissions)
// in the Amazon EC2 User Guide. For more information, see Store and restore an
// AMI using Amazon S3 (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-store-restore.html)
// in the Amazon EC2 User Guide.
func (c *Client) DescribeStoreImageTasks(ctx context.Context, params *DescribeStoreImageTasksInput, optFns ...func(*Options)) (*DescribeStoreImageTasksOutput, error) {
	if params == nil {
		params = &DescribeStoreImageTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStoreImageTasks", params, optFns, c.addOperationDescribeStoreImageTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStoreImageTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStoreImageTasksInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - task-state - Returns tasks in a certain state ( InProgress | Completed |
	//   Failed )
	//   - bucket - Returns task information for tasks that targeted a specific bucket.
	//   For the filter value, specify the bucket name.
	// When you specify the ImageIds parameter, any filters that you specify are
	// ignored. To use the filters, you must remove the ImageIds parameter.
	Filters []types.Filter

	// The AMI IDs for which to show progress. Up to 20 AMI IDs can be included in a
	// request.
	ImageIds []string

	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// . You cannot specify this parameter and the ImageIds parameter in the same call.
	MaxResults *int32

	// The token returned from a previous paginated request. Pagination continues from
	// the end of the items returned by the previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeStoreImageTasksOutput struct {

	// The token to include in another request to get the next page of items. This
	// value is null when there are no more items to return.
	NextToken *string

	// The information about the AMI store tasks.
	StoreImageTaskResults []types.StoreImageTaskResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStoreImageTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeStoreImageTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeStoreImageTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStoreImageTasks"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStoreImageTasks(options.Region), middleware.Before); err != nil {
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

// DescribeStoreImageTasksAPIClient is a client that implements the
// DescribeStoreImageTasks operation.
type DescribeStoreImageTasksAPIClient interface {
	DescribeStoreImageTasks(context.Context, *DescribeStoreImageTasksInput, ...func(*Options)) (*DescribeStoreImageTasksOutput, error)
}

var _ DescribeStoreImageTasksAPIClient = (*Client)(nil)

// DescribeStoreImageTasksPaginatorOptions is the paginator options for
// DescribeStoreImageTasks
type DescribeStoreImageTasksPaginatorOptions struct {
	// The maximum number of items to return for this request. To get the next page of
	// items, make another request with the token returned in the output. For more
	// information, see Pagination (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html#api-pagination)
	// . You cannot specify this parameter and the ImageIds parameter in the same call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeStoreImageTasksPaginator is a paginator for DescribeStoreImageTasks
type DescribeStoreImageTasksPaginator struct {
	options   DescribeStoreImageTasksPaginatorOptions
	client    DescribeStoreImageTasksAPIClient
	params    *DescribeStoreImageTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeStoreImageTasksPaginator returns a new
// DescribeStoreImageTasksPaginator
func NewDescribeStoreImageTasksPaginator(client DescribeStoreImageTasksAPIClient, params *DescribeStoreImageTasksInput, optFns ...func(*DescribeStoreImageTasksPaginatorOptions)) *DescribeStoreImageTasksPaginator {
	if params == nil {
		params = &DescribeStoreImageTasksInput{}
	}

	options := DescribeStoreImageTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeStoreImageTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeStoreImageTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeStoreImageTasks page.
func (p *DescribeStoreImageTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeStoreImageTasksOutput, error) {
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

	result, err := p.client.DescribeStoreImageTasks(ctx, &params, optFns...)
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

// StoreImageTaskCompleteWaiterOptions are waiter options for
// StoreImageTaskCompleteWaiter
type StoreImageTaskCompleteWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// StoreImageTaskCompleteWaiter will use default minimum delay of 5 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, StoreImageTaskCompleteWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeStoreImageTasksInput, *DescribeStoreImageTasksOutput, error) (bool, error)
}

// StoreImageTaskCompleteWaiter defines the waiters for StoreImageTaskComplete
type StoreImageTaskCompleteWaiter struct {
	client DescribeStoreImageTasksAPIClient

	options StoreImageTaskCompleteWaiterOptions
}

// NewStoreImageTaskCompleteWaiter constructs a StoreImageTaskCompleteWaiter.
func NewStoreImageTaskCompleteWaiter(client DescribeStoreImageTasksAPIClient, optFns ...func(*StoreImageTaskCompleteWaiterOptions)) *StoreImageTaskCompleteWaiter {
	options := StoreImageTaskCompleteWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = storeImageTaskCompleteStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &StoreImageTaskCompleteWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for StoreImageTaskComplete waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *StoreImageTaskCompleteWaiter) Wait(ctx context.Context, params *DescribeStoreImageTasksInput, maxWaitDur time.Duration, optFns ...func(*StoreImageTaskCompleteWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for StoreImageTaskComplete waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *StoreImageTaskCompleteWaiter) WaitForOutput(ctx context.Context, params *DescribeStoreImageTasksInput, maxWaitDur time.Duration, optFns ...func(*StoreImageTaskCompleteWaiterOptions)) (*DescribeStoreImageTasksOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeStoreImageTasks(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for StoreImageTaskComplete waiter")
}

func storeImageTaskCompleteStateRetryable(ctx context.Context, input *DescribeStoreImageTasksInput, output *DescribeStoreImageTasksOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("StoreImageTaskResults[].StoreTaskState", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Completed"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("StoreImageTaskResults[].StoreTaskState", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("StoreImageTaskResults[].StoreTaskState", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "InProgress"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return true, nil
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeStoreImageTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStoreImageTasks",
	}
}
