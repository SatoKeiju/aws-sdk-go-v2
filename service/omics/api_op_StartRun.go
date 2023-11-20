// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/document"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a workflow run. To duplicate a run, specify the run's ID and a role ARN.
// The remaining parameters are copied from the previous run. The total number of
// runs in your account is subject to a quota per Region. To avoid needing to
// delete runs manually, you can set the retention mode to REMOVE . Runs with this
// setting are deleted automatically when the run quoata is exceeded.
func (c *Client) StartRun(ctx context.Context, params *StartRunInput, optFns ...func(*Options)) (*StartRunOutput, error) {
	if params == nil {
		params = &StartRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartRun", params, optFns, c.addOperationStartRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartRunInput struct {

	// To ensure that requests don't run multiple times, specify a unique ID for each
	// request.
	//
	// This member is required.
	RequestId *string

	// A service role for the run.
	//
	// This member is required.
	RoleArn *string

	// A log level for the run.
	LogLevel types.RunLogLevel

	// A name for the run.
	Name *string

	// An output URI for the run.
	OutputUri *string

	// Parameters for the run.
	Parameters document.Interface

	// A priority for the run.
	Priority *int32

	// The retention mode for the run.
	RetentionMode types.RunRetentionMode

	// The run's group ID.
	RunGroupId *string

	// The ID of a run to duplicate.
	RunId *string

	// A storage capacity for the run in gigabytes.
	StorageCapacity *int32

	// Tags for the run.
	Tags map[string]string

	// The run's workflow ID.
	WorkflowId *string

	// The run's workflow type.
	WorkflowType types.WorkflowType

	noSmithyDocumentSerde
}

type StartRunOutput struct {

	// The run's ARN.
	Arn *string

	// The run's ID.
	Id *string

	// The destination for workflow outputs.
	RunOutputUri *string

	// The run's status.
	Status types.RunStatus

	// The run's tags.
	Tags map[string]string

	// The universally unique identifier for a run.
	Uuid *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartRun"); err != nil {
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
	if err = addEndpointPrefix_opStartRunMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opStartRunMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartRun(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opStartRunMiddleware struct {
}

func (*endpointPrefix_opStartRunMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opStartRunMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "workflows-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opStartRunMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opStartRunMiddleware{}, "ResolveEndpointV2", middleware.After)
}

type idempotencyToken_initializeOpStartRun struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartRun) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartRunInput ")
	}

	if input.RequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.RequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartRunMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartRun{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartRun",
	}
}
