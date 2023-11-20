// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Get information about a specific bot alias.
func (c *Client) DescribeBotAlias(ctx context.Context, params *DescribeBotAliasInput, optFns ...func(*Options)) (*DescribeBotAliasOutput, error) {
	if params == nil {
		params = &DescribeBotAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBotAlias", params, optFns, c.addOperationDescribeBotAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBotAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBotAliasInput struct {

	// The identifier of the bot alias to describe.
	//
	// This member is required.
	BotAliasId *string

	// The identifier of the bot associated with the bot alias to describe.
	//
	// This member is required.
	BotId *string

	noSmithyDocumentSerde
}

type DescribeBotAliasOutput struct {

	// A list of events that affect a bot alias. For example, an event is recorded
	// when the version that the alias points to changes.
	BotAliasHistoryEvents []types.BotAliasHistoryEvent

	// The identifier of the bot alias.
	BotAliasId *string

	// The locale settings that are unique to the alias.
	BotAliasLocaleSettings map[string]types.BotAliasLocaleSettings

	// The name of the bot alias.
	BotAliasName *string

	// The current status of the alias. When the alias is Available , the alias is
	// ready for use with your bot.
	BotAliasStatus types.BotAliasStatus

	// The identifier of the bot associated with the bot alias.
	BotId *string

	// The version of the bot associated with the bot alias.
	BotVersion *string

	// Specifics of how Amazon Lex logs text and audio conversations with the bot
	// associated with the alias.
	ConversationLogSettings *types.ConversationLogSettings

	// A timestamp of the date and time that the alias was created.
	CreationDateTime *time.Time

	// The description of the bot alias.
	Description *string

	// A timestamp of the date and time that the alias was last updated.
	LastUpdatedDateTime *time.Time

	// A list of the networks to which the bot alias you described belongs.
	ParentBotNetworks []types.ParentBotNetwork

	// Determines whether Amazon Lex will use Amazon Comprehend to detect the
	// sentiment of user utterances.
	SentimentAnalysisSettings *types.SentimentAnalysisSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBotAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBotAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBotAlias{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBotAlias"); err != nil {
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
	if err = addOpDescribeBotAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBotAlias(options.Region), middleware.Before); err != nil {
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

// DescribeBotAliasAPIClient is a client that implements the DescribeBotAlias
// operation.
type DescribeBotAliasAPIClient interface {
	DescribeBotAlias(context.Context, *DescribeBotAliasInput, ...func(*Options)) (*DescribeBotAliasOutput, error)
}

var _ DescribeBotAliasAPIClient = (*Client)(nil)

// BotAliasAvailableWaiterOptions are waiter options for BotAliasAvailableWaiter
type BotAliasAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BotAliasAvailableWaiter will use default minimum delay of 10 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, BotAliasAvailableWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeBotAliasInput, *DescribeBotAliasOutput, error) (bool, error)
}

// BotAliasAvailableWaiter defines the waiters for BotAliasAvailable
type BotAliasAvailableWaiter struct {
	client DescribeBotAliasAPIClient

	options BotAliasAvailableWaiterOptions
}

// NewBotAliasAvailableWaiter constructs a BotAliasAvailableWaiter.
func NewBotAliasAvailableWaiter(client DescribeBotAliasAPIClient, optFns ...func(*BotAliasAvailableWaiterOptions)) *BotAliasAvailableWaiter {
	options := BotAliasAvailableWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = botAliasAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BotAliasAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BotAliasAvailable waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *BotAliasAvailableWaiter) Wait(ctx context.Context, params *DescribeBotAliasInput, maxWaitDur time.Duration, optFns ...func(*BotAliasAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for BotAliasAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *BotAliasAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeBotAliasInput, maxWaitDur time.Duration, optFns ...func(*BotAliasAvailableWaiterOptions)) (*DescribeBotAliasOutput, error) {
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

		out, err := w.client.DescribeBotAlias(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for BotAliasAvailable waiter")
}

func botAliasAvailableStateRetryable(ctx context.Context, input *DescribeBotAliasInput, output *DescribeBotAliasOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("botAliasStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Available"
		value, ok := pathValue.(types.BotAliasStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotAliasStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botAliasStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.BotAliasStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotAliasStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("botAliasStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Deleting"
		value, ok := pathValue.(types.BotAliasStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.BotAliasStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeBotAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBotAlias",
	}
}
