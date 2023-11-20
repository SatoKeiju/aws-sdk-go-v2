// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the state of the specified instances with respect to the specified
// load balancer. If no instances are specified, the call describes the state of
// all instances that are currently registered with the load balancer. If instances
// are specified, their state is returned even if they are no longer registered
// with the load balancer. The state of terminated instances is not returned.
func (c *Client) DescribeInstanceHealth(ctx context.Context, params *DescribeInstanceHealthInput, optFns ...func(*Options)) (*DescribeInstanceHealthOutput, error) {
	if params == nil {
		params = &DescribeInstanceHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceHealth", params, optFns, c.addOperationDescribeInstanceHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeInstanceHealth.
type DescribeInstanceHealthInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The IDs of the instances.
	Instances []types.Instance

	noSmithyDocumentSerde
}

// Contains the output for DescribeInstanceHealth.
type DescribeInstanceHealthOutput struct {

	// Information about the health of the instances.
	InstanceStates []types.InstanceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeInstanceHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeInstanceHealth{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInstanceHealth"); err != nil {
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
	if err = addOpDescribeInstanceHealthValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceHealth(options.Region), middleware.Before); err != nil {
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

// DescribeInstanceHealthAPIClient is a client that implements the
// DescribeInstanceHealth operation.
type DescribeInstanceHealthAPIClient interface {
	DescribeInstanceHealth(context.Context, *DescribeInstanceHealthInput, ...func(*Options)) (*DescribeInstanceHealthOutput, error)
}

var _ DescribeInstanceHealthAPIClient = (*Client)(nil)

// AnyInstanceInServiceWaiterOptions are waiter options for
// AnyInstanceInServiceWaiter
type AnyInstanceInServiceWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// AnyInstanceInServiceWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, AnyInstanceInServiceWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeInstanceHealthInput, *DescribeInstanceHealthOutput, error) (bool, error)
}

// AnyInstanceInServiceWaiter defines the waiters for AnyInstanceInService
type AnyInstanceInServiceWaiter struct {
	client DescribeInstanceHealthAPIClient

	options AnyInstanceInServiceWaiterOptions
}

// NewAnyInstanceInServiceWaiter constructs a AnyInstanceInServiceWaiter.
func NewAnyInstanceInServiceWaiter(client DescribeInstanceHealthAPIClient, optFns ...func(*AnyInstanceInServiceWaiterOptions)) *AnyInstanceInServiceWaiter {
	options := AnyInstanceInServiceWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = anyInstanceInServiceStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &AnyInstanceInServiceWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for AnyInstanceInService waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *AnyInstanceInServiceWaiter) Wait(ctx context.Context, params *DescribeInstanceHealthInput, maxWaitDur time.Duration, optFns ...func(*AnyInstanceInServiceWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for AnyInstanceInService waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *AnyInstanceInServiceWaiter) WaitForOutput(ctx context.Context, params *DescribeInstanceHealthInput, maxWaitDur time.Duration, optFns ...func(*AnyInstanceInServiceWaiterOptions)) (*DescribeInstanceHealthOutput, error) {
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

		out, err := w.client.DescribeInstanceHealth(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for AnyInstanceInService waiter")
}

func anyInstanceInServiceStateRetryable(ctx context.Context, input *DescribeInstanceHealthInput, output *DescribeInstanceHealthOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("InstanceStates[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "InService"
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
				return false, nil
			}
		}
	}

	return true, nil
}

// InstanceDeregisteredWaiterOptions are waiter options for
// InstanceDeregisteredWaiter
type InstanceDeregisteredWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InstanceDeregisteredWaiter will use default minimum delay of 15 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, InstanceDeregisteredWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeInstanceHealthInput, *DescribeInstanceHealthOutput, error) (bool, error)
}

// InstanceDeregisteredWaiter defines the waiters for InstanceDeregistered
type InstanceDeregisteredWaiter struct {
	client DescribeInstanceHealthAPIClient

	options InstanceDeregisteredWaiterOptions
}

// NewInstanceDeregisteredWaiter constructs a InstanceDeregisteredWaiter.
func NewInstanceDeregisteredWaiter(client DescribeInstanceHealthAPIClient, optFns ...func(*InstanceDeregisteredWaiterOptions)) *InstanceDeregisteredWaiter {
	options := InstanceDeregisteredWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = instanceDeregisteredStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &InstanceDeregisteredWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for InstanceDeregistered waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *InstanceDeregisteredWaiter) Wait(ctx context.Context, params *DescribeInstanceHealthInput, maxWaitDur time.Duration, optFns ...func(*InstanceDeregisteredWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for InstanceDeregistered waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *InstanceDeregisteredWaiter) WaitForOutput(ctx context.Context, params *DescribeInstanceHealthInput, maxWaitDur time.Duration, optFns ...func(*InstanceDeregisteredWaiterOptions)) (*DescribeInstanceHealthOutput, error) {
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

		out, err := w.client.DescribeInstanceHealth(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for InstanceDeregistered waiter")
}

func instanceDeregisteredStateRetryable(ctx context.Context, input *DescribeInstanceHealthInput, output *DescribeInstanceHealthOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("InstanceStates[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "OutOfService"
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

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidInstance" == apiErr.ErrorCode() {
			return false, nil
		}
	}

	return true, nil
}

// InstanceInServiceWaiterOptions are waiter options for InstanceInServiceWaiter
type InstanceInServiceWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InstanceInServiceWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, InstanceInServiceWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeInstanceHealthInput, *DescribeInstanceHealthOutput, error) (bool, error)
}

// InstanceInServiceWaiter defines the waiters for InstanceInService
type InstanceInServiceWaiter struct {
	client DescribeInstanceHealthAPIClient

	options InstanceInServiceWaiterOptions
}

// NewInstanceInServiceWaiter constructs a InstanceInServiceWaiter.
func NewInstanceInServiceWaiter(client DescribeInstanceHealthAPIClient, optFns ...func(*InstanceInServiceWaiterOptions)) *InstanceInServiceWaiter {
	options := InstanceInServiceWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = instanceInServiceStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &InstanceInServiceWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for InstanceInService waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *InstanceInServiceWaiter) Wait(ctx context.Context, params *DescribeInstanceHealthInput, maxWaitDur time.Duration, optFns ...func(*InstanceInServiceWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for InstanceInService waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *InstanceInServiceWaiter) WaitForOutput(ctx context.Context, params *DescribeInstanceHealthInput, maxWaitDur time.Duration, optFns ...func(*InstanceInServiceWaiterOptions)) (*DescribeInstanceHealthOutput, error) {
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

		out, err := w.client.DescribeInstanceHealth(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for InstanceInService waiter")
}

func instanceInServiceStateRetryable(ctx context.Context, input *DescribeInstanceHealthInput, output *DescribeInstanceHealthOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("InstanceStates[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "InService"
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

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidInstance" == apiErr.ErrorCode() {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeInstanceHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInstanceHealth",
	}
}
