// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"strconv"
	"time"
)

// Describes the tenant databases in a DB instance that uses the multi-tenant
// configuration. Only RDS for Oracle CDB instances are supported.
func (c *Client) DescribeTenantDatabases(ctx context.Context, params *DescribeTenantDatabasesInput, optFns ...func(*Options)) (*DescribeTenantDatabasesOutput, error) {
	if params == nil {
		params = &DescribeTenantDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTenantDatabases", params, optFns, c.addOperationDescribeTenantDatabasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTenantDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTenantDatabasesInput struct {

	// The user-supplied DB instance identifier, which must match the identifier of an
	// existing instance owned by the Amazon Web Services account. This parameter isn't
	// case-sensitive.
	DBInstanceIdentifier *string

	// A filter that specifies one or more database tenants to describe. Supported
	// filters:
	//   - tenant-db-name - Tenant database names. The results list only includes
	//   information about the tenant databases that match these tenant DB names.
	//   - tenant-database-resource-id - Tenant database resource identifiers.
	//   - dbi-resource-id - DB instance resource identifiers. The results list only
	//   includes information about the tenants contained within the DB instances
	//   identified by these resource identifiers.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeTenantDatabases
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	MaxRecords *int32

	// The user-supplied tenant database name, which must match the name of an
	// existing tenant database on the specified DB instance owned by your Amazon Web
	// Services account. This parameter isn’t case-sensitive.
	TenantDBName *string

	noSmithyDocumentSerde
}

type DescribeTenantDatabasesOutput struct {

	// An optional pagination token provided by a previous DescribeTenantDatabases
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// An array of the tenant databases requested by the DescribeTenantDatabases
	// operation.
	TenantDatabases []types.TenantDatabase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTenantDatabasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTenantDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTenantDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTenantDatabases"); err != nil {
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
	if err = addOpDescribeTenantDatabasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTenantDatabases(options.Region), middleware.Before); err != nil {
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

// DescribeTenantDatabasesAPIClient is a client that implements the
// DescribeTenantDatabases operation.
type DescribeTenantDatabasesAPIClient interface {
	DescribeTenantDatabases(context.Context, *DescribeTenantDatabasesInput, ...func(*Options)) (*DescribeTenantDatabasesOutput, error)
}

var _ DescribeTenantDatabasesAPIClient = (*Client)(nil)

// DescribeTenantDatabasesPaginatorOptions is the paginator options for
// DescribeTenantDatabases
type DescribeTenantDatabasesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTenantDatabasesPaginator is a paginator for DescribeTenantDatabases
type DescribeTenantDatabasesPaginator struct {
	options   DescribeTenantDatabasesPaginatorOptions
	client    DescribeTenantDatabasesAPIClient
	params    *DescribeTenantDatabasesInput
	nextToken *string
	firstPage bool
}

// NewDescribeTenantDatabasesPaginator returns a new
// DescribeTenantDatabasesPaginator
func NewDescribeTenantDatabasesPaginator(client DescribeTenantDatabasesAPIClient, params *DescribeTenantDatabasesInput, optFns ...func(*DescribeTenantDatabasesPaginatorOptions)) *DescribeTenantDatabasesPaginator {
	if params == nil {
		params = &DescribeTenantDatabasesInput{}
	}

	options := DescribeTenantDatabasesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTenantDatabasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTenantDatabasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTenantDatabases page.
func (p *DescribeTenantDatabasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTenantDatabasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeTenantDatabases(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// TenantDatabaseAvailableWaiterOptions are waiter options for
// TenantDatabaseAvailableWaiter
type TenantDatabaseAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// TenantDatabaseAvailableWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, TenantDatabaseAvailableWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeTenantDatabasesInput, *DescribeTenantDatabasesOutput, error) (bool, error)
}

// TenantDatabaseAvailableWaiter defines the waiters for TenantDatabaseAvailable
type TenantDatabaseAvailableWaiter struct {
	client DescribeTenantDatabasesAPIClient

	options TenantDatabaseAvailableWaiterOptions
}

// NewTenantDatabaseAvailableWaiter constructs a TenantDatabaseAvailableWaiter.
func NewTenantDatabaseAvailableWaiter(client DescribeTenantDatabasesAPIClient, optFns ...func(*TenantDatabaseAvailableWaiterOptions)) *TenantDatabaseAvailableWaiter {
	options := TenantDatabaseAvailableWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = tenantDatabaseAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &TenantDatabaseAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for TenantDatabaseAvailable waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *TenantDatabaseAvailableWaiter) Wait(ctx context.Context, params *DescribeTenantDatabasesInput, maxWaitDur time.Duration, optFns ...func(*TenantDatabaseAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for TenantDatabaseAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *TenantDatabaseAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeTenantDatabasesInput, maxWaitDur time.Duration, optFns ...func(*TenantDatabaseAvailableWaiterOptions)) (*DescribeTenantDatabasesOutput, error) {
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

		out, err := w.client.DescribeTenantDatabases(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for TenantDatabaseAvailable waiter")
}

func tenantDatabaseAvailableStateRetryable(ctx context.Context, input *DescribeTenantDatabasesInput, output *DescribeTenantDatabasesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("TenantDatabases[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
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
		pathValue, err := jmespath.Search("TenantDatabases[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleted"
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
		pathValue, err := jmespath.Search("TenantDatabases[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-parameters"
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
		pathValue, err := jmespath.Search("TenantDatabases[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "incompatible-restore"
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

	return true, nil
}

// TenantDatabaseDeletedWaiterOptions are waiter options for
// TenantDatabaseDeletedWaiter
type TenantDatabaseDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// TenantDatabaseDeletedWaiter will use default minimum delay of 30 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, TenantDatabaseDeletedWaiter will use default max delay of 120
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
	Retryable func(context.Context, *DescribeTenantDatabasesInput, *DescribeTenantDatabasesOutput, error) (bool, error)
}

// TenantDatabaseDeletedWaiter defines the waiters for TenantDatabaseDeleted
type TenantDatabaseDeletedWaiter struct {
	client DescribeTenantDatabasesAPIClient

	options TenantDatabaseDeletedWaiterOptions
}

// NewTenantDatabaseDeletedWaiter constructs a TenantDatabaseDeletedWaiter.
func NewTenantDatabaseDeletedWaiter(client DescribeTenantDatabasesAPIClient, optFns ...func(*TenantDatabaseDeletedWaiterOptions)) *TenantDatabaseDeletedWaiter {
	options := TenantDatabaseDeletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = tenantDatabaseDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &TenantDatabaseDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for TenantDatabaseDeleted waiter. The maxWaitDur
// is the maximum wait duration the waiter will wait. The maxWaitDur is required
// and must be greater than zero.
func (w *TenantDatabaseDeletedWaiter) Wait(ctx context.Context, params *DescribeTenantDatabasesInput, maxWaitDur time.Duration, optFns ...func(*TenantDatabaseDeletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for TenantDatabaseDeleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *TenantDatabaseDeletedWaiter) WaitForOutput(ctx context.Context, params *DescribeTenantDatabasesInput, maxWaitDur time.Duration, optFns ...func(*TenantDatabaseDeletedWaiterOptions)) (*DescribeTenantDatabasesOutput, error) {
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

		out, err := w.client.DescribeTenantDatabases(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for TenantDatabaseDeleted waiter")
}

func tenantDatabaseDeletedStateRetryable(ctx context.Context, input *DescribeTenantDatabasesInput, output *DescribeTenantDatabasesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("length(TenantDatabases) == `0`", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "true"
		bv, err := strconv.ParseBool(expectedValue)
		if err != nil {
			return false, fmt.Errorf("error parsing boolean from string %w", err)
		}
		value, ok := pathValue.(bool)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected bool value got %T", pathValue)
		}

		if value == bv {
			return false, nil
		}
	}

	if err != nil {
		var errorType *types.DBInstanceNotFoundFault
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeTenantDatabases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTenantDatabases",
	}
}
