// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns descriptive information about an Fargate profile.
func (c *Client) DescribeFargateProfile(ctx context.Context, params *DescribeFargateProfileInput, optFns ...func(*Options)) (*DescribeFargateProfileOutput, error) {
	if params == nil {
		params = &DescribeFargateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFargateProfile", params, optFns, c.addOperationDescribeFargateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFargateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFargateProfileInput struct {

	// The name of the Amazon EKS cluster associated with the Fargate profile.
	//
	// This member is required.
	ClusterName *string

	// The name of the Fargate profile to describe.
	//
	// This member is required.
	FargateProfileName *string

	noSmithyDocumentSerde
}

type DescribeFargateProfileOutput struct {

	// The full description of your Fargate profile.
	FargateProfile *types.FargateProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFargateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFargateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFargateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFargateProfile"); err != nil {
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
	if err = addOpDescribeFargateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFargateProfile(options.Region), middleware.Before); err != nil {
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

// DescribeFargateProfileAPIClient is a client that implements the
// DescribeFargateProfile operation.
type DescribeFargateProfileAPIClient interface {
	DescribeFargateProfile(context.Context, *DescribeFargateProfileInput, ...func(*Options)) (*DescribeFargateProfileOutput, error)
}

var _ DescribeFargateProfileAPIClient = (*Client)(nil)

// FargateProfileActiveWaiterOptions are waiter options for
// FargateProfileActiveWaiter
type FargateProfileActiveWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// FargateProfileActiveWaiter will use default minimum delay of 10 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, FargateProfileActiveWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeFargateProfileInput, *DescribeFargateProfileOutput, error) (bool, error)
}

// FargateProfileActiveWaiter defines the waiters for FargateProfileActive
type FargateProfileActiveWaiter struct {
	client DescribeFargateProfileAPIClient

	options FargateProfileActiveWaiterOptions
}

// NewFargateProfileActiveWaiter constructs a FargateProfileActiveWaiter.
func NewFargateProfileActiveWaiter(client DescribeFargateProfileAPIClient, optFns ...func(*FargateProfileActiveWaiterOptions)) *FargateProfileActiveWaiter {
	options := FargateProfileActiveWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = fargateProfileActiveStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &FargateProfileActiveWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for FargateProfileActive waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *FargateProfileActiveWaiter) Wait(ctx context.Context, params *DescribeFargateProfileInput, maxWaitDur time.Duration, optFns ...func(*FargateProfileActiveWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for FargateProfileActive waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *FargateProfileActiveWaiter) WaitForOutput(ctx context.Context, params *DescribeFargateProfileInput, maxWaitDur time.Duration, optFns ...func(*FargateProfileActiveWaiterOptions)) (*DescribeFargateProfileOutput, error) {
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

		out, err := w.client.DescribeFargateProfile(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for FargateProfileActive waiter")
}

func fargateProfileActiveStateRetryable(ctx context.Context, input *DescribeFargateProfileInput, output *DescribeFargateProfileOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("fargateProfile.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED"
		value, ok := pathValue.(types.FargateProfileStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.FargateProfileStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("fargateProfile.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.FargateProfileStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.FargateProfileStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	return true, nil
}

// FargateProfileDeletedWaiterOptions are waiter options for
// FargateProfileDeletedWaiter
type FargateProfileDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// FargateProfileDeletedWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, FargateProfileDeletedWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeFargateProfileInput, *DescribeFargateProfileOutput, error) (bool, error)
}

// FargateProfileDeletedWaiter defines the waiters for FargateProfileDeleted
type FargateProfileDeletedWaiter struct {
	client DescribeFargateProfileAPIClient

	options FargateProfileDeletedWaiterOptions
}

// NewFargateProfileDeletedWaiter constructs a FargateProfileDeletedWaiter.
func NewFargateProfileDeletedWaiter(client DescribeFargateProfileAPIClient, optFns ...func(*FargateProfileDeletedWaiterOptions)) *FargateProfileDeletedWaiter {
	options := FargateProfileDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = fargateProfileDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &FargateProfileDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for FargateProfileDeleted waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *FargateProfileDeletedWaiter) Wait(ctx context.Context, params *DescribeFargateProfileInput, maxWaitDur time.Duration, optFns ...func(*FargateProfileDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for FargateProfileDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *FargateProfileDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeFargateProfileInput, maxWaitDur time.Duration, optFns ...func(*FargateProfileDeletedWaiterOptions)) (*DescribeFargateProfileOutput, error) {
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

		out, err := w.client.DescribeFargateProfile(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for FargateProfileDeleted waiter")
}

func fargateProfileDeletedStateRetryable(ctx context.Context, input *DescribeFargateProfileInput, output *DescribeFargateProfileOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("fargateProfile.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETE_FAILED"
		value, ok := pathValue.(types.FargateProfileStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.FargateProfileStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeFargateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFargateProfile",
	}
}
