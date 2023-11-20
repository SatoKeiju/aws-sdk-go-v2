// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Manually triggers a pipeline to create an image.
func (c *Client) StartImagePipelineExecution(ctx context.Context, params *StartImagePipelineExecutionInput, optFns ...func(*Options)) (*StartImagePipelineExecutionOutput, error) {
	if params == nil {
		params = &StartImagePipelineExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartImagePipelineExecution", params, optFns, c.addOperationStartImagePipelineExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartImagePipelineExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartImagePipelineExecutionInput struct {

	// Unique, case-sensitive identifier you provide to ensure idempotency of the
	// request. For more information, see Ensuring idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html)
	// in the Amazon EC2 API Reference.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the image pipeline that you want to manually
	// invoke.
	//
	// This member is required.
	ImagePipelineArn *string

	noSmithyDocumentSerde
}

type StartImagePipelineExecutionOutput struct {

	// The client token that uniquely identifies the request.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the image that the request created.
	ImageBuildVersionArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartImagePipelineExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartImagePipelineExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartImagePipelineExecution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartImagePipelineExecution"); err != nil {
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
	if err = addIdempotencyToken_opStartImagePipelineExecutionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartImagePipelineExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartImagePipelineExecution(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartImagePipelineExecution struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartImagePipelineExecution) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartImagePipelineExecution) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartImagePipelineExecutionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartImagePipelineExecutionInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartImagePipelineExecutionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartImagePipelineExecution{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartImagePipelineExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartImagePipelineExecution",
	}
}
